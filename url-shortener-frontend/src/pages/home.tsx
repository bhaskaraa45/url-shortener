import React, { useState, useEffect } from "react";
import TableComponent from "../components/tableCompo.tsx";
import DataModel from "../models/datamodel.ts"
import ApiServices from "../services/apiServices.ts";
import '../styles/styles.css'
import { initializeApp } from "firebase/app";
import { getAuth, User } from "firebase/auth";
import { redirect } from "react-router-dom"; // Import Redirect from react-router-dom
import MenuListComposition from "../components/menuComponent.tsx";
import { useNavigate } from 'react-router-dom';

const firebaseConfig = {
    apiKey: "AIzaSyBnCNk7cot7SKQH4KZLi0o3BQtl6JZir6U",
    authDomain: "url-shortener-aa45.firebaseapp.com",
    projectId: "url-shortener-aa45",
    storageBucket: "url-shortener-aa45.appspot.com",
    messagingSenderId: "1058499734182",
    appId: "1:1058499734182:web:24c2bf304c5d9ef9222267",
    measurementId: "G-QDQ1K5VZEH"
};
const app = initializeApp(firebaseConfig)

function HomePage() {
    const [longUrl, setLongUrl] = useState("");
    const [shortUrl, setShortUrl] = useState("https://url.bhaskaraa45.me/");
    const [shortenedUrls, setShortenedUrls] = useState<DataModel[]>([]);
    const [user, setUser] = useState<User>();
    const [isLoggedIn, setLoginStatus] = useState<boolean>(false);
    const myDomain: string = "https://url.bhaskaraa45.me/"

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
                .catch(error => {
                    // console.error('Error fetching shortened URLs:', error);
                });
        }
        return () => unsubscribe();
    }, [isLoggedIn]);


    const handleGenerateClick = () => {
        const parsedUrl = new URL(shortUrl)
        console.log(parsedUrl.pathname)
        ApiServices.generateShortUrl(longUrl, parsedUrl.pathname.substring(1))
            .then(newShortUrl => {
                if (shortenedUrls == undefined) {
                    setShortenedUrls([])
                }
                setShortenedUrls(prevShortenedUrls => [newShortUrl, ...prevShortenedUrls]);
            })
            .catch(error => {
                // Handle error
            });
        setLongUrl("");
    };

    const handleShortUrlChnage = (url: string) => {
        if (!url.startsWith(myDomain)) {
            return;
        }
        if (url.length >= myDomain.length) {
            setShortUrl(url)
        }
    }


    const defaultPic: string = "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Windows_10_Default_Profile_Picture.svg/512px-Windows_10_Default_Profile_Picture.svg.png?20221210150350"
    var imgUrl: string = defaultPic

    // if (user != undefined) {
    //     imgUrl = user.photoURL ? user.photoURL : defaultPic
    //     console.log(imgUrl)
    // }

    return (
        // <div className="main rounded-lg border-rose-500 shadow-lg">
        <div className="main rounded-lg border-rose-500 ">
            <div className="title text-3xl font-semibold text-white/[.85]">
                Shrink your long URL
                <MenuListComposition imgUrl={imgUrl} userName="bhaskar" />
            </div>
            <div className="inputContainers">
                <label className="inputLabel text-lg font-normal text-white/[.85]">
                    Enter a long URL
                    <br />
                    <input
                        className="urlInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md"
                        autoComplete='off'
                        type='text'
                        placeholder="e.g. https://subdomain.long-domain.com/long-path"
                        value={longUrl}
                        onChange={(e) => setLongUrl(e.target.value)}
                    />
                </label>
                <div className="gapInputs">

                </div>
                <label className="inputLabel text-lg font-normal text-white/[.85]">
                    Enter custom path (Optional)
                    <br />
                    <form className="pathForm">
                        <input
                            className="pathInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md"
                            autoComplete='off'
                            type='text'
                            placeholder="e.g. your-path"
                            value={shortUrl}
                            onChange={
                                (e) => handleShortUrlChnage(e.target.value)
                            }
                        />
                        <button onClick={()=> {console.log("helllo")}} className="checkButton">
                            check
                        </button>
                    </form>

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

            <TableComponent dataList={shortenedUrls == undefined || shortenedUrls == null ? [] : shortenedUrls} />

        </div>
    );
}

export default HomePage;
