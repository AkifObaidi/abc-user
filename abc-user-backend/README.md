
# 🧠 ABC User Management — Backend (Golang)

Welcome to the **ABC User Management Backend**! This backend service is built using **Golang** and provides a RESTful API to manage users. Below you’ll find how to use the API, configure your environment, and run the service locally or with Docker.

---

## 📦 API Endpoints

### ▶️ Get a List of Users

- **GET** `/users`  
  Fetch all users with support for:
  - `search`: filter by name/email
  - `page`: pagination page number
  - `limit`: number of users per page

```http
GET /users?search=john&page=2&limit=5
```

---

### 🆕 Create a New User

- **POST** `/users`  
  Send user data in the request body to create a new user.

```json
{
  "name": "John",
  "email": "john@example.com",
  "age": 25
}
```

---

### 🔍 Get a User by ID

- **GET** `/users/:id`  
  Fetch a single user by their ID.

---

### ✏️ Update an Existing User

- **PUT** `/users/:id`  
  Update user details with the payload below.

```json
{
  "name": "John Updated",
  "email": "new@example.com",
  "age": 26
}
```

---

### ❌ Delete a User

- **DELETE** `/users/:id`  
  Deletes a user and returns **204 No Content** if successful.

---

## ⚙️ Configuration — Environment Variables

This app uses a `.env` file for configuration. Here’s an example:

```env
SERVER_PORT=8080

DB_TYPE=mysql          # mysql or mongo
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=abc_users

LOG_LEVEL=info
```

---

## 🚀 Run Locally

### ▶️ Run with Docker

```bash
docker run -p 8080:8080 --env-file .env abc-user-backend
```

###  To run the app locally using Go:

```bash
go run .
```

Make sure your `.env` file is properly configured before running.

---

### To run Testcases

```bash
go test ./...
```

##  🔗 Frontend

- **abc-user-frontend** — Vue.js frontend client for this API.
