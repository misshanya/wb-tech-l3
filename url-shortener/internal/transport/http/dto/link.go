package dto

type CreateShortLinkRequest struct {
	URL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ShortURL string `json:"short_url"`
}
