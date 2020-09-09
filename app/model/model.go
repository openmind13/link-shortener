package model

// JSON entities

// RequestAddRandom - request for adding new random short url
type RequestAddRandom struct {
	LongURL string `json:"longurl"`
}

// ResponseAddRandom - response on request adding random short url
type ResponseAddRandom struct {
	ShortURL string `json:"shorturl"`
}

// RequestAddCustom - request for adding new custom short url
type RequestAddCustom struct {
	LongURL  string `json:"longurl"`
	ShortURL string `json:"shorturl"`
}

// ResponseAddCustom - response on request adding custom short url
type ResponseAddCustom struct {
	ShortURL string `json:"shorturl"`
}
