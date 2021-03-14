package models

// Book Model -> table name 'book' in the db
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateBook -> to validate json input for creating a new book record
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
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

// CreateBook : creates a new book record in the 'book' table and returns it
func (book *Book) CreateBook() (*Book, error) {
	err := DB.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}
