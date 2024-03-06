class DataModel {
    url: string;
    shorturl: string;
    clicked: number;
    id: number;

    constructor(url: string, shorturl: string, clicked: number, id: number) {
        this.url = url;
        this.shorturl = shorturl;
        this.clicked = clicked;
        this.id = id;
    }
}

export default DataModel;
