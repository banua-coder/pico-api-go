# Pico API Go - COVID-19 Data API

A Go backend service that provides REST API endpoints for COVID-19 data in Indonesia, including national cases and province-level statistics.

## Features

- 🦠 National COVID-19 cases data with daily and cumulative statistics
- 🗺️ Province-level COVID-19 data with enhanced ODP/PDP grouping
- 📊 R-rate (reproductive rate) data when available
- 🔍 Date range filtering for all endpoints
- 📄 **Hybrid pagination system** - efficient for apps, complete for charts
- 📚 **Interactive API Documentation** - Auto-generated OpenAPI/Swagger docs
- 🎯 **Smart query parameters** - flexible data retrieval options
- 🚀 Fast and efficient MySQL database integration
- 🔧 Clean architecture with repository and service layers
- 🛡️ CORS support for web frontend integration
- 📝 Structured logging and error handling
- 💾 Environment-based configuration
- 🚀 **Automatic deployment** with GitHub Actions

## 📚 API Documentation

### Interactive Swagger UI
- **Local development**: http://localhost:8080/swagger/index.html
- **Production**: https://pico-api.banuacoder.com/swagger/index.html

### OpenAPI Specification
- YAML: [`docs/swagger.yaml`](docs/swagger.yaml)
- JSON: [`docs/swagger.json`](docs/swagger.json)

## API Endpoints

### Health Check
- `GET /api/v1/health` - Service health status and database connectivity

### National Data
- `GET /api/v1/national` - Get all national cases
- `GET /api/v1/national?start_date=2020-03-01&end_date=2020-12-31` - Get national cases by date range
- `GET /api/v1/national/latest` - Get latest national case data

### Province Data
- `GET /api/v1/provinces` - Get all provinces with latest case data (default)
- `GET /api/v1/provinces?exclude_latest_case=true` - Get basic province list without case data
- `GET /api/v1/provinces/cases` - Get all province cases (paginated by default)
- `GET /api/v1/provinces/cases?all=true` - Get all province cases (complete dataset)
- `GET /api/v1/provinces/cases?limit=100&offset=50` - Get province cases with custom pagination
- `GET /api/v1/provinces/{provinceId}/cases` - Get cases for specific province (paginated)
- `GET /api/v1/provinces/{provinceId}/cases?all=true` - Get all cases for specific province

### 🆕 Enhanced Query Parameters

**Pagination (All province endpoints):**
- `limit` (int): Records per page (default: 50, max: 1000)
- `offset` (int): Records to skip (default: 0)
- `all` (boolean): Return complete dataset without pagination

**Date Filtering:**
- `start_date` (YYYY-MM-DD): Filter from date
- `end_date` (YYYY-MM-DD): Filter to date

**Province Enhancement:**
- `exclude_latest_case` (boolean): Return basic province list without case data (default includes latest case data)

### 📄 Response Types

**Paginated Response:**
```json
{
  "status": "success",
  "data": {
    "data": [...],
    "pagination": {
      "limit": 50,
      "offset": 0,
      "total": 1000,
      "page": 1,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

**Complete Data Response:**
```json
{
  "status": "success", 
  "data": [...]
}
```

## 🆕 Enhanced Data Structure

### Grouped ODP/PDP Data

Province case data now includes structured ODP (Person Under Observation) and PDP (Patient Under Supervision) data:

```json
{
  "daily": {
    "positive": 150,
    "odp": {
      "active": 5,
      "finished": 20
    },
    "pdp": {
      "active": 8, 
      "finished": 25
    }
  },
  "cumulative": {
    "positive": 5000,
    "odp": {
      "active": 50,
      "finished": 750,
      "total": 800
    },
    "pdp": {
      "active": 20,
      "finished": 580, 
      "total": 600
    }
  }
}
```

## Usage Examples

### For Web Applications (Efficient Loading)
```javascript
// Load first page (default: 50 records)
const response = await fetch('/api/v1/provinces/cases');
const { data, pagination } = response.data;

// Load next page
if (pagination.has_next) {
    const nextPage = await fetch(`/api/v1/provinces/cases?offset=${pagination.offset + pagination.limit}`);
}
```

### For Charts & Analytics (Complete Dataset)
```javascript
// Get complete dataset for time series charts
const response = await fetch('/api/v1/provinces/cases?all=true&start_date=2024-01-01');
const allData = response.data;

// Perfect for Chart.js, D3.js, etc.
const chartData = allData.map(item => ({
    x: item.date,
    y: item.cumulative.positive
}));
```

### For Province-Specific Analysis
```javascript
// Get all Jakarta data
const response = await fetch('/api/v1/provinces/31/cases?all=true');

// Get provinces with their latest statistics (default behavior)
const provincesResponse = await fetch('/api/v1/provinces');
```

## Setup and Installation

### Prerequisites
- Go 1.25+ 
- MySQL database
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/banua-coder/pico-api-go.git
cd pico-api-go
```

2. Copy environment configuration:
```bash
cp .env.example .env
```

3. Update the `.env` file with your database configuration:
```env
DB_HOST=your_db_host
DB_PORT=3306
DB_USERNAME=your_db_username
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
SERVER_HOST=localhost
SERVER_PORT=8080
ENV=development
```

4. Install dependencies:
```bash
go mod tidy
```

5. Run the application:
```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

### Building for Production

```bash
go build -o pico-api-go cmd/main.go
./pico-api-go
```

### Regenerating API Documentation

After modifying handlers or adding new endpoints, regenerate the Swagger docs:

```bash
# Install swag tool (one-time setup)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/main.go -o ./docs
```

## Database Schema

The API uses three main tables:

### national_cases
- Daily national COVID-19 statistics
- Includes positive, recovered, deceased cases
- Cumulative data and R-rate when available

### provinces
- Indonesian province information
- Uses official province codes (e.g., Aceh: 11, Sulawesi Tengah: 72)

### province_cases
- Province-level COVID-19 statistics
- Includes ODP (Orang Dalam Pemantauan) and PDP (Pasien Dalam Pengawasan) tracking
- Links to national_cases for date information

## Shared Hosting Deployment

This API is designed to work with shared hosting environments:

1. Build the binary for your target platform
2. Upload the binary and `.env` file to your hosting provider
3. Ensure your hosting provider supports Go applications
4. Configure the database connection in the `.env` file
5. Start the application

## Development

### Git Flow
This project uses Git Flow for development:

```bash
# Start a new feature
git flow feature start feature-name

# Finish a feature
git flow feature finish feature-name
```

### Project Structure
```
├── cmd/                    # Application entry points
│   └── main.go
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP handlers and routes
│   ├── middleware/       # HTTP middleware
│   ├── models/          # Data models
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
├── pkg/                  # Public packages
│   └── database/        # Database connection
├── .env.example         # Environment configuration template
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch using Git Flow
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.