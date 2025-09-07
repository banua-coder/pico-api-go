# Pico API Go - COVID-19 Data API

A Go backend service that provides REST API endpoints for COVID-19 data in Indonesia, including national cases and province-level statistics.

## Features

- 🦠 National COVID-19 cases data with daily and cumulative statistics
- 🗺️ Province-level COVID-19 data including ODP/PDP tracking
- 📊 R-rate (reproductive rate) data when available
- 🔍 Date range filtering for all endpoints
- 🚀 Fast and efficient MySQL database integration
- 🔧 Clean architecture with repository and service layers
- 🛡️ CORS support for web frontend integration
- 📝 Structured logging and error handling
- 💾 Environment-based configuration

## API Endpoints

### Health Check
- `GET /api/v1/health` - Service health status

### National Data
- `GET /api/v1/national` - Get all national cases
- `GET /api/v1/national?start_date=2020-03-01&end_date=2020-12-31` - Get national cases by date range
- `GET /api/v1/national/latest` - Get latest national case data

### Province Data
- `GET /api/v1/provinces` - Get all provinces
- `GET /api/v1/provinces/cases` - Get all province cases
- `GET /api/v1/provinces/cases?start_date=2020-03-01&end_date=2020-12-31` - Get province cases by date range
- `GET /api/v1/provinces/{provinceId}/cases` - Get cases for specific province
- `GET /api/v1/provinces/{provinceId}/cases?start_date=2020-03-01&end_date=2020-12-31` - Get province cases by date range

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