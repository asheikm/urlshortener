// package model
package model

// UrlStruct
type UrlStruct struct {
	LongUrl  string `json:"url"`
	ShortUrl string `json:"shorturl"`
}

type InputUrl struct {
	Url string `json:"url"`
}
