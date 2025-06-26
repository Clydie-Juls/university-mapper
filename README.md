# University Mapper

A full-stack web application to map and explore university data, featuring real-world data scraping, API integration, concurrency, and advanced algorithms like cosine similarity for searching.

---

## Features

- Full-stack app built with Go (backend) and Next.js (frontend)
- Scrapes and seeds real-world university data into a PostgreSQL database
- API integration with concurrency for efficient data handling
- Implements cosine similarity and other algorithms for data processing
- Responsive UI built with Tailwind CSS and React
- Error handling throughout the stack
- Public access (no authentication) to maximize usability

---

## Getting Started

### Prerequisites

- Docker (for PostgreSQL)
- Go 1.18+ (for backend and scraper)
- Node.js and npm (for frontend)

---

### Setup Instructions

Run these commands in sequence to get the app up and running:

```bash
# 1. Start PostgreSQL with Docker
docker run --name university-mapper-db -e POSTGRES_PASSWORD=12345678 -p 5432:5432 -d postgres

# 2. Seed the database by running the scraper
cd scraper
go run .

# 3. Run the backend API server (in a new terminal)
cd ../api
go run main.go

# 4. Run the frontend app (in another terminal)
cd ../client
npm install
npm run dev

