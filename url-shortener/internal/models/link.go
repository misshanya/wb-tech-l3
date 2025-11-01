package models

type Link struct {
	ID    int64
	URL   string `json:"url"`
	Short string `json:"short"`
}
