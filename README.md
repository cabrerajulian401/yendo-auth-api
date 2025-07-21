# Yendo | Backend Authorization Service (Go)

**A scalable proof-of-concept for a core transaction authorization service, engineered with Go**

**Learned Go from almost scratch and created Project in 1 Weekend: 07/18/2025 - 07/20/2025**

---

## 1. Project Genesis & Objective

This project is the direct result of a self-directed initiative to understand and model the core technical challenges of Yendo's business. Following a conversation on Thursday, I dedicated the weekend to learning Go from the ground up and applying that knowledge to build a functional service that simulates one of Yendo's most critical operations: Credit Authorizations.

The objective is twofold:
1.  **To demonstrate rapid, focused learning and execution:** Proving that new technologies can be mastered and applied to deliver functional prototypes under a tight deadline.
2.  **To build a relevant, mission-driven application:** This isn't a generic API. It's a foundational piece of an underwriting engine specifically designed for a company that provides credit to those with non-traditional financial profiles.

Detailed learning notes and project plans from this weekend-long effort are included in the repository to document the development process.

## 2. Business Logic: Why This Matters to Yendo

For a company like Yendo, the initial authorization is more than just a transaction; it's the entry point to a sophisticated risk assessment process. This service simulates that first step, providing a robust and scalable foundation for a system that must:

* **Handle High-Volume Requests:** Reliably process thousands of credit applications in a secure manner
* **Capture Critical Data:** Securely ingest the necessary data points for making an underwriting decision.
* **Ensure Low Latency:** Provide a fast response to customers and partners.

This prototype proves out the core API structure required to support these critical business needs.

## 3. Core Engineering Principles

This service was built to exemplify key backend engineering tenets essential for a FinTech platform:

* **Performance:** Using Go's compiled nature and avoiding heavy frameworks, the service is optimized for low-latency request processing and a minimal memory footprint.
* **Scalability:** The stateless architecture allows for seamless horizontal scaling behind a load balancer, ensuring the system can grow with user demand.
* **Concurrency:** Leveraging Go's native goroutines, the service is architected to handle a high volume of concurrent requests efficiently, crucial for a real-time financial application.
* **Maintainability:** The codebase is clean, commented, and uses idiomatic Go, making it easy to understand, maintain, and extend.

## 4. API Specification

### POST /authorize

Submits a new credit application for authorization.

* **Request Body:**
    ```json
    {
        "id": "txn_1a2b3c4d5e",
        "cardNumber": "4111222233334444",
        "amount": 2500.00,
        "timestamp": 1679340000
    }
    ```
* **Success Response (`201 Created`):**
    ```json
    {
        "status": "approved",
        "transactionId": "txn_1a2b3c4d5e"
    }
    ```

### GET /authorizations

Retrieves a list of all processed credit applications.

* **Success Response (`200 OK`):**
    ```json
    [
        {
            "id": "txn_1a2b3c4d5e",
            "cardNumber": "4111222233334444",
            "amount": 2500.00,
            "timestamp": 1679340000
        }
    ]
    ```

## 5. Local Deployment

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

## 6. Future Roadmap to a Production-Ready Underwriting Engine

This outlines the strategic evolution of the service from a prototype to a core component of Yendo's business infrastructure.

* **Phase 1: Enhance the Data Model for Underwriting**
    * **Action:** Expand the `AuthorizationRequest` struct and database schema to include Yendo-specific data points critical for collateral-based lending.
    * **New Fields:** `MakeOfCar`, `ValueOfCar`, `FicoScore`.
    * **Benefit:** This transforms the service from a simple transaction processor into a true data-gathering tool for the underwriting engine, directly aligning it with Yendo's business model.

* **Phase 2: Database Integration & Persistence**
    * **Action:** Integrate a production-grade SQL database (e.g., **PostgreSQL**) and implement a data access layer.
    * **Benefit:** Ensures data durability, enables complex stateful queries, and provides a single source of truth for all applications.

* **Phase 3: Containerization & CI/CD**
    * **Action:** Create a **Dockerfile** to containerize the service and set up a CI/CD pipeline (e.g., GitHub Actions) to automate testing and deployment.
    * **Benefit:** Streamlines the development lifecycle and ensures consistent, reliable deployments.

* **Phase 4: Observability & Monitoring**
    * **Action:** Instrument the code with structured logging and export metrics (e.g., request latency, error rates) to a system like **Prometheus**.
    * **Benefit:** Provides critical visibility into system health and performance.

* **Phase 5: Advanced Security & Validation**
    * **Action:** Implement strict request validation and secure the API with authentication middleware (e.g., JWT).
    * **Benefit:** Hardens the service against invalid data and unauthorized access, essential for a financial application. 
