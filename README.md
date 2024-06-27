# Golang Three-Tier Architecture Unit Test Project

This project is a Golang implementation of a three-tier architecture example, complete with comprehensive unit tests. The project utilizes Gin as the web framework, GORM as the ORM tool, and Testify for testing.

## Project Structure

The project follows a three-tier architecture:

1. Presentation Layer: Uses Gin to handle HTTP requests and responses
2. Business Logic Layer: Contains all business logic
3. Data Access Layer: Utilizes GORM for database operations

## Key Technologies

- [Gin](https://github.com/gin-gonic/gin): A lightweight web framework
- [GORM](https://gorm.io): An excellent Golang ORM library
- [Testify](https://github.com/stretchr/testify): A toolkit with assertion and mocking support

## Running Tests

To run all unit tests, execute the following command in the project root directory:

```bash
go test ./...
```
