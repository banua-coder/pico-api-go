# API Documentation

This directory contains auto-generated OpenAPI/Swagger documentation for the Sulawesi Tengah COVID-19 Data API.

## Files

- `swagger.yaml` - OpenAPI 3.0 specification in YAML format
- `swagger.json` - OpenAPI 3.0 specification in JSON format  
- `docs.go` - Generated Go file containing the documentation

## Usage

### Interactive Documentation

When the API server is running, you can access the interactive Swagger UI at:
- **Local development**: http://localhost:8080/swagger/index.html
- **Production**: https://pico-api.banuacoder.com/swagger/index.html

### Swagger Specification Files

You can also use the specification files directly with various tools:

```bash
# Validate the OpenAPI spec
swagger-codegen validate -i docs/swagger.yaml

# Generate client SDKs
swagger-codegen generate -i docs/swagger.yaml -l javascript -o clients/js
swagger-codegen generate -i docs/swagger.yaml -l python -o clients/python

# Import into Postman, Insomnia, or other API tools
# Use the swagger.json or swagger.yaml file
```

## Regenerating Documentation

To regenerate the documentation after code changes:

```bash
swag init -g cmd/main.go -o ./docs
```

This will update all files in this directory based on the Go code annotations.

## Key Features Documented

- 🏥 **Health Check** - API status and database connectivity
- 🏛️ **Sulawesi Tengah Focus** - Primary COVID-19 data for Central Sulawesi
- 🇮🇩 **National Data** - COVID-19 statistics for Indonesia (reference context)
- 🗺️ **Province Data** - Provincial COVID-19 information with latest case data by default
- 📊 **Province Cases** - Detailed case data with hybrid pagination support
- 📄 **Pagination** - Comprehensive pagination metadata and flexible data retrieval
- 🏷️ **Enhanced Data Structure** - Grouped ODP/PDP fields with proper daily/cumulative separation

## Response Models

All response models include proper JSON schema definitions with:
- Type validation
- Required field specifications  
- Example values
- Field descriptions
- Nested object relationships

This makes it easy to generate client code, validate API responses, and understand the data structure.