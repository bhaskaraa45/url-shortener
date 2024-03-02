class DataModel {
    url: string;
    shorturl: string;
    clicked: number;

    constructor(url: string, shorturl: string, clicked: number) {
        this.url = url;
        this.shorturl = shorturl;
        this.clicked = clicked;
    }
}

export default DataModel;
