package models

type Book struct {
	ID     string `json:"id"`
	ISBN   string `bson:"isbn"`
	Title  string `json:"title"`
	Price  float32
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
