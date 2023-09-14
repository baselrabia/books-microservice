# Books Microservice

The Books Microservice is a simple REST API built using Go (Golang) and the Gin framework. It allows you to manage a book store by performing the following actions:

- Add a new book.
- Update an existing book.
- Delete all books.
- Get a list of all books.
- Get a specific book by its ID.

## Getting Started

## Running the Application with Docker Compose

To run the Books Microservice using Docker Compose, follow these steps:

1. Make sure you have Docker and Docker Compose installed on your system. If not, you can download and install them from [Docker's official website](https://www.docker.com/get-started).

2. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/books-microservice.git
    ```
3. Change to the project directory:

    ```bash
    cd books-microservice
    ```
4. Build the Docker image for the microservice and start the container:

    ```bash
    docker-compose up --build
    ```
This command will build the Docker image using the specified Go version and start the microservice in a container.
### Prerequisites For Running on Local

- Go (Golang) should be installed on your system. If not, you can download and install it from [here](https://golang.org/doc/install).

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/baselrabia/books-microservice.git
    ```

2. Change to the project directory:
   ```bash
    cd books-microservice
    ```
3. Install the required Go packages, including Gin:
    ```bash
    go get -u github.com/gin-gonic/gin
   ```

### Running the Microservice
To start the Books Microservice, run the following command:

```bash
go run main.go
```
 The microservice will start and be accessible at http://localhost:8080.

## API Endpoints
### Add a New Book
Endpoint: POST /books

Request Body:

```json
{
    "id": 1,
    "title": "Sample Book",
    "author": "John Doe",
    "price": 19.99
}
```
### Update a Book
Endpoint: PUT /books/:id (Replace :id with the book ID)

Request Body:

```json
{
    "id": 1,
    "title": "Updated Book",
    "author": "Jane Smith",
    "price": 24.99
}
```
### Delete All Books
Endpoint: DELETE /books
### Get All Books
Endpoint: GET /books
### Get a Specific Book by ID
Endpoint: GET /books/:id (Replace :id with the book ID)
## Example Usage
You can use tools like `curl`, Postman, or any REST client to interact with the API endpoints. Here are some example `curl` commands:

- Adding a New Book:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "title": "Sample Book", "author": "John Doe", "price": 19.99}' http://localhost:8080/books
```


- Updating a Book:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Updated Book", "author": "Jane Smith", "price": 24.99}' http://localhost:8080/books/1
```

- Deleting All Books:

```bash
curl -X DELETE http://localhost:8080/books
```

- Getting All Books:

```bash
curl http://localhost:8080/books
```

- Getting a Specific Book by ID:

```bash
curl http://localhost:8080/books/1
```
Feel free to adapt these commands or use a tool like Postman to interact with the microservice as needed.

 

 