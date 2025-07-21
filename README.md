# Yendo | Backend Authorization Service (Go)

**A high-performance, scalable proof-of-concept for a core transaction authorization service, engineered with Go.**

---

## 1. Overview

This project is a backend service designed to simulate the high-throughput transaction authorization flow required by a modern FinTech platform like Yendo. It provides a RESTful API for processing authorization requests and retrieving transaction data, built to showcase key principles of backend engineering: performance, simplicity, and scalability.

The system is written entirely in Go, leveraging its powerful concurrency model and standard library to create a lightweight, efficient, and dependency-free service. While this prototype utilizes an in-memory data store for performance testing and rapid development, its architecture is designed for a seamless transition to a persistent, production-grade database system.

## 2. Core Engineering Principles Demonstrated

This service was built to exemplify several key backend engineering tenets:

* **Performance:** By using Go and avoiding heavy frameworks, the service is optimized for low-latency request processing and a small memory footprint, critical for financial transactions.
* **Concurrency:** Go's native support for goroutines and channels makes it trivial to handle thousands of concurrent requests efficiently, ensuring the system remains responsive under load.
* **Scalability:** The service is stateless, meaning multiple instances can be deployed behind a load balancer to scale horizontally without any code changes.
* **Maintainability:** With clean, commented, and idiomatic Go code, the service is easy to understand, maintain, and extend.

## 3. System Architecture

The architecture is simple and robust, consisting of two main components:

1.  **HTTP Server:** A lightweight server built using Go's native `net/http` package. It listens for incoming API requests, routes them to the appropriate handler, and manages the request/response lifecycle.
2.  **Request Handlers:** Functions that contain the core business logic. They are responsible for decoding JSON payloads, processing the authorization logic, and encoding JSON responses.

## 4. API Specification

### POST /authorize

Submits a new transaction for authorization.

* **Description:** The primary endpoint for processing a new transaction. It accepts a JSON payload and returns a status confirmation.
* **Request Body:**

    ```json
    {
        "id": "txn_1a2b3c4d5e",
        "cardNumber": "4111222233334444",
        "amount": 99.95,
        "timestamp": 1679340000
    }
    ```

* **Success Response (200 OK):**

    ```json
    {
        "status": "approved",
        "transactionId": "txn_1a2b3c4d5e"
    }
    ```

### GET /authorizations

Retrieves a list of all transactions processed by the service.

* **Description:** An internal endpoint for retrieving the full history of authorization requests for administrative or debugging purposes.
* **Success Response (200 OK):**

    ```json
    [
        {
            "id": "txn_1a2b3c4d5e",
            "cardNumber": "4111222233334444",
            "amount": 99.95,
            "timestamp": 1679340000
        }
    ]
    ```

## 5. Local Deployment

To build and run the service locally, ensure Go is installed.

1.  **Clone the repository:**
    ```bash
    git clone [your-repository-url]
    cd yendo-auth-api
    ```

2.  **Run the service:**
    ```bash
    go run main.go
    ```

3.  The server will start on `http://localhost:8080`.

## 6. Engineering Roadmap

This outlines the strategic evolution of the service from a prototype to a production-ready system.

* **Phase 1: Database Integration & Persistence**
    * **Action:** Integrate a production-grade SQL database (e.g., **PostgreSQL**) and implement a data access layer to handle all database operations.
    * **Benefit:** Ensures data durability and enables complex, stateful queries.

* **Phase 2: Containerization & CI/CD**
    * **Action:** Create a **Dockerfile** to containerize the service. Set up a CI/CD pipeline (e.g., using GitHub Actions) to automate testing and deployment.
    * **Benefit:** Streamlines the development lifecycle and ensures consistent, reliable deployments.

* **Phase 3: Observability & Monitoring**
    * **Action:** Instrument the code with structured logging and export metrics (e.g., request latency, error rates) to a monitoring system like **Prometheus**.
    * **Benefit:** Provides critical visibility into the system's health and performance in a production environment.

* **Phase 4: Advanced Security & Validation**
    * **Action:** Implement strict request validation and secure the API with authentication middleware (e.g., JWT).
    * **Benefit:** Hardens the service against invalid data and unauthorized access, essential for a financial application.
