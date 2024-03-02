import React from "react";
import DataModel from "../models/datamodel";
import '../styles/styles.css'

function TableComponent({ dataList }: { dataList: DataModel[] }) {
    return (
        <div className="tableData relative overflow-x-auto shadow-md sm:rounded-lg">
            <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                        <th scope="col" className="px-6 py-3">
                            Shorten Url
                        </th>
                        <th scope="col" className="px-6 py-3">
                            Redirection URL
                        </th>
                        <th scope="col" className="px-6 py-3">
                            Clicked
                        </th>
                    </tr>
                </thead>

                <tbody>
                    {dataList.length === 0 ? (
                        <tr>
                            <td colSpan={3} className="px-6 py-4 text-center">
                                No data available
                            </td>
                        </tr>
                    ) : (
                        dataList.map((data, index) => (
                            <tr key={index} className="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700">
                                <td className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    <a href={data.shorturl} target="_blank" className="font-medium  hover:underline">{data.shorturl}</a>
                                </td>
                                <td className="px-6 py-4">
                                    <a href={data.url} target="_blank" className="font-medium hover:underline">{data.url}</a>
                                </td>
                                <td className="px-6 py-4">
                                    {data.clicked}
                                </td>
                            </tr>
                        ))
                    )}
                </tbody>
            </table>
        </div>
    );
}

export default TableComponent;
