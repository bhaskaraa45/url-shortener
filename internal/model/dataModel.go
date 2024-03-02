package model

type DataModel struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shorturl"`
	Clicked  int    `json:"clicked"`
}
