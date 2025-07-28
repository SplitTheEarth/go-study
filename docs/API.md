# StudyApp API Documentation

## Overview

The StudyApp API is a RESTful API that provides endpoints for managing users, decks, questions, and study sessions. All endpoints return JSON responses unless otherwise specified.

**Base URL:** `http://localhost:8080`

## Table of Contents

- [Authentication](#authentication)
- [Users](#users)
- [Decks](#decks)
- [Questions](#questions)
- [Study Sessions](#study-sessions)
- [Error Handling](#error-handling)
- [Response Formats](#response-formats)

## Authentication

Currently, the API uses simple email/password authentication. Future versions will implement JWT tokens.

### Login
```http
POST /api/login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "user_id": 1,
  "username": "john_doe",
  "email": "user@example.com",
  "score": 150
}
```

**Error Response (401 Unauthorized):**
```
Invalid credentials
```

## Users

### Create User (Sign Up)
```http
POST /api/users
```

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "user@example.com",
  "score": 0
}
```

**Error Responses:**
- `400 Bad Request` - Invalid JSON or missing fields
- `409 Conflict` - Username or email already exists

## Decks

### Get All Decks
```http
GET /api/decks
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Go Basics",
    "description": "Fundamental concepts of the Go language",
    "createdAt": "2024-01-15T10:30:00Z"
  },
  {
    "id": 2,
    "name": "SQL Fundamentals",
    "description": "Basic SQL operations and queries",
    "createdAt": "2024-01-16T14:20:00Z"
  }
]
```

### Get Deck by ID
```http
GET /api/decks?id={deck_id}
```

**Parameters:**
- `id` (required) - The deck ID

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Go Basics",
  "description": "Fundamental concepts of the Go language",
  "createdAt": "2024-01-15T10:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Missing deck ID
- `404 Not Found` - Deck not found

### Create Deck
```http
POST /api/decks
```

**Request Body:**
```json
{
  "name": "JavaScript Fundamentals",
  "description": "Core concepts of JavaScript programming"
}
```

**Response (201 Created):**
```json
{
  "id": 3,
  "name": "JavaScript Fundamentals",
  "description": "Core concepts of JavaScript programming",
  "createdAt": "2024-01-17T09:15:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid JSON or missing required fields
- `500 Internal Server Error` - Database error

### Get Deck by ID (Alternative Route)
```http
GET /api/decks/{id}
```

**Parameters:**
- `id` (path parameter) - The deck ID

**Response:** Same as GET /api/decks?id={deck_id}

## Questions

### Add Question to Deck
```http
POST /api/questions
```

**Request Body:**
```json
{
  "deckId": 1,
  "questionText": "What does the := operator do in Go?",
  "answer": "It declares and initializes a variable",
  "options": [
    "It declares and initializes a variable",
    "It assigns a value to an existing variable",
    "It compares two values",
    "It creates a constant"
  ]
}
```

**Response (201 Created):**
```json
{
  "id": 5,
  "deckId": 1,
  "questionText": "What does the := operator do in Go?",
  "answer": "It declares and initializes a variable",
  "options": [
    "It declares and initializes a variable",
    "It assigns a value to an existing variable",
    "It compares two values",
    "It creates a constant"
  ],
  "createdAt": "2024-01-17T11:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid JSON or missing required fields
- `500 Internal Server Error` - Database error

### Get Questions by Deck
```http
GET /api/questions?deck_id={deck_id}
```

**Parameters:**
- `deck_id` (required) - The deck ID to get questions for

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "deckId": 1,
    "questionText": "What does the := operator do in Go?",
    "answer": "It declares and initializes a variable",
    "options": [
      "It declares and initializes a variable",
      "It assigns a value to an existing variable",
      "It compares two values",
      "It creates a constant"
    ],
    "createdAt": "2024-01-15T10:45:00Z"
  },
  {
    "id": 2,
    "deckId": 1,
    "questionText": "Which keyword is used to create a function in Go?",
    "answer": "func",
    "options": [
      "function",
      "func",
      "def",
      "fn"
    ],
    "createdAt": "2024-01-15T10:50:00Z"
  }
]
```

**Error Responses:**
- `400 Bad Request` - Missing deck_id parameter
- `500 Internal Server Error` - Database error

## Study Sessions

### Submit Answer
```http
POST /api/submit
```

**Request Body:**
```json
{
  "user_id": 1,
  "question": 5,
  "answer": "It declares and initializes a variable"
}
```

**Response (200 OK) - Correct Answer:**
```json
{
  "correct": true,
  "message": "Correct answer!",
  "score": 175
}
```

**Response (200 OK) - Incorrect Answer:**
```json
{
  "correct": false,
  "message": "Incorrect. The correct answer is: It declares and initializes a variable",
  "score": 150
}
```

**Error Responses:**
- `400 Bad Request` - Invalid JSON or missing required fields
- `404 Not Found` - Question or user not found
- `500 Internal Server Error` - Database error

## General Endpoints

### Health Check
```http
GET /api/status
```

**Response (200 OK):**
```json
{
  "status": "OK"
}
```

### Hello (Root)
```http
GET /
```

**Response:** Serves the main HTML file

## Error Handling

The API uses standard HTTP status codes to indicate the success or failure of requests:

- `200 OK` - The request was successful
- `201 Created` - A new resource was successfully created
- `400 Bad Request` - The request was invalid or missing required parameters
- `401 Unauthorized` - Authentication failed
- `404 Not Found` - The requested resource was not found
- `409 Conflict` - The request conflicts with the current state (e.g., duplicate username)
- `500 Internal Server Error` - An error occurred on the server

### Error Response Format

Most error responses return plain text messages:
```
Invalid credentials
```

Some endpoints may return JSON error objects:
```json
{
  "error": "Invalid request format",
  "details": "Missing required field: username"
}
```

## Response Formats

### User Object
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "user@example.com",
  "score": 150
}
```

### Deck Object
```json
{
  "id": 1,
  "name": "Go Basics",
  "description": "Fundamental concepts of the Go language",
  "createdAt": "2024-01-15T10:30:00Z"
}
```

### Question Object
```json
{
  "id": 1,
  "deckId": 1,
  "questionText": "What does the := operator do in Go?",
  "answer": "It declares and initializes a variable",
  "options": [
    "It declares and initializes a variable",
    "It assigns a value to an existing variable",
    "It compares two values"
  ],
  "createdAt": "2024-01-15T10:45:00Z"
}
```

## Database Schema

### Users Table
- `id` (INTEGER, PRIMARY KEY, AUTO_INCREMENT)
- `username` (TEXT, NOT NULL, UNIQUE)
- `email` (TEXT, NOT NULL, UNIQUE)
- `password_hash` (TEXT, NOT NULL)
- `score` (INTEGER, NOT NULL, DEFAULT 0)

### Decks Table
- `id` (INTEGER, PRIMARY KEY, AUTO_INCREMENT)
- `name` (TEXT, NOT NULL)
- `description` (TEXT)
- `created_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP)

### Questions Table
- `id` (INTEGER, PRIMARY KEY, AUTO_INCREMENT)
- `deck_id` (INTEGER, NOT NULL, FOREIGN KEY)
- `question_text` (TEXT, NOT NULL)
- `answer` (TEXT, NOT NULL)
- `options` (TEXT) - JSON array stored as text
- `created_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP)

### User_Answers Table
- `id` (INTEGER, PRIMARY KEY, AUTO_INCREMENT)
- `user_id` (INTEGER, NOT NULL, FOREIGN KEY)
- `question_id` (INTEGER, NOT NULL, FOREIGN KEY)
- `user_answer` (TEXT, NOT NULL)
- `is_correct` (BOOLEAN, NOT NULL)
- `answered_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP)

## Example Usage

### Complete Flow Example

1. **Create a user:**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

2. **Login:**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

3. **Create a deck:**
```bash
curl -X POST http://localhost:8080/api/decks \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Deck","description":"A deck for testing"}'
```

4. **Add a question:**
```bash
curl -X POST http://localhost:8080/api/questions \
  -H "Content-Type: application/json" \
  -d '{"deckId":1,"questionText":"What is 2+2?","answer":"4","options":["3","4","5","6"]}'
```

5. **Get questions for study:**
```bash
curl http://localhost:8080/api/questions?deck_id=1
```

6. **Submit an answer:**
```bash
curl -X POST http://localhost:8080/api/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"question":1,"answer":"4"}'
```

## Notes

- All timestamps are in RFC3339 format (ISO 8601)
- The API currently doesn't implement authentication middleware, so endpoints don't require authentication headers
- Question options are stored as JSON arrays in the database
- User scores are updated automatically when answers are submitted
- The API supports CORS for frontend integration

## Future Enhancements

- JWT token-based authentication
- Rate limiting
- Pagination for large result sets
- Advanced filtering and sorting
- File upload for images in questions
- Bulk operations for questions
- User progress tracking endpoints
- Deck sharing and collaboration features