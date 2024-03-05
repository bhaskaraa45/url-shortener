class ApiServices {

    static endpoint: string = process.env.REACT_APP_BACKEND_DOMAIN;
    static async generateShortUrl(url: string) {
        // const endpoint: string = "localhost:8080";
        console.log(ApiServices.endpoint)
        try {
            const response = await fetch(`http://${ApiServices.endpoint}/add`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "url": url
                })
            });

            console.log("HELLO WORLD")

            if (!response.ok) {
                throw new Error('Failed to generate short URL');
            }

            const data = await response.json();
            console.log(data["shorturl"]);
            return data;
        } catch (error) {
            console.error(error);
        }
    }

    static async getAllUrls() {
        try {
            const response = await fetch(`http://${ApiServices.endpoint}/getAll`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            const data = await response.json();
            return data;
        } catch (error) {
            console.error(error);
        }
    }

    static async getOgUrl(shorturl: string) {
        try {
            const response = await fetch(`http://${ApiServices.endpoint}/${shorturl}`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            const data = await response.json();
            return data["url"];
        } catch (error) {
            console.error(error);
        }
    }

    static async verifyUser(idToken: string) {
        try {
            const response = await fetch(`http://${ApiServices.endpoint}/verify`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "idToken": idToken
                })
            });
            return response
            
        } catch (error) {
            console.error(error);
        }
    }
}

export default ApiServices;
