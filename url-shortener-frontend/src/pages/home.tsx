import React, { useState, useEffect } from "react";
import TableComponent from "../components/tableCompo.tsx";
import DataModel from "../models/datamodel.ts"
import ApiServices from "../services/apiServices.ts";
import '../styles/styles.css'

function HomePage() {
    const [longUrl, setLongUrl] = useState("");
    const [shortenedUrls, setShortenedUrls] = useState<DataModel[]>([]);

    useEffect(() => {
        ApiServices.getAllUrls()
            .then(response => {
                setShortenedUrls(response);
            })
            .catch(error => {
                // console.error('Error fetching shortened URLs:', error);
            });
    }, []);


    const handleGenerateClick = () => {
        console.log(longUrl);
        ApiServices.generateShortUrl(longUrl)
            .then(newShortUrl => {
                // Append the new short URL to the existing list
                setShortenedUrls(prevShortenedUrls => [...prevShortenedUrls, newShortUrl]);
            })
            .catch(error => {
                // Handle error
            });
        setLongUrl("");
    };
    
    
    console.log(shortenedUrls)

    return (
        <div className="main rounded-lg border-rose-500 shadow-lg">
            <div className="title text-3xl font-semibold text-white/[.85]">
                Shrink your long URL
            </div>
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

            <TableComponent dataList={shortenedUrls!=null  ? shortenedUrls: []} />

        </div>
    );
}

export default HomePage;
