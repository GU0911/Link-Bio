# Link-Bio

[![Go](https://img.shields.io/badge/Go-v1.20%2B-blue?logo=go&logoColor=white)](https://go.dev/)
[![License: Unlicensed](https://img.shields.io/badge/License-Unlicensed-red)](https://choosealicense.com/licenses/unlicensed/)
[![Docker](https://img.shields.io/badge/Docker-ready-blue?logo=docker&logoColor=white)](https://www.docker.com/)

## Project Description 📝

This project, Link-Bio, provides a streamlined and efficient backend service for creating and managing personalized link-in-bio pages. In today's digital landscape, a concise and customizable "link in bio" is crucial for directing social media followers and target audiences to relevant content, products, or services. Link-Bio simplifies this process by offering a robust API for managing links, enabling users to easily create, update, and delete links associated with their profile.

This service is designed for individuals, influencers, businesses, and anyone seeking a centralized hub to share multiple links through a single, shareable URL. It addresses the limitation of most social media platforms that only allow one link in the profile description. By leveraging the power of Go, Link-Bio offers a performant and scalable solution, making it suitable for handling a large number of users and links. The use of Docker and docker-compose further simplifies deployment and ensures consistent behavior across different environments.

The architecture focuses on maintainability and extensibility, making it easy to integrate new features and adapt to evolving requirements. By providing a clean API and well-defined data models, Link-Bio aims to empower users with a flexible and reliable tool for managing their online presence and driving traffic to their desired destinations. The project uses a standard Go project structure for easy navigation and understanding.

## Key Features ✨

*   **Link Creation:** Allows users to easily create new links with custom titles and URLs. Benefit: Enables users to consolidate all important links into a single, manageable resource.
*   **Link Retrieval:** Provides an API to retrieve all links associated with a specific user. Benefit: Enables dynamic display of links on the user's bio page.
*   **Link Update:** Enables users to modify existing links, including their titles and URLs. Benefit: Keeps the link-in-bio page up-to-date with the latest content.
*   **Link Deletion:** Allows users to remove outdated or irrelevant links. Benefit: Ensures the link-in-bio page remains clean and focused.
*   **Database Persistence:** Persists link data in a database for reliable storage and retrieval. Benefit: Guarantees data integrity and availability.
*   **Dockerized Deployment:** Simplifies deployment and ensures consistent behavior across different environments. Benefit: Makes it easy to deploy and run the application in various environments, from local development to production.
*   **.env Configuration:** Enables easy configuration of the application environment through environment variables. Benefit: Allows for easy customization of database connections and other settings without modifying the code.

## Tech Stack & Tools 🛠️

| Technology   | Description                                                 |
| :----------- | :---------------------------------------------------------- |
| Go           | Programming language used for the backend service.         |
| Docker       | Containerization platform for packaging and deploying the application.  |
| Docker Compose | Tool for defining and running multi-container Docker applications.   |
| .env         | Used for managing environment variables.                   |
| Git          | Version control system used for tracking code changes.         |

## Installation & Running Locally 🚀

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.20 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

**Steps:**

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/Gaeuly/Link-Bio.git
    ```

2.  **Navigate to the project directory:**

    ```bash
    cd Link-Bio
    ```

3.  **Configure Environment Variables:**

    Create a `.env` file based on the `.env.example` file. Replace the placeholder values with your actual configuration. At a minimum, you'll need to set database connection details.

    ```bash
    cp .env.example .env
    # Edit the .env file with your preferred editor, e.g., nano, vim, or VS Code.
    nano .env
    ```
    Make sure to set the correct database parameters in the .env file.

4. **Build and Run the Docker containers:**

    ```bash
    docker-compose up --build
    ```

    This command will build the Docker image (if it doesn't exist) and start the application containers.  The `--build` flag ensures that the image is rebuilt if the `Dockerfile` has been modified.

5.  **Access the API:**

    The API should now be accessible at `http://localhost`.  You'll need to consult any relevant API documentation to determine the specific endpoints.

6.  **Run the Application Directly (Without Docker):**
    Alternatively, you can run the application directly using Go:
    ```bash
    go mod download
    go run main.go
    ```

## How to Contribute 🤝

We welcome contributions to Link-Bio! To contribute:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes.
4.  Write tests to ensure your changes are working correctly.
5.  Submit a pull request with a clear description of your changes.