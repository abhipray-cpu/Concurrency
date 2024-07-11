# Concurrent Web Crawler Project

This project is a concurrent web crawler developed in Go, featuring a user interface built with Vue.js. It allows users to interact with crawled data and add new URLs to crawl. The system is designed to efficiently process multiple URLs in parallel, leveraging the concurrency model of Go.The UI is available at port 80

## Technology Stack

- **Backend**: Go (Golang)
- **Frontend**: Vue.js
- **Database**: MongoDB, Elasticsearch
- **Containerization**: Docker

## Features

- **Concurrent Crawling**: Utilizes Go's concurrency features to crawl multiple URLs simultaneously.
- **Interactive UI**: Built with Vue.js, the UI provides a seamless experience for interacting with the crawled data.
- **Data Storage**: Uses MongoDB for storing crawl metadata and Elasticsearch for indexing and searching the crawled content.
- **Docker Integration**: The entire application is containerized using Docker, simplifying deployment and environment setup.

## Getting Started

To get the project up and running on your local machine, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone <repository-url>

2. **Run THE PROJECT**:

   ```bash
    cd ./project
    make up_build
