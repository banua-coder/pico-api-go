# Sulawesi Tengah COVID-19 Data API Documentation

Version: 2.0.2  
Base URL: `https://pico-api.banuacoder.com/api/v1`

## Overview

This API provides COVID-19 data primarily focused on Sulawesi Tengah (Central Sulawesi), with additional national and provincial data for context. The API supports both paginated responses (for efficient data loading) and complete datasets (for analytics and charts).

## Response Format

### Success Response
```json
{
  "status": "success", 
  "data": { ... }
}
```

### Paginated Response
```json
{
  "status": "success",
  "data": {
    "data": [...],
    "pagination": {
      "limit": 50,
      "offset": 0, 
      "total": 1000,
      "total_pages": 20,
      "page": 1,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

### Error Response
```json
{
  "status": "error",
  "message": "Error description"
}
```

## Pagination Parameters

All endpoints support hybrid pagination:

| Parameter | Type | Default | Max | Description |
|-----------|------|---------|-----|-------------|
| `limit` | int | 50 | 1000 | Number of records per page |
| `offset` | int | 0 | - | Number of records to skip |
| `all` | boolean | false | - | Return all data (bypasses pagination) |

## Enhanced Province Data Structure

Province case data now includes grouped ODP/PDP fields:

```json
{
  "day": 100,
  "date": "2024-01-15T00:00:00Z",
  "daily": {
    "positive": 150,
    "recovered": 120,
    "deceased": 10,
    "active": 20,
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
    "recovered": 4500,
    "deceased": 300,
    "active": 200,
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
  },
  "statistics": {
    "percentages": {
      "active": 4.0,
      "recovered": 90.0,
      "deceased": 6.0
    },
    "reproduction_rate": {
      "value": 1.2,
      "upper_bound": 1.5,
      "lower_bound": 0.9
    }
  }
}
```

## Endpoints

### 1. Health Check

**GET** `/health`

Check API health and database connectivity.

**Response:**
```json
{
  "status": "success",
  "data": {
    "status": "healthy",
    "service": "COVID-19 API", 
    "version": "2.0.2",
    "timestamp": "2024-01-15T10:30:00Z",
    "database": {
      "status": "healthy",
      "connections": {
        "open": 2,
        "idle": 1,
        "in_use": 1,
        "max_open": 5,
        "wait_count": 0
      }
    }
  }
}
```

### 2. National Cases

**GET** `/national`

Get national COVID-19 cases data.

**Query Parameters:**
- `start_date` (string, optional): Start date (YYYY-MM-DD)
- `end_date` (string, optional): End date (YYYY-MM-DD)

**Examples:**
```bash
# Get all national data
GET /national

# Get data for specific date range
GET /national?start_date=2024-01-01&end_date=2024-01-31
```

### 3. Latest National Case

**GET** `/national/latest`

Get the most recent national case data.

### 4. Provinces

**GET** `/provinces`

Get list of all provinces with their latest COVID-19 case data (default behavior).

**Query Parameters:**
- `exclude_latest_case` (boolean, optional): Return basic province list without case data

**Examples:**
```bash
# Provinces with latest case data (default)
GET /provinces

# Basic province list without case data
GET /provinces?exclude_latest_case=true
```

### 5. Province Cases

**GET** `/provinces/cases`  
**GET** `/provinces/{provinceId}/cases`

Get COVID-19 cases for all provinces or a specific province.

**Path Parameters:**
- `provinceId` (string, optional): Province ID (e.g., "31" for Jakarta)

**Query Parameters:**
- `limit` (int, optional): Records per page (default: 50, max: 1000)
- `offset` (int, optional): Records to skip (default: 0)
- `all` (boolean, optional): Return all data without pagination
- `start_date` (string, optional): Start date (YYYY-MM-DD)  
- `end_date` (string, optional): End date (YYYY-MM-DD)

**Examples:**

```bash
# Paginated province cases (default: 50 records)
GET /provinces/cases

# Custom pagination
GET /provinces/cases?limit=100&offset=200

# All data (for charts/analytics)
GET /provinces/cases?all=true

# Specific province with pagination
GET /provinces/31/cases?limit=30

# Specific province, all data
GET /provinces/31/cases?all=true

# Date range with pagination
GET /provinces/cases?start_date=2024-01-01&end_date=2024-01-31&limit=100

# Date range, all data (for time series charts)
GET /provinces/cases?start_date=2024-01-01&end_date=2024-01-31&all=true
```

**Response Structure:**

*Paginated Response:*
```json
{
  "status": "success",
  "data": {
    "data": [
      { "day": 1, "date": "2024-01-15", ... },
      { "day": 2, "date": "2024-01-14", ... }
    ],
    "pagination": {
      "limit": 50,
      "offset": 0,
      "total": 1000,
      "total_pages": 20,
      "page": 1,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

*All Data Response:*
```json
{
  "status": "success", 
  "data": [
    { "day": 1, "date": "2024-01-15", ... },
    { "day": 2, "date": "2024-01-14", ... }
  ]
}
```

## Usage Patterns

### 1. Efficient Data Loading (Default)
```javascript
// Load first page with 50 records
const response = await fetch('/api/v1/provinces/cases');
const { data, pagination } = response.data;

// Load next page
if (pagination.has_next) {
  const nextPage = await fetch(`/api/v1/provinces/cases?offset=${pagination.offset + pagination.limit}`);
}
```

### 2. Charts & Analytics
```javascript
// Get complete dataset for time series chart
const response = await fetch('/api/v1/provinces/cases?all=true&start_date=2024-01-01&end_date=2024-12-31');
const allData = response.data;

// Perfect for Chart.js, D3.js, etc.
const chartData = allData.map(item => ({
  x: item.date,
  y: item.cumulative.positive
}));
```

### 3. Province-Specific Analysis
```javascript
// Get all Jakarta data for detailed analysis
const response = await fetch('/api/v1/provinces/31/cases?all=true');
const jakartaData = response.data;
```

## Error Handling

### Common HTTP Status Codes
- `200` - Success
- `400` - Bad Request (invalid parameters)
- `404` - Not Found
- `500` - Internal Server Error
- `503` - Service Unavailable (database issues)

### Error Response Examples
```json
{
  "status": "error",
  "message": "Invalid date format. Use YYYY-MM-DD"
}
```

```json
{
  "status": "error", 
  "message": "Province not found"
}
```

## Rate Limiting

- No rate limiting currently implemented
- Consider implementing if needed for production use

## CORS

CORS is enabled for all origins to support web applications.

## Data Sources

- Data is sourced from official Indonesian health authorities
- Updates are typically daily
- Historical data available from the beginning of the pandemic

## Best Practices

1. **Use pagination by default** for better performance
2. **Use `all=true` only for analytics/charts** to avoid large payloads  
3. **Implement client-side caching** for frequently accessed data
4. **Use date ranges** to limit data scope when possible
5. **Handle errors gracefully** and implement retry logic
6. **Monitor response times** and adjust pagination limits as needed

## Changelog

### Version 2.0.2
- ✅ Enhanced ODP/PDP data grouping structure
- ✅ Implemented hybrid pagination system
- ✅ Added provinces with latest case data endpoint
- ✅ Improved API response structure
- ✅ Added comprehensive pagination metadata

### Version 2.0.1  
- Fixed database column typos
- Fixed RT display issues

### Version 2.0.0
- Major API restructure
- Added province-level data
- Enhanced statistics calculations