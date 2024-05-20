package model

type Eligible struct {
	Day         string `json:"day"`
	Month       string `json:"month"`
	Year        string `json:"year"`
	Eligibility bool   `json:"eligibility"`
}
