package model

type DataModel struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	ShortUrl string `json:"shorturl"`
	Clicked  int    `json:"clicked"`
	UserID   int    `json:"user_id"`
}
