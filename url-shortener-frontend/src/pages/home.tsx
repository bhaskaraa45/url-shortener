import React, { useState, useEffect } from "react";
import { useNavigate } from 'react-router-dom';
import TableComponent from "../components/tableCompo.tsx";
import DataModel from "../models/datamodel.ts"
import ApiServices from "../services/apiServices.ts";
import '../styles/styles.css'
import { initializeApp } from "firebase/app";
import { getAuth, User } from "firebase/auth";
import MenuListComposition from "../components/menuComponent.tsx";

const firebaseConfig = {
    apiKey: "AIzaSyBnCNk7cot7SKQH4KZLi0o3BQtl6JZir6U",
    authDomain: "url-shortener-aa45.firebaseapp.com",
    projectId: "url-shortener-aa45",
    storageBucket: "url-shortener-aa45.appspot.com",
    messagingSenderId: "1058499734182",
    appId: "1:1058499734182:web:24c2bf304c5d9ef9222267",
    measurementId: "G-QDQ1K5VZEH"
};

initializeApp(firebaseConfig);

function HomePage() {
    const myDomain:string = process.env.REACT_APP_BACKEND_DOMAIN;

    const [longUrl, setLongUrl] = useState("");
    const [shortUrl, setShortUrl] = useState(`${myDomain}/`);
    const [shortenedUrls, setShortenedUrls] = useState<DataModel[]>([]);
    const [user, setUser] = useState<User>();
    const [isLoggedIn, setLoginStatus] = useState<boolean>(false);
    const [isAvlbl, setAvaliability] = useState<boolean>();
    const [isValidUrl, setValidity] = useState<boolean>();

    const navigate = useNavigate();

    useEffect(() => {
        const unsubscribe = getAuth().onAuthStateChanged(user => {
            if (user) {
                setUser(user);
                setLoginStatus(true)
            } else {
                setUser(undefined)
                setLoginStatus(false);
                navigate("/auth")
            }
        });

        if (isLoggedIn) {
            ApiServices.getAllUrls()
                .then(response => {
                    if (response) {
                        setShortenedUrls(response);
                    }
                })
        }
        return () => unsubscribe();
    }, [isLoggedIn, navigate]);


    const handleGenerateClick = () => {
        const backend = new URL(myDomain)

        if (!validateUrl(longUrl)) {
            setValidity(false)
            return
        }
        if (longUrl.length === 0) {
            setValidity(false)
            return
        }
        if (longUrl.includes(backend.hostname)) {
            setValidity(false)
            return
        }

        if (!shortUrl.startsWith(myDomain)) {
            setAvaliability(false)
        }

        setValidity(true)
        const parsedUrl = new URL(shortUrl)
        ApiServices.generateShortUrl(longUrl, parsedUrl.pathname.substring(1))
            .then(newShortUrl => {
                if (shortenedUrls === undefined) {
                    setShortenedUrls([])
                }
                setShortenedUrls(prevShortenedUrls => [newShortUrl, ...prevShortenedUrls]);
                setLongUrl("");
                setAvaliability(undefined)
                setShortUrl(`${myDomain}/`)
            })
            .catch(error => {
                // Handle error
            });
    };

    const handleShortUrlChange = (url: string) => {
        if (!url.startsWith(`${myDomain}/`)) {
            return;
        }
        if (url.length >= myDomain.length) {
            setShortUrl(url);
        }
    }

    const handleCheckButton = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        const parsedUrl = new URL(shortUrl);
        const path = parsedUrl.pathname.substring(1)
        if (path.length === 0) {
            return
        }
        ApiServices.checkAvaliability(path).then((res) => {
            if (res) {
                setAvaliability(true);
            } else {
                setAvaliability(false);
            }
        });
    };

    function validateUrl(urlString: string): boolean {
        try {
            new URL(urlString);
            return true;
        } catch (error) {
            return false;
        }
    }
    const deleteFunction = (id: number) => {
        ApiServices.delete(id).then((res) => {
            if (res) {
                const index = shortenedUrls.findIndex(item => item.id === id);
                if (index !== -1) {
                    const updatedShortenedUrls = [...shortenedUrls.slice(0, index), ...shortenedUrls.slice(index + 1)];
                    setShortenedUrls(updatedShortenedUrls);
                }
            }
        })
    };

    const defaultPic: string = "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Windows_10_Default_Profile_Picture.svg/512px-Windows_10_Default_Profile_Picture.svg.png?20221210150350"
    var imgUrl: string = defaultPic;
    var userName:string = "user"
    if (user) {
        const temp = user.displayName
        if (temp!=null) {
            userName = temp.split(" ")[0]
        }
    }

    return (
        <div className="main rounded-lg border-rose-500">
            <div className="title text-3xl font-semibold text-white/[.85]">
                Shrink your long URL
                <MenuListComposition imgUrl={imgUrl} userName={userName} />
            </div>
            <div className="inputContainers">
                <label className="inputLabel text-lg font-normal text-white/[.85]">
                    Enter a long URL
                    <br />
                    <input
                        className={isValidUrl === undefined ? "urlInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md" : isValidUrl ? "urlInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md" : "urlInput border-2 rounded-md border-rose-500 focus:border-rose-500 text-base shadow-md"}
                        autoComplete='off'
                        type='text'
                        placeholder="e.g. https://subdomain.long-domain.com/long-path"
                        value={longUrl}
                        onChange={(e) => setLongUrl(e.target.value)}
                    />
                    <p className={isValidUrl === undefined ? "disable" : isValidUrl ? "disable text-rose-500" : "text-rose-500"}>
                        **Invalid URL
                    </p>
                </label>
                <div className="gapInputs"></div>
                <label className="inputLabel text-lg font-normal text-white/[.85]">
                    Enter custom path (Optional)
                    <br />
                    <form className="pathForm">
                        <input
                            className={isAvlbl === undefined ? "pathInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md" : isAvlbl ? "pathInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md" : "pathInput border-2 rounded-md border-rose-500 focus:border-rose-500 text-base shadow-md"}
                            autoComplete='off'
                            type='text'
                            placeholder="e.g. your-path"
                            value={shortUrl}
                            onChange={(e) => handleShortUrlChange(e.target.value)}
                        />
                        <button onClick={handleCheckButton} className="checkButton">
                            check
                        </button>
                    </form>
                    <p className={isAvlbl === undefined ? "disable" : isAvlbl ? "text-green-500" : "disable text-green-500"}>
                        **Available
                    </p>
                    <p className={isAvlbl === undefined ? "disable" : isAvlbl ? "disable text-rose-500" : "text-rose-500"}>
                        **Not Available
                    </p>
                </label>
            </div>
            <button
                type="button"
                className="generateBtn rounded-md shadow-md"
                onClick={handleGenerateClick}
            >
                Generate
            </button>

            <div className="listTitle text-2xl font-medium text-white/[.85]">
                Your shortened URLs
            </div>

            <TableComponent dataList={shortenedUrls === undefined || shortenedUrls === null ? [] : shortenedUrls} deleteFunction={deleteFunction} />
        </div>
    );
}

export default HomePage;
