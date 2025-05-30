package domain

type Category struct {
	CategoryID     string
	Description    string
	ParentCategory *Category
}
