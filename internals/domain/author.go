package domain

type Author struct {
	ID         uint   `bson:"id"`
	FirstName  string `bson:"first_name"`
	LastName   string `bson:"last_name"`
	Patronymic string `bson:"patronymic"`
}
