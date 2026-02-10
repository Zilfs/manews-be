# Manews - News Portal API

Manews is a RESTful API for a news portal built with Go using the Fiber web framework. It follows Clean Architecture principles and provides endpoints for managing users, categories, and content.

## Table of Contents

- [Architecture](#architecture)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Database Structure](#database-structure)
- [API Endpoints](#api-endpoints)
- [Setup and Installation](#setup-and-installation)
- [Docker Support](#docker-support)

## Architecture

This application follows the Clean Architecture pattern with the following layers:

### Layered Structure

1. **Presentation Layer (`internal/adapter/handler`)**
   - Handles HTTP requests and responses
   - Validates incoming requests
   - Maps between HTTP requests/responses and domain models

2. **Domain Layer (`internal/core/domain`)**
   - Contains business logic and entities
   - Defines interfaces for repositories and services
   - Independent of infrastructure concerns

3. **Infrastructure Layer (`internal/adapter/repository`, `internal/adapter/cloudflare`)**
   - Implements repository interfaces using GORM
   - Integrates with external services (Cloudflare R2)

4. **Application Layer (`internal/core/service`)**
   - Orchestrates business logic
   - Implements domain services
   - Coordinates between repositories and handlers

### Key Components

- **Fiber**: High-performance web framework
- **GORM**: ORM library for database operations
- **Viper**: Configuration management
- **JWT**: Authentication and authorization
- **Cloudflare R2**: Image storage
- **PostgreSQL**: Primary database

## Configuration

The application is configured using environment variables. Copy `.env.example` to `.env` and adjust the values as needed:

```bash
APP_ENV="development"
APP_PORT="8080"

DATABASE_PORT=5432
DATABASE_HOST=localhost
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=manews
DATABASE_MAX_OPEN_CONNECTION=10
DATABASE_MAX_IDLE_CONNECTION=10

JWT_SECRET_KEY=your-secret-key
JWT_ISSUER=manews

CLOUDFLARE_R2_API_KEY=your-r2-api-key
CLOUDFLARE_R2_API_SECRET=your-r2-api-secret
CLOUDFLARE_R2_BUCKET_NAME=your-bucket-name
CLOUDFLARE_R2_TOKEN=your-r2-token
CLOUDFLARE_R2_ACCOUNT_ID=your-account-id
CLOUDFLARE_R2_PUBLIC_URL=https://your-domain.r2.cloudflarestorage.com
```

## Dependencies

Main dependencies used in this project:

| Package                                 | Version  | Purpose                     |
| --------------------------------------- | -------- | --------------------------- |
| github.com/gofiber/fiber/v2             | v2.52.10 | Web framework               |
| gorm.io/gorm                            | v1.31.1  | ORM for database operations |
| github.com/gofiber/contrib/swagger      | v1.3.0   | API documentation           |
| github.com/golang-jwt/jwt/v4            | v4.5.2   | JWT token handling          |
| github.com/spf13/viper                  | v1.21.0  | Configuration management    |
| github.com/aws/aws-sdk-go-v2/service/s3 | v1.94.0  | Cloudflare R2 integration   |
| github.com/go-playground/validator/v10  | v10.30.1 | Request validation          |

For a complete list of dependencies, see `go.mod`.

## Database Structure

The application uses PostgreSQL with the following tables:

### Users Table

```sql
CREATE TABLE IF NOT EXISTS "users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NULL DEFAULT current_timestamp
);
```

### Categories Table

```sql
CREATE TABLE IF NOT EXISTS "categories" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    slug VARCHAR(200) UNIQUE NOT NULL,
    created_by_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NULL DEFAULT current_timestamp
);
```

### Contents Table

```sql
CREATE TABLE IF NOT EXISTS "contents" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    excerpt VARCHAR(250) NOT NULL,
    description TEXT NOT NULL,
    image TEXT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PUBLISH',
    tags TEXT NOT NULL,
    created_by_id INT REFERENCES users(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NULL DEFAULT current_timestamp
);
```

## API Endpoints

### Authentication

- `POST /api/login` - User login

### Admin Endpoints (Require Authentication)

#### Categories

- `GET /api/admin/categories/` - Get all categories
- `POST /api/admin/categories/` - Create a new category
- `PUT /api/admin/categories/:categoryId` - Update a category
- `GET /api/admin/categories/:categoryId` - Get category by ID
- `DELETE /api/admin/categories/:categoryId` - Delete a category

#### Contents

- `GET /api/admin/contents/` - Get all contents
- `POST /api/admin/contents/` - Create new content
- `PUT /api/admin/contents/:contentId` - Update content
- `GET /api/admin/contents/:contentId` - Get content by ID
- `DELETE /api/admin/contents/:contentId` - Delete content
- `POST /api/admin/contents/upload-image` - Upload content image to Cloudflare R2

#### Users

- `GET /api/admin/users/profile` - Get user profile
- `PUT /api/admin/users/update-password` - Update user password

### Frontend Endpoints (Public)

- `GET /api/fe/categories` - Get all categories for frontend
- `GET /api/fe/contents` - Get all published contents with filtering
- `GET /api/fe/contents/:contentID` - Get content detail by ID

## Setup and Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Zilfs/manews-be.git
   cd manews-be
   ```

2. Copy and configure the environment file:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run database migrations:

   ```bash
   # Using migrate tool
   migrate -path database/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
   ```

5. Run the application:
   ```bash
   go run main.go start
   ```

## Docker Support

The application includes Docker support for easy deployment:

### Docker Compose

```yaml
version: "3.8"

services:
  db:
    image: postgres:13
    environment:
      POSTGRES_USER: ${DATABASE_USER:-postgres}
      POSTGRES_PASSWORD: ${DATABASE_PASSWORD:-password}
      POSTGRES_DB: ${DATABASE_NAME:-manews}
    ports:
      - "${DATABASE_PORT:-5432}:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

### Building and Running with Docker

```bash
# Build the Docker image
docker build -t manews .

# Run with Docker Compose
docker-compose up -d
```

The application will be available at `http://localhost:8080`
