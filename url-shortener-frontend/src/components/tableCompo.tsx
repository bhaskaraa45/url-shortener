import React from "react";
import DataModel from "../models/datamodel";
import '../styles/styles.css'
import Popup from 'reactjs-popup';
import 'reactjs-popup/dist/index.css';
import ModalComponent from "./deleteWarningModal.tsx";
import SessionExpiredModal from "./sessionExpireModal.tsx";

function TableComponent({ dataList, deleteFunction }: { dataList: DataModel[] , deleteFunction: Function}) {
    const backendDomain = process.env.REACT_APP_BACKEND_DOMAIN
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
                        <th scope="col" className="px-6 py-3">

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
                                    <a href={"http://" + backendDomain + "/" + data.shorturl} target="_blank" className="font-medium  hover:underline"> {backendDomain}/{data.shorturl}</a>
                                </td>
                                <td className="px-6 py-4">
                                    <a href={data.url} target="_blank" className="font-medium hover:underline">{data.url}</a>
                                </td>
                                <td className="px-6 py-4">
                                    {data.clicked}
                                </td>
                                <td className="px-6 py-4">
                                    {/* <Popup trigger={<button >
                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 30" width="22px" height="22px">
                                            <path fill="#9ca3af" d="M 14.984375 2.4863281 A 1.0001 1.0001 0 0 0 14 3.5 L 14 4 L 8.5 4 A 1.0001 1.0001 0 0 0 7.4863281 5 L 6 5 A 1.0001 1.0001 0 1 0 6 7 L 24 7 A 1.0001 1.0001 0 1 0 24 5 L 22.513672 5 A 1.0001 1.0001 0 0 0 21.5 4 L 16 4 L 16 3.5 A 1.0001 1.0001 0 0 0 14.984375 2.4863281 z M 6 9 L 7.7929688 24.234375 C 7.9109687 25.241375 8.7633438 26 9.7773438 26 L 20.222656 26 C 21.236656 26 22.088031 25.241375 22.207031 24.234375 L 24 9 L 6 9 z" />
                                        </svg>
                                    </button>} position="center center">
                                        <div className="popUpMain"> POPUP</div>
                                    </Popup> */}
                                    {/* <ModalComponent deleteFunction={deleteFunction} id={data.id}/> */}
                                    <SessionExpiredModal/>

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
