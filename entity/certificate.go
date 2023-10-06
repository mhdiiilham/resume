package entity

type Certificate struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Issuer string `json:"issuer"`
	URL    string `json:"url"`
}
