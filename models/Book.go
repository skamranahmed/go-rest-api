package models

// Book Model -> table name 'book' in the db
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetBooks : returns all the entries of the 'book' table
func (book *Book) GetBooks() (*[]Book, error) {
	books := []Book{}

	err := DB.Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	
	return &books, nil
}
