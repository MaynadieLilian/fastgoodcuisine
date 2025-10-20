package model

type Recipe struct {
	Title        string `json:"title"`
	Time         string `json:"time"`
	Difficulty   string `json:"difficulty"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
}
