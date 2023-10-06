package entity

type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}
