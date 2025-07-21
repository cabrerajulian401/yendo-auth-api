# Yendo | Backend Transaction Service (Go)

**A scalable proof-of-concept for a core transaction authorization service, engineered with Go.**

**Learned Go Basics and created Project in 1 Weekend: 07/18/2025 - 07/20/2025**

---

## 1. Project Preface & Objectives

### Preface:

This project is the direct result of a self-directed initiative to understand and model the core technical concepts of Yendo's business. Following a conversation on Thursday, I dedicated the weekend to learning Go from the ground up and applying that knowledge to build a functional service that simulates one of Yendo's most critical operations: credit card transactions.

###  Objectives:
 1.  To demonstrate rapid, focused learning and execution: Proving that I can learn new technologies and apply them to deliver functional prototypes under a tight deadline.
 2.  To build a simple & scalable transaction backend service: This simulates backend credit card operations and serves as a lightweight framework for storing & creating customer data transactions. **This service should be able to generate a Transaction ID and Timestamp for a Transaction and be able to record/showcase all Transactions within Server using API Endpoints**

#### Detailed Markdown learning notes of Go Documentation and project plans from this weekend-long effort are aslo included in the repository to document my learning process.

## 2. Business Logic: Why This Matters to Yendo

For a FinTech credit card company, this transaction server simulates the first step of providing a robust and scalable foundation for the following:

* **Handling High-Volume Requests:** To process thousands of credit transactions in a secure manner using specific endpoints and endpoint error logic in the future.
* **Capturing Critical Data:** Securely ingest necessary data points like Amount, Transaction ID, and Timestamp for keeping record of transaction records  using Go's encoding/decoding JSON packages and HTTP server logic.
* **Ensuring Low Latency:** Provide a fast response to customers and partners by using Go for its concurrency.

This prototype proves out the core API structure required to support these critical business needs.

## 3. Core Engineering Principles

This service was built to exemplify basic customer transctions.

* **Performance for a Better Customer Experience:** A customer waiting for a transaction to be approved is a moment of friction. This service is fast by design, using Go's compiled code and efficient in-memory data structures (defined by Go `structs` and `slices`). This translates to near-instantaneous transactions, building customer trust.

* **Type-Safe JSON Processing:** The service ensures data integrity through Go's strict type system and its native `encoding/json` package. By defining a clear `AuthorizationRequest` struct with `json` tags, we enforce a strict contract for all incoming data. The server efficiently **unmarshals** incoming JSON into this struct, rejecting any malformed requests, and **marshals** the Go struct back into JSON for responses, guaranteeing predictable and secure data exchange without external dependencies.

* **Concurrency for Uninterrupted Service:** In the real world, user traffic isn't predictable. A new partnership or peak-hour demand can cause a sudden surge in requests. This service is built using Go's native concurrency model, where each incoming request is handled independently and efficiently. This means the platform can manage thousands of simultaneous transactions without slowing down, ensuring a stable and reliable service for every customer, even during peak load. This is why Go is useful for this.

* **Scalability for Future Growth:** What happens when Yendo grows from 1,000 to 1,000,000 users? A scalable system means you don't have to rebuild everything. The server logic is "stateless," allowing it to be easily replicated across multiple servers to handle massive user growth without a drop in performance without needing previous memory of past requests.

## 4. API Specification & Server Functionality

### POST /authorize

Submits a new credit card transaction for authorization, requiring the card number and transaction amount as JSON parameters.

* **Request Body:**
    ```json
    {
       "cardNumber": "4111222233334444",
       "amount": 2500.00
    }
    ```

Then returns the full body now including a transaction ID and a timestamp of when the transaction occurred.
* **Success Response (`201 Created`):**
    ```json
    {
        "id": "txn_1a2b3c4d5e",
        "cardNumber": "4111222233334444",
        "amount": 2500.00,
        "timestamp": 1679340000
    }
    ```

### GET /authorizations

Retrieves a list of all processed transactions. This is a read-only operation that does not change the server state.

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

## 5. Local Deployment & Testing

1.  **Clone the repository:**
    ```bash
    git clone [your-repository-url]
    cd yendo-auth-api
    ```

2. Install all necessary `package` import dependencies in `main.go` 

  
4.  **Run the service:**
    ```bash
    go run main.go
    ```
5.  The server will start on `http://localhost:8080`.


6.  **Test POST Request:** Create a new transaction using cURL. This requires the card number and transaction amount then returns the server-generated `id` and `timestamp` in a full JSON object.
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"cardNumber":"1234-5678-9012-3456", "amount":99.99}' http://localhost:8080/authorize
    ```

7.  **Test GET Request:** Retrieve all records from the in-memory store. This returns an array of all transaction objects.
    ```bash
    curl http://localhost:8080/authorizations
    ```

## 6. Future Roadmap to a Production-Ready Server

This outlines the strategic evolution of the service from a prototype to a core component of Yendo's business infrastructure.

* **Phase 1: Enhance the Data Fields**
    * **Action:** Expand the `AuthorizationRequest` struct to include Yendo-specific data points critical for collateral-based lending.
    * **New Fields:** `MakeOfCar`, `ValueOfCar`, `FicoScore`.
    * **Benefit:** This transforms the service from a simple transaction processor into a true data-gathering tool for company-specific data points and makes Customer data points more specific

* **Phase 2: Database Integration & Persistence**
    * **Action:** Integrate a production-grade SQL database (e.g., **PostgreSQL**) and implement a data access layer.
    * **Benefit:** Ensures data durability, enables complex stateful queries, and provides a single source of truth for all applications.

* **Phase 3: Observability & Monitoring**
    * **Action:** Instrument the code with structured logging and export metrics (e.g., request latency, error rates) to a system.
    * **Benefit:** Provides critical visibility into system health and performance.

