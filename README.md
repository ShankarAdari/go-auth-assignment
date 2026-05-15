# Golang Backend Assignment

A production-ready authentication service built with Golang, Gin, and GORM.

## Features
- **Signup API**: Creates a new user with password hashing.
- **Login API**: Validates credentials and returns a JWT token.
- **Role-based Access**:
    - `User`: Can access their own profile.
    - `Admin`: Can access the list of all users.
- **Middleware**: JWT validation and Role protection.
- **Database**: SQLite (Automated migration).

## Tech Stack
- **Language**: Go (Golang)
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Auth**: JWT (JSON Web Tokens)
- **Security**: Bcrypt (Password Hashing)

## API Endpoints
| Method | Endpoint   | Access    | Description |
|--------|------------|-----------|-------------|
| POST   | `/signup`  | Public    | Register a new user |
| POST   | `/login`   | Public    | Login and get JWT |
| GET    | `/profile` | Auth User | View your profile |
| GET    | `/users`   | Admin Only| List all users |

## Setup Instructions

1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **Testing the APIs**:
   - Use Postman or cURL.
   - For protected routes, include the JWT in the header:
     `Authorization: Bearer <your_token>`

## Implementation Details

- **Auth Flow**: The client sends credentials to `/login`. The server verifies them against the database and returns a signed JWT. The client then sends this JWT in subsequent requests to access protected resources.
- **Security**: Passwords are never stored in plain text. We use Bcrypt with a salt cost of 14.
- **Route Protection**: `AuthMiddleware` extracts the JWT, validates its signature, and checks the expiration. If valid, it attaches the user data to the request context.
