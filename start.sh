#!/bin/bash

echo "Building and running Go API and Database containers..."

# Run docker-compose up with --build and -d flags
docker-compose up --build -d

echo ""
echo "Process completed!"
echo "The Go API should be running at http://localhost:8080"
echo "Use 'docker-compose logs -f api' to view application logs."