# Factory Login

A secure authentication service built with Go that implements the Factory Method design pattern for flexible login strategies.

## Project Overview

Factory Login is a robust authentication service that allows users to log in using different authentication methods (email or mobile number/MSISDN). The project demonstrates the implementation of Clean Architecture and Domain-Driven Design (DDD) principles in a Go application.

### Key Features

- Multiple login strategies (email and MSISDN)
- CAPTCHA validation for enhanced security
- Rate limiting to prevent brute force attacks
- JWT token-based authentication
- Clean Architecture and DDD implementation

## Architecture

The project follows Clean Architecture principles with clear separation of concerns:

### Layers

1. **Domain Layer** (`internal/domain/`)
   - Contains business entities, value objects, repository interfaces, and domain services
   - Represents the core business rules and concepts

2. **Application Layer** (`internal/application/`)
   - Contains use cases, DTOs, and business logic orchestration
   - Implements the Factory Method pattern for login strategies

3. **Infrastructure Layer** (`internal/infrastructure/`)
   - Contains implementations of repositories and external services
   - Handles integration with external systems (Redis, Google reCAPTCHA)

4. **Interfaces Layer** (`internal/interfaces/`)
   - Contains HTTP handlers and adapters
   - Manages dependency injection and request/response handling

### Design Patterns

- **Factory Method Pattern**: Used to create different login strategies
- **Strategy Pattern**: Encapsulates different login algorithms
- **Repository Pattern**: Abstracts data access
- **Dependency Injection**: Used for wiring components together

## Technologies

- **Go** (version 1.24): Core programming language
- **Redis**: Used for rate limiting
- **Google reCAPTCHA**: For bot protection
- **JWT**: For token-based authentication

## Setup and Configuration

### Prerequisites

- Go 1.24 or higher
- Redis server
- Google reCAPTCHA account

### Environment Variables

Create a `.env` file based on the provided `.env.example`:

```
GOOGLE_RECAPTCHA_SECRET=your_recaptcha_secret
REDIS_URL=your_redis_url
MAX_RATE_LIMIT=5
MAX_RATE_LIMIT_DURATION_MINUTES=15
SKIP_CAPTCHA=false
```

### Running the Application

1. Clone the repository
2. Set up environment variables
3. Run the application:

```bash
go run cmd/main.go
```

The server will start on port 8080.

## API Usage

### Login Endpoint

```
POST /login
```

Request body:
```json
{
  "identifier": "user@example.com",
  "password": "password123",
  "captcha_code": "captcha_response_from_frontend",
  "type": 1
}
```

Where `type` is:
- `0`: MSISDN (mobile number)
- `1`: Email

Successful response:
```json
{
  "success": true,
  "message": "login successful",
  "access_token": "jwt_token_string",
  "user_id": "user_id",
  "expires_at": 1628097600
}
```

## Project Structure

```
├── cmd
│   └── main.go                 # Application entry point
├── internal
│   ├── application             # Application layer
│   │   ├── dto                 # Data Transfer Objects
│   │   ├── strategy            # Login strategies
│   │   └── usecase             # Use cases
│   ├── domain                  # Domain layer
│   │   ├── entity              # Domain entities
│   │   ├── repository          # Repository interfaces
│   │   ├── service             # Domain service interfaces
│   │   └── valueobject         # Value objects
│   ├── infrastructure          # Infrastructure layer
│   │   ├── external            # External services
│   │   └── persistence         # Repository implementations
│   └── interfaces              # Interfaces layer
│       └── http                # HTTP handlers
├── .env.example                # Example environment variables
└── go.mod                      # Go module definition
```

## Health Check

The application provides a health check endpoint:

```
GET /health
```

Response:
```
OK
```