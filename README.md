
# Music Shop

A web application for an online musical instruments store with user authentication and role-based access control.

## Features


## User Roles

- **User**: Cart management, bonus system, product comments
- **Quest**: Can see musical instruments, add to cart.
- **Administrator**: Full system access, user management, analytics

## Tech Stack

- Backend: Golang
- Database: SQLlite
- Authentication: JWT / OAuth

## Installation

```bash
git clone <repository-url>
cd music-shop
go mod tidy
go run ./cmd/api
```

## Configuration

Create a `.env` file with required environment variables:

```
PORT=8080
DB_HOST=localhost
DB_NAME=shop
DB_USER=root
DB_PASS=pass
JWT_SECRET=verysecret
```
