import React, { useState } from "react";
import { useNavigate } from 'react-router-dom';
import ApiServices from "../services/apiServices.ts";

function SessionExpiredModal() {
    const [isVisible, setIsVisible] = useState(false);
    const navigate = useNavigate();


    const toggleModal = () => {
        setIsVisible(!isVisible);
    };

    const handleLoginClick = () => {
        ApiServices.handleFirebseLogout().then(() => {
            navigate("/auth");
            setIsVisible(false);
        });
    };

    return (
        <>
            <button
                data-modal-target="session-expired-modal"
                data-modal-toggle="session-expired-modal"
                type="button"
                onClick={toggleModal}
            >
                Session Expired
            </button>

            {isVisible && (
                <div
                    id="session-expired-modal"
                    tabIndex={-1}
                    className="fixed top-0 right-0 left-0 z-50 flex justify-center items-center w-full h-full bg-black bg-opacity-50 backdrop-filter backdrop-blur-sm"
                >
                    <div className="relative p-4 w-full max-w-md">
                        <div className="relative bg-white rounded-lg shadow">
                            <div className="p-4 md:p-5 text-center">
                                <h3 className="mb-5 text-lg font-normal text-gray-500">
                                    Your session has been expired. Please log in again.
                                </h3>
                                <button
                                    data-modal-hide="session-expired-modal"
                                    type="button"
                                    className="text-white bg-blue-600 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm inline-flex items-center px-5 py-2.5"
                                    onClick={handleLoginClick}
                                >
                                    Log In
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </>
    );
}

export default SessionExpiredModal;
