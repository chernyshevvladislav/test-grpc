package domain

type Book struct {
	ID    uint   `bson:"id"`
	Title string `bson:"title"`
}
