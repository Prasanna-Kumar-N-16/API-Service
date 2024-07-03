# Market-360-Backend
## Prerequisites
- Go: https://golang.org/dl/

## Getting Started
- To clone the repository: git clone https://github.com/ThinklusiveTech/Market-360-Backend.git

## Configuration
- The application uses a JSON configuration file to set up different parameters. 
- Ensure that you have a valid config.json file with the necessary configuration.

## Features
### API End Points

- /api/v1/ping: A simple ping endpoint to check if the server is running.
- /api/v1/login: Login API with associated handlers.


## Golang
- Follow the official Go formatting conventions.

### Go Binary

#### Creating Go Binary:
- make build : Generates a go binary

#### Removing Go Binary:
- make remove: Deletes the binary


### Packages:
- Import only the required packages from external sources.
#### Adding Packages

- Command to get the required packages: `go get <package-path>`
- For example, to get the Gin package: `go get github.com/gin-gonic/gin`

#### Update Packages
- Command to update all the packages to their latest versions: `go get -u`
- Command to update a specific package to the latest version: `go get -u <package-path>`

### Cleaning Modules
- make clean: will clean all the modules   

### Testing

- Write unit tests for functions using testing libraries.
- Ensure that your tests cover critical functionality.
- make test: will run test on all the files
- make test_coverage: will say the code coverage

## Guidelines
- In this project, we follow a set of coding standards and best practices to maintain code quality and consistency. 
- Please make sure to adhere to the following guidelines when contributing to this project:

### 1. Component Naming:
- Use lowercase or PascalCase for functions used within file.
- Choose descriptive and meaningful names for functions.
### 2. File Structure

- Organize your files and folders logically.

### 3. Code Style

- Maintain a consistent code style throughout the project.

### 4. Documentation

- Add comments and documentation to your code to make it more understandable.
- Document public functions, APIs, and complex logic in code comments or separate documentation files.

### 5. Error Handling

- Implement proper error handling in functions and API calls to provide a good user experience.

### 6. Reusability

- Aim for reusable functions or methods to avoid code duplication.
- Reuse them wherever possible.

### 7. Dependency Management

- Use Go Modules for managing dependencies
- Define and document dependencies in the 'go.mod' file
