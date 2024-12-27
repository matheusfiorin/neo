# Space Monitor API

Go REST API providing NASA's Near-Earth Object (NEO) data.

## Setup

1. Install Go
```bash
brew install go
```

2. Set environment variables in `.env`:
```
NASA_API_KEY=your_api_key
PORT=8080
```

3. Run server
```bash
go run cmd/api/main.go
```

## Endpoints

### Health Check
```bash
GET /health
```

### Near-Earth Objects
```bash
GET /neo
```

Query parameters:
- `start_date`: YYYY-MM-DD (default: today)
- `end_date`: YYYY-MM-DD (default: today + 7 days)

Example:
```bash
curl "http://localhost:8080/neo?start_date=2024-01-01&end_date=2024-01-07"
```

## Flutter Integration

Ideal for hazard monitoring dashboard featuring:
- Close approach countdown
- Distance visualization
- Hazard levels
- Size comparisons

## Project Structure
```
├── cmd/
│   └── api/          # Entry point
├── internal/
│   ├── handlers/     # HTTP handlers
│   ├── models/       # Data models
│   └── services/     # Business logic
└── go.mod           
```