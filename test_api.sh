#!/bin/bash

# Test script for national COVID-19 API endpoints
# This script tests the new JSON response structure

API_URL="http://localhost:8080"

echo "Testing National COVID-19 API Endpoints"
echo "========================================"
echo ""

# Test 1: Get all national cases
echo "1. Testing GET /api/national"
echo "----------------------------"
curl -s "${API_URL}/api/national" | jq '.[0]' 2>/dev/null || echo "API not running or endpoint not available"
echo ""

# Test 2: Get latest national case
echo "2. Testing GET /api/national/latest"
echo "------------------------------------"
curl -s "${API_URL}/api/national/latest" | jq '.' 2>/dev/null || echo "API not running or endpoint not available"
echo ""

# Test 3: Get national cases with date range
echo "3. Testing GET /api/national with date range"
echo "---------------------------------------------"
curl -s "${API_URL}/api/national?start_date=2020-03-01&end_date=2020-03-10" | jq '.[0]' 2>/dev/null || echo "API not running or endpoint not available"
echo ""

echo "Expected JSON structure:"
echo "========================"
cat << 'EOF'
{
  "day": 1,
  "date": "2020-03-01T17:00:00Z",
  "daily": {
    "positive": 2,
    "recovered": 0,
    "deceased": 0,
    "active": 2
  },
  "cumulative": {
    "positive": 2,
    "recovered": 0,
    "deceased": 0,
    "active": 2
  },
  "statistics": {
    "percentages": {
      "active": 100,
      "recovered": 0,
      "deceased": 0
    },
    "reproduction_rate": {
      "value": 1.2,
      "upper_bound": 1.5,
      "lower_bound": 0.9
    }
  }
}
EOF