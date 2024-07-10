## Project Overview

This project is a sophisticated, concurrent web crawler designed to efficiently navigate and extract data from the web. It leverages Go's concurrency features to perform high-speed web crawling, parsing, and data storage. The application's architecture is modular, facilitating easy expansion and maintenance.

### Key Components

- **cmd/crawler/main.go**: Serves as the application's entry point, bootstrapping the web crawler with necessary configurations and initiating the crawling process.

- **pkg/**: Hosts the essential packages that constitute the web crawler's core functionalities:
  - **fetcher/**: Responsible for retrieving web pages, equipped with mechanisms for rate limiting and retrying failed requests.
  - **parser/**: Analyzes the fetched web pages to extract links and other pertinent information.
  - **scheduler/**: Manages the order and priority of web pages to be crawled, optimizing resource utilization.
  - **storage/**: Oversees the storage of extracted data, potentially interfacing with various database systems or file storage solutions.

- **config/config.go**: Manages application configurations, dynamically reading settings from environment variables or configuration files.

- **internal/**: Contains specialized packages not intended for external use but essential for the application's functionality:
  - **models/**: Defines the core data structures utilized throughout the application.
  - **utils/**: Supplies common utility functions that support various application components.

- **test/**: Includes comprehensive unit and integration tests, ensuring the reliability and stability of the application.

- **.env**: Specifies development-related environment variables, such as API keys and database connections.

- **Dockerfile**: Provides Docker with the necessary instructions to containerize the application, streamlining deployment processes.

- **Makefile**: Outlines common development tasks (e.g., build, run, test) in a simplified manner, enhancing the development workflow.

- **README.md**: Delivers a concise overview of the project, including setup guidelines and examples of how to use the application.

This project's structure is designed for scalability and maintainability, making it well-suited for production environments.

### Running the Project

To run the web crawler, a Makefile is provided to simplify common tasks such as building, running, and testing the application. Here's how you can use it:

1. **Build the Application**: Compile the application into an executable.
   ```bash
   make build