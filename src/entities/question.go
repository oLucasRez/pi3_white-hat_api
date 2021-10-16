package entities

type Question struct {
	Id       int    `bson:"_id" json:"id"`
	Focus    string `json:"focus"`
	Category string `json:"category"`
	Content  string `json:"content"`
}
