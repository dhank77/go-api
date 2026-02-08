# Go API - Cashier Application

A RESTful API built with Go using layered architecture for a point-of-sale (POS) cashier application.

## Tech Stack

- **Go** - Programming language
- **PostgreSQL** (Supabase) - Database
- **pgx** - PostgreSQL driver
- **godotenv** - Environment variable management
- **Viper** - Configuration management

## Project Structure

```
go-api/
├── main.go                 # Application entry point
├── config/                 # Configuration management
├── database/               # Database connection
├── models/                 # Data models/entities
├── repository/             # Data access layer
├── service/                # Business logic layer
├── handler/                # HTTP handlers (controllers)
└── providers/              # Dependency injection
```

## Layered Architecture

This project follows a **layered architecture** pattern for clean separation of concerns:

```
┌─────────────────────────────────────┐
│           HTTP Request              │
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│            Handler Layer            │
│  - Parse HTTP request               │
│  - Validate input                   │
│  - Return HTTP response             │
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│            Service Layer            │
│  - Business logic                   │
│  - Data transformation              │
│  - Orchestration                    │
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│          Repository Layer           │
│  - Database queries                 │
│  - Data persistence                 │
│  - SQL operations                   │
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│             Database                │
└─────────────────────────────────────┘
```

### Layer Responsibilities

| Layer | Responsibility |
|-------|----------------|
| **Handler** | HTTP request/response handling, input validation, routing |
| **Service** | Business logic, data transformation, orchestration |
| **Repository** | Database operations, SQL queries, data persistence |
| **Models** | Data structures, entities, DTOs |
| **Providers** | Dependency injection, service registration |

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/products?name=indom` | Search products by name |
| GET | `/products/{id}` | Get product by ID (with category) |
| POST | `/products` | Create new product |
| PUT | `/products/{id}` | Update product |
| DELETE | `/products/{id}` | Delete product |

### Categories

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | Get all categories |
| GET | `/categories/{id}` | Get category by ID |
| POST | `/categories` | Create new category |
| PUT | `/categories/{id}` | Update category |
| DELETE | `/categories/{id}` | Delete category |

### Transactions

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/checkout` | Process checkout transaction |

**Checkout Request Body:**
```json
{
  "items": [
    { "product_id": 1, "quantity": 2 },
    { "product_id": 3, "quantity": 1 }
  ]
}
```

### Reports

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/report/hari-ini` | Get today's sales summary |
| GET | `/api/report?start_date=2026-01-01&end_date=2026-02-01` | Get sales summary by date range |

**Report Response:**
```json
{
  "total_revenue": 45000,
  "total_transactions": 5,
  "top_product": { "name": "Indomie Goreng", "qty_sold": 12 }
}
```

## Environment Variables

Create a `.env` file in the root directory:

```env
DATABASE_URL=postgres://user:password@host:port/database
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your database URL
3. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## Database Schema

### Products Table
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL,
    category_id INT REFERENCES categories(id)
);
```

### Transactions Tables
```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);
```

## Key Features

- **Search by Name**: Products can be searched using ILIKE query for partial matching
- **Stock Management**: Automatic stock deduction on checkout
- **Batch Insert**: Transaction details are inserted in a single query to avoid N+1 problem
- **Date Range Reports**: Flexible reporting with date filtering
- **Database Transactions**: Checkout uses database transactions for data consistency
