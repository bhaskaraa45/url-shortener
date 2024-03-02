import React from "react";
import '../index.css'
import '../styles/styles.css'

function HomePage() {
    return (
        <div className="main rounded-lg border-rose-500 shadow-lg">
            <div className="title text-3xl font-semibold text-white/[.85]">
                Shrink your long URL
            </div>
            <label className="inputLabel text-lg font-medium text-white/[.85]">
                Enter a long URL
                <br />
                <input className="urlInput border-2 rounded-md border-grey focus:border-indigo-500 text-base shadow-md"
                    autoComplete='off'
                    type='text'
                    placeholder="e.g. https://subdomain.long-domain.com/long-path"
                >
                </input>
            </label>
            <button type="button" className="generateBtn rounded-md shadow-md">
                Generate
            </button>
            
        </div>
    );
}

export default HomePage;