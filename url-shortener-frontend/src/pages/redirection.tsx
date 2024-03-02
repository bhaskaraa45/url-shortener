import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";

function RedirectionPage() {
    const location = useLocation();


    useEffect(() => {
        var pathname = location.pathname;
        const temp = pathname.split("/");
        if (temp.length === 2) {
            pathname = temp[1];
        }else{
            window.location.href = "https://bhaskaraa45.me"
        }
        

    }, [location.pathname]); // Adding location.pathname to the dependency array


    return (
        <div>
            <h2>Redirection Page</h2>
            <p>Path: {location.pathname}</p>
        </div>
    );
}

export default RedirectionPage;
