# Golang REST API using GIN and GORM!
## Introduction
A boilerplate code for creating a REST API using Go, GIN and GORM. 
##  Project Structure
    .
    ├── ...
    ├── handler
    │   ├── bookHandler.go      # contains handlers for different routes
    ├── models                  
    │   ├── Book.go             # contains the book model
    │   ├── init.go             # initialize the db connection and automigration
    ├── seed             
    │   ├── seeder.go           # populates the db with dummy data
    ├── .env                    # contains environment variables
    ├── main.go                 # entry point for our app
    └── ...


## How to work on this project on your local system?
- Clone the project
    > #### `git clone https://github.com/skamranahmed/go-rest-api`
- Navigate to the project folder
    > #### `cd go-rest-api`
- Download the project dependcies
    > #### `go mod download`
- Update the .env file as per your configuration
- To run the project
    > #### `go run main.go`

## API Routes
- Get all books
    > #### `GET localhost:8000/books`

- Create a new book record
    > #### `POST localhost:8000/books`

- Get a single book
    > #### `GET localhost:8000/books/:bookID`

- Update an existing book record
    > #### `PUT localhost:8000/books/:bookID`

- Delete an existing book record
    > #### `DELETE localhost:8000/books/:bookID`

