# StudyApp Documentation

This directory contains documentation for the StudyApp project.

## Available Documentation

### API Documentation
- **[API.md](API.md)** - Complete API documentation in Markdown format
- **[api.html](api.html)** - Interactive HTML version of the API documentation

### How to Use

#### Markdown Version
The `API.md` file contains the complete API documentation in Markdown format. You can:
- View it directly on GitHub
- Convert it to other formats using tools like Pandoc
- Use it as reference while developing

#### HTML Version
The `api.html` file is a standalone HTML document that provides:
- Interactive navigation
- Syntax highlighting for code examples
- Professional styling
- Mobile-responsive design

To view the HTML documentation:
1. Open `api.html` in your web browser, or
2. Serve it through your Go application by adding a route in your server

### Adding Documentation Route (Optional)

You can serve the HTML documentation through your Go server by adding this route to your `router.go`:

```go
// Add this import
"path/filepath"

// Add this route in your SetupRoutes function
r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join("docs", "api.html"))
})
```

Then access the documentation at: `http://localhost:8080/docs`

## API Overview

The StudyApp API provides the following main endpoints:

- **Authentication**: `/api/login`
- **Users**: `/api/users`
- **Decks**: `/api/decks`
- **Questions**: `/api/questions`
- **Study Sessions**: `/api/submit`

## Quick Start

1. **Create a user**:
   ```bash
   curl -X POST http://localhost:8080/api/users \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
   ```

2. **Login**:
   ```bash
   curl -X POST http://localhost:8080/api/login \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","password":"password123"}'
   ```

3. **Create a deck**:
   ```bash
   curl -X POST http://localhost:8080/api/decks \
     -H "Content-Type: application/json" \
     -d '{"name":"My First Deck","description":"Learning the basics"}'
   ```

For complete examples and detailed endpoint information, see the full API documentation files.

## Contributing

When adding new endpoints or modifying existing ones:

1. Update the API documentation in both `API.md` and `api.html`
2. Add examples and test the endpoints
3. Update the database schema section if needed
4. Test all examples in the documentation

## Notes

- All timestamps are in RFC3339 format (ISO 8601)
- The API currently doesn't require authentication tokens
- Question options are stored as JSON arrays
- Base URL for local development: `http://localhost:8080`
