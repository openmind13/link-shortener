package model

// APIRequest interface
type APIRequest interface {
}

// AddCustomRequest - request for adding new custom short url
type AddCustomRequest struct {
	LongURL  string `json:"longurl"`
	ShortURL string `json:"shorturl"`
}

// AddRequest - request for adding new random short url
type AddRequest struct {
	LongURL string `json:"longurl"`
}
