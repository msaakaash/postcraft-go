<h1 align="center">Postcraft Go</h1>

A lightweight and robust RESTful API built with Go for managing blog posts, including creation, retrieval, updating, and deletion.

## Features
- Create, Read, Update, and Delete (CRUD) posts or blogs
- RESTful API design with clean endpoints
- JSON request and response format
- Error handling with appropriate HTTP status codes
- Built using Go, Gin, GORM for performance and simplicity


## Tech Stack
- [Gin-Gonic](https://github.com/gin-gonic/gin) web framework
- [GORM](https://gorm.io/) ORM library for database operations
- Postman (for API testing)
- PostgreSQL

## Getting Started
Follow these steps to get a local copy of the project up and running:

### Prerequisites
- [Go](https://golang.org/dl/) installed (version 1.18 or higher recommended)
- Git installed
- [GORM](https://gorm.io/) (automatically installed via `go get` during setup)
- (Optional) [Postman](https://www.postman.com/) or any API testing tool to test endpoints


### Installation

1. **Clone the repository**
```bash
git clone https://github.com/msaakaash/postcraft-go.git
cd postcraft-go
```

2. **Initialize Go module**
If you have just cloned the repo and it doesn't contain a `go.mod` file, initialize the Go module by running:
```bash
go mod init postcraft-go
```

3. **Install dependencies**
Run the following command to download and install required packages, including Gin-Gonic:
```bash
go mod tidy
```

4. **Run the application**
Start the API server using:
```bash
go run main.go
```

5. **Test the API**
> **Note:** By default, the server runs on `http://localhost:3000`. Use Postman or curl to interact with the API.

### Example cURL request to get all posts
```bash
curl http://localhost:3000/posts
```
Feel free to customize and extend the API as needed!

## API Endpoints

| Method | Endpoint    | Description                       |
|--------|-------------|-----------------------------------|
| POST   | /posts      | Create a new blog post            |
| GET    | /posts      | Retrieve all blog posts           |
| GET    | /posts/:id  | Retrieve a specific blog post by ID |
| PUT    | /posts/:id  | Update a blog post by ID          |
| DELETE | /posts/:id  | Delete a blog post by ID          |



## Contributing
Contributions are welcome! Please see [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## Code of Conduct
Please read our [Code of Conduct](docs/CODE_OF_CONDUCT.md) before contributing to this project.

## Security
If you discover a vulnerability, please refer to our [Security Policy](docs/SECURITY.md) for instructions on how to report it responsibly.

## License  
This project is licensed under the [MIT LICENSE](LICENSE).


## Author

[**Aakaash M S**](https://github.com/msaakaash)
