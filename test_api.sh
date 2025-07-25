#!/bin/bash

# A script to test the study app API endpoints.
#
# Usage:
# 1. Make sure your Go server code is in the current directory.
# 2. Make this script executable: chmod +x test_api.sh
# 3. Run the script: ./test_api.sh

# --- Configuration ---
BASE_URL="http://localhost:8080"
DB_FILE="./internal/studyapp.db" # Assuming this is the path in your Go code

# --- Helper Functions ---
# Function to print a nice header for each test section
print_header() {
    echo ""
    echo "=========================================================="
    echo "=> $1"
    echo "=========================================================="
}

# --- Script Start ---
echo "Starting API test script..."

# 1. Cleanup old database file to ensure a clean slate
if [ -f "$DB_FILE" ]; then
    echo "Removing old database file: $DB_FILE"
    rm "$DB_FILE"
fi

# 2. Start the Go server in the background
echo "Building and running the Go server in the background..."
go build -o studyapp_server ./cmd/server/main.go
./studyapp_server &
SERVER_PID=$! # Get the Process ID of the last background command

# Give the server a moment to start up
sleep 2

# 3. Run the tests
# Trap command ensures the server is killed even if the script is interrupted
trap "kill $SERVER_PID" EXIT

# === User Authentication ===
print_header "Testing User Signup"
curl -s -X POST -H "Content-Type: application/json" \
    -d '{"username":"testuser", "email":"test@example.com", "password":"somepassword123"}' \
    "$BASE_URL/api/users" | jq .

print_header "Testing Login (Failure - Wrong Password)"
curl -s -X POST -H "Content-Type: application/json" \
    -d '{"email":"test@example.com", "password":"wrongpassword"}' \
    "$BASE_URL/api/login"

print_header "Testing Login (Success)"
curl -s -X POST -H "Content-Type: application/json" \
    -d '{"email":"test@example.com", "password":"somepassword123"}' \
    "$BASE_URL/api/login" | jq .

# === Decks ===
print_header "Testing Create Deck 1"
DECK1_RESPONSE=$(curl -s -X POST -H "Content-Type: application/json" \
    -d '{"name":"Go Basics", "description":"Fundamental concepts of the Go language."}' \
    "$BASE_URL/api/decks")
echo "$DECK1_RESPONSE" | jq .
DECK1_ID=$(echo "$DECK1_RESPONSE" | jq .id)

print_header "Testing Create Deck 2"
curl -s -X POST -H "Content-Type: application/json" \
    -d '{"name":"SQL Basics", "description":"Fundamental concepts of SQL."}' \
    "$BASE_URL/api/decks" | jq .

print_header "Testing List All Decks"
curl -s -X GET "$BASE_URL/api/decks" | jq .

print_header "Testing Get Deck by ID (ID: $DECK1_ID)"
curl -s -X GET "$BASE_URL/api/decks?id=$DECK1_ID" | jq .


# === Questions ===
print_header "Testing Add Question to Deck $DECK1_ID"
curl -s -X POST -H "Content-Type: application/json" \
    -d '{"deckId":'"$DECK1_ID"', "questionText":"What does the `:=` operator do in Go?", "answer":"It declares and initializes a variable.", "options":["It assigns a value", "It compares two values"]}' \
    "$BASE_URL/api/questions" | jq .

print_header "Testing List Questions for Deck $DECK1_ID"
curl -s -X GET "$BASE_URL/api/questions?deck_id=$DECK1_ID" | jq .


# --- Cleanup ---
print_header "Tests finished. Shutting down server."
# The 'trap' command at the start will handle killing the server process.
