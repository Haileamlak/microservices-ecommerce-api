# eCommerce Microservices API

A minimal, production-ready eCommerce backend built with Golang, gRPC, HTTP, MongoDB, and Hexagonal Architecture. The system is divided into independent microservices for scalability and maintainability.

## Microservices

| Service       | Description                                | Protocols Supported |
|---------------|--------------------------------------------|---------------------|
| `product-ms`  | Handles product CRUD and catalog queries   | HTTP, gRPC          |
| `order-ms`    | Manages orders and interacts with payment  | HTTP, gRPC          |
| `payment-ms`  | Handles payments via Stripe                | HTTP only          |
| `user-ms`     | Manages user authentication and sessions   | HTTP, gRPC          |

***All services interact with user-ms for authentication and authorization.***

## Tech Stack

* Language: **Go**
* DB: **MongoDB**
* APIs: **HTTP + gRPC**
* Architecture: **Hexagonal (Clean)**
* Auth: **JWT**
* Payment: **Stripe**

---

## Project Structure (Hexagonal Architecture)

Each microservice follows the same structure:

```

ecommerce/
├── product-ms/
│   ├── internal/
│   │   ├── domain/           # Entity Definitions
│   │   ├── dto/              # Data Transfer Objects
│   │   ├── handler/          # HTTP Handlers
│   │   ├── domain/           # Entity Definitions
│   │   ├── repository/       # Interfaces & DB Implementation
│   │   ├── usecase/          # Business Logic
│   │   ├── infrastructure/
│   │   │   ├── client/       # gRPC files
│   │   │   ├── config/       # Configuration
│   │   │   ├── database/     # Database connection
│   │   │   └── middleware/   # Auth middleware, logging, etc.
│   ├── cmd/              # Main function
│   ├── pkg/              # Utility functions
│   ├── docs/             # Swagger documentation
│   ├── go.mod            # Go module
│   ├── go.sum            # Go sum
│   ├── README.md         # README
│   └── .env              # Environment variables
...

````

---

## Running Locally

### Prerequisites

- Go 1.20+
- MongoDB (local or Atlas)
- [Stripe Developer Account](https://dashboard.stripe.com/register)
- `protoc` (Protocol Buffer Compiler)

### Clone and Install

```bash
git clone https://github.com/haileamlak/ecommerce-with-microservices.git
cd ecommerce-with-microservices/{service-name}
go mod tidy
go run cmd/main.go
````

---

## API Overview

### Authentication (JWT-based)

* `user-ms` issues and verifies JWTs.
* Other services use gRPC calls to `user-ms` for token validation.

### Sample Endpoints

All Services Have Swagger Documentation

* `GET /swagger/*` (works for all services)

---

#### `user-ms`

* `POST /register`
* `POST /login`

* gRPC: `Login`, `VerifyToken`


#### `product-ms`

* `POST /products`
* `GET /products/{id}`
* `GET /products`
* `PUT /products/{id}`
* `DELETE /products/{id}`

* gRPC: `CreateProduct`, `GetProduct`, `GetProducts`, `UpdateProduct`, `DeleteProduct`

#### `order-ms`

* `POST /orders`
* `GET /orders/{id}`
* `GET /orders/user/{userId}`
* `PUT /orders/{id}`
* `DELETE /orders/{id}`
* gRPC: `CreateOrder`, `GetOrder`

#### `payment-ms`

* `POST /initiate-payment`
* `POST /webhook`

---

## Authentication Middleware

* Placed under: `infrastructure/middleware/auth.go`
* Calls `user-ms` over gRPC to verify tokens
* Injects `userID` into request context

---

## Stripe Integration

* Set your Stripe secret key in `.env` or as an environment variable.
* Payments are created from `payment-ms` and verified through webhooks or polling.

---

## License

No-license — free to use, modify, and distribute.
