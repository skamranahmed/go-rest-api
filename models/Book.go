package models

// Book Model -> table name 'book' in the db
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateBookInput -> to validate json input for creating a new book record
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput -> to validate json input for updating an existing book record
type UpdateBookInput struct {
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

// GetBook : finds and returns the book record with the given id
func GetBook(id uint) (*Book, error) {
	book := Book{}
	err := DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return &book, nil
}

// CreateBook : creates a new book record in the 'book' table and returns it
func (book *Book) CreateBook() (*Book, error) {
	err := DB.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

// UpdateBook : updates an existing book record
func (book *Book) UpdateBook() (*Book, error) {
	err := DB.Model(&book).Updates(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

// DeleteBook : deletes a book record
func (book *Book) DeleteBook() (*Book, error) {
	err := DB.Delete(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return &Book{}, nil
}
