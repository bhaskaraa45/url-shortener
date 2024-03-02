import React from "react";
import '../index.css'
import '../styles/styles.css'
import TableComponent from "../components/tableCompo.tsx";
import DataModel from "../models/datamodel.ts"

function HomePage() {
    const dummyData = new DataModel("https://bhaskar-gamma.vercel.app/","https://shrink.bhaskaraa45.me/portfolio", 0)
    return (
        <div className="main rounded-lg border-rose-500 shadow-lg">
            <div className="title text-3xl font-semibold text-white/[.85]">
                Shrink your long URL
            </div>
            <label className="inputLabel text-lg font-normal text-white/[.85]">
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

            <div className="listTitle text-2xl font-medium text-white/[.85]">
                Your shortened URLs
            </div>

            <TableComponent dataList={[dummyData, dummyData]} />

        </div>
    );
}

export default HomePage;