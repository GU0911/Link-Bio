# Link-Bio

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)](https://go.dev/)
[![License: Unlicensed](https://img.shields.io/badge/License-Unlicensed-red)](https://choosealicense.com/licenses/unlicensed/)
[![Docker](https://img.shields.io/badge/Docker-Enabled-blue?logo=docker)](https://www.docker.com/)

## Project Description 📝

This project, Link-Bio, is a simple and robust Link Bio API built with Go.  It provides a backend service to manage and serve a collection of links, often used to create a personalized landing page containing various social media profiles, websites, and other important resources. The API allows users to create, read, update, and delete links, making it easy to customize and maintain their online presence.

The Link-Bio API is designed to be easily deployable and scalable, thanks to its Docker containerization. This allows for consistent performance across different environments, from local development to production deployments.  The project showcases best practices in Go development, including clear separation of concerns, robust error handling, and comprehensive unit testing (as demonstrated by the `link_handler_test.go` file).

This project is ideal for developers who want to learn how to build a web service in Go, containerize it with Docker, and manage data using a relational database (implied by `database.go`). It also serves as a foundation for creating more complex link management applications, such as link shortening services or advanced analytics dashboards. The modular structure and clear documentation make it easy to extend and adapt the API to specific requirements.

## Key Features ✨

*   **Link Management:**  Create, read, update, and delete links through a RESTful API. This allows users to easily manage their online presence from a central location.
*   **Docker Containerization:** The application is containerized using Docker, enabling easy deployment and scalability across different environments. This ensures consistency and simplifies the deployment process.
*   **Configuration Management:** Utilizes `.env` files and configuration packages to manage environment-specific settings. This improves code maintainability and simplifies configuration changes.
*   **Database Integration:**  Integrates with a relational database (implementation details in `database.go`) to store and manage link data. This provides persistence and allows for efficient data retrieval.
*   **Middleware Support:** Includes middleware for request logging, providing valuable insights into API usage and performance. This aids in debugging and monitoring the application.
*   **Input Validation:** Implements input validation using a validator package (`validator.go`), ensuring data integrity and preventing invalid data from being stored. This enhances the robustness of the application.
*   **REST API:** Exposes the service through a REST API, making it easy to integrate with various front-end applications and services.

## Tech Stack & Tools 🛠️

| Category | Technology/Tool  | Description                                                                             |
| -------- | ---------------- | --------------------------------------------------------------------------------------- |
| Language | Go              | The primary programming language used for the API development.                            |
| Framework| Standard Library | Utilizes Go's standard library for building web services.                               |
| Containerization | Docker        | Used for containerizing the application for easy deployment.                         |
| Orchestration | Docker Compose | Used for defining and managing multi-container Docker applications.                       |
| Configuration | .env files        | Used for storing environment-specific configuration settings.                       |
| Data | Relational Database (Implied)   | Database for storing link information. The specific type is not explicitly defined but likely PostgreSQL or MySQL.    |
| Logging | Custom (src/middleware/logging.go)       | Logs request and response information for debugging and monitoring.                      |
| Validation| Custom (src/util/validator.go)          | Implements input validation to ensure data integrity.                              |

## Installation & Running Locally 🚀

1.  **Prerequisites:**
    *   Go (version 1.21 or later)
    *   Docker
    *   Docker Compose

2.  **Clone the repository:**

    ```bash
    git clone https://github.com/Gaeuly/Link-Bio.git
    ```

3.  **Navigate to the project directory:**

    ```bash
    cd Link-Bio
    ```

4.  **Create `.env` file:**
    Copy the contents of `.env.example` to a new file named `.env` and update the values as needed.
    ```bash
    cp .env.example .env
    ```

5.  **Build and run the application using Docker Compose:**

    ```bash
    docker-compose up --build
    ```

   This command will build the Docker image and start the application in a container. The API will be accessible at the port specified in the `docker-compose.yml` file and `.env` file (usually port 8080 or similar).

6. **Alternatively, you can run the application without Docker:**

    First, you'll need to ensure Go dependencies are downloaded:

    ```bash
    go mod download
    ```
    Then, you can execute:
     ```bash
     go run main.go
     ```

## How to Contribute 🤝

We welcome contributions to the Link-Bio project! If you'd like to contribute, please follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and commit them with clear, descriptive commit messages.
4.  Submit a pull request to the `main` branch.

Please ensure that your code adheres to the project's coding standards and includes relevant unit tests. We appreciate your help in making Link-Bio a better project!