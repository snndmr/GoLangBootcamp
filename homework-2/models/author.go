package models

type Author struct {
	Id   int
	Name string
}

// NewAuthor constructor of Author
func NewAuthor(id int, name string) *Author {
	return &Author{
		Id:   id,
		Name: name,
	}
}
