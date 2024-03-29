import { getAuth, signOut } from "firebase/auth";

class ApiServices {
    static endpoint: string = process.env.REACT_APP_BACKEND_DOMAIN;

    static async handleFirebseLogout() {
        const auth = getAuth();
        signOut(auth).then(() => {
            ApiServices.logout().then((res) => {
                if (res) {
                    this.logout();
                }
            })
        }).catch((error) => {
        });
    }


    static async generateShortUrl(url: string, shorturl: string) {
        try {
            const response = await fetch(`${ApiServices.endpoint}/add`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "url": url,
                    "shorturl": shorturl,
                })
            });

            if (!response.ok) {
                throw new Error('Failed to generate short URL');
            }

            const data = await response.json();
            return data;
        } catch (error) {
            console.error(error);
        }
    }

    static async getAllUrls() {
        try {
            const response = await fetch(`${ApiServices.endpoint}/getAll`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
           
            if (!response.ok) {
                throw new Error('Failed to get all urls');
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error(error);
        }
    }

    static async getOgUrl(shorturl: string) {
        try {
            const response = await fetch(`${ApiServices.endpoint}/${shorturl}`, {
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
            const response = await fetch(`${ApiServices.endpoint}/verify`, {
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

    static async logout() {
        try {
            const response = await fetch(`${ApiServices.endpoint}/logout`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });

            return response.ok
        } catch (error) {
            console.error(error);
        }
    }

    static async checkAvaliability(shorturl: string) {
        try {
            const response = await fetch(`${ApiServices.endpoint}/check`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "shorturl": shorturl
                })
            });

            if (response.status === 200) {
                return true
            } else {
                return false;
            }

        } catch (error) {
            console.error(error);
        }
    }
    static async delete(id: number) {
        try {
            const response = await fetch(`${ApiServices.endpoint}/delete/${id}`, {
                method: 'DELETE',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            return response.ok
        } catch (error) {
            console.error(error);
        }
    }
}

export default ApiServices;
