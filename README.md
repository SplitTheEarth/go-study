# Study App

A web-based study application built with Go that allows users to create decks of flashcards, add questions, and track their learning progress through interactive quizzes.

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/SplitTheEarth/go-study
   cd study-app
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build -o studyapp ./cmd/server
   ```

## Usage

### Starting the Server

```bash
# Run directly with Go
go run ./cmd/server/main.go

# Or run the built binary
./studyapp
```

The server will start on port 8080 by default. You can set a custom port using the `PORT` environment variable:

```bash
PORT=3000 go run ./cmd/server/main.go
```

### API Endpoints

#### Authentication
- `POST /api/users` - Register a new user
- `POST /api/login` - User login

#### Decks
- `GET /api/decks` - List all decks or get specific deck by ID
- `POST /api/decks` - Create a new deck

#### Questions
- `GET /api/questions?deck_id={id}` - Get questions for a specific deck
- `POST /api/questions` - Add a question to a deck

#### Answers
- `POST /api/submit` - Submit an answer to a question

#### General
- `GET /` - Serve the main application page
- `GET /api/status` - API health check

## Testing

The project includes a comprehensive test script that covers all API endpoints:

```bash
# Make the test script executable
chmod +x test_api.sh

# Run the tests
./test_api.sh
```

The test script will:
1. Clean up any existing database
2. Start the server in the background
3. Test user registration and login
4. Test deck creation and retrieval
5. Test question creation and listing
6. Test answer submission
7. Clean up and stop the server

## Development

### Prerequisites
- Go 1.24.5 or later
- SQLite3

### Dependencies
- `github.com/go-chi/chi/v5` - HTTP router
- `github.com/mattn/go-sqlite3` - SQLite driver
- `golang.org/x/crypto` - Cryptographic functions

### Database Schema

The application automatically creates the necessary SQLite tables on startup:
- `users` - User accounts and authentication
- `decks` - Study deck information
- `questions` - Questions belonging to decks
- `user_answers` - Answer submissions and scoring

### Environment Variables

- `PORT` - Server port (default: 8080)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## TODO / Future Enhancements

- [ ] Add session management and JWT tokens
- [ ] Implement spaced repetition algorithm
- [ ] Add deck sharing functionality
- [ ] Create admin dashboard
- [ ] Add study statistics and analytics
- [ ] Implement deck import/export
- [ ] Add multimedia support for questions (images, audio)
- [ ] Create mobile-responsive frontend
- [ ] Add real-time study sessions
- [ ] Implement user profiles and achievements
