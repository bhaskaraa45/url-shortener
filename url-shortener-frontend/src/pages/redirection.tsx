import React, { useState, useEffect, useRef } from "react";
import { useLocation } from "react-router-dom";
import ApiServices from "../services/apiServices.ts";

function RedirectionPage() {
    const location = useLocation();
    const isFirstTimeRef = useRef(true);
    const [redirectUrl, setRedirectUrl] = useState<string>("");

    const defaultUrl = "https://shrink.bhaskaraa45.me"

    useEffect(() => {
        if (isFirstTimeRef.current) {
            const fetchData = async () => {
                const pathname = location.pathname;
                const temp = pathname.split("/");
                if (temp.length === 2) {
                    const shortUrl = temp[1];
                    try {
                        const response = await ApiServices.getOgUrl(shortUrl);
                        setRedirectUrl(response);
                    } catch (error) {
                        // Handle error
                        console.error(error);
                    }
                } else {
                    setRedirectUrl(defaultUrl);
                }
            };

            fetchData();
            isFirstTimeRef.current = false;
        }
    }, []); // Empty dependency array ensures that the effect runs only once

    if (redirectUrl.length > 5) {
        window.location.href = redirectUrl;
    }

    return (
        <div>
            {/* You can render any additional content here */}
        </div>
    );
}

export default RedirectionPage;
