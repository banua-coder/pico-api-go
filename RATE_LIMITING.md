# Rate Limiting

This API implements rate limiting to ensure fair usage and protect against abuse. The rate limiter uses a sliding window algorithm to track requests per client IP address.

## Configuration

Rate limiting can be configured using environment variables:

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `RATE_LIMIT_ENABLED` | `true` | Enable or disable rate limiting |
| `RATE_LIMIT_REQUESTS_PER_MINUTE` | `100` | Maximum requests per minute per IP |
| `RATE_LIMIT_BURST_SIZE` | `20` | Burst size for initial requests |
| `RATE_LIMIT_WINDOW_SIZE` | `1m` | Time window for rate limiting |

## Response Headers

All API responses include the following rate limiting headers:

- `X-RateLimit-Limit`: The maximum number of requests allowed in the current window
- `X-RateLimit-Remaining`: The number of requests remaining in the current window
- `X-RateLimit-Reset`: Unix timestamp when the rate limit window resets (only on 429 responses)
- `Retry-After`: Number of seconds to wait before making another request (only on 429 responses)

## Rate Limit Exceeded

When the rate limit is exceeded, the API returns:

- **Status Code**: `429 Too Many Requests`
- **Response Body**:
  ```json
  {
    "status": "error",
    "error": "Rate limit exceeded. Too many requests."
  }
  ```

## Client IP Detection

The rate limiter identifies clients by IP address using the following priority:

1. `X-Forwarded-For` header (for load balancers/proxies)
2. `X-Real-IP` header (for reverse proxies)
3. `RemoteAddr` from the connection (fallback)

## Implementation Details

- **Algorithm**: Sliding window rate limiter
- **Storage**: In-memory (per instance)
- **Cleanup**: Automatic cleanup of old client records every 5 minutes
- **Thread Safety**: Fully concurrent with proper mutex locking

## Best Practices for Clients

1. **Check Headers**: Always check the `X-RateLimit-*` headers to understand your current quota
2. **Handle 429 Responses**: Implement exponential backoff when receiving 429 responses
3. **Use Retry-After**: Respect the `Retry-After` header value before retrying
4. **Distribute Requests**: Avoid bursting all requests at once; distribute them evenly

## Example Usage

```bash
# Check current rate limit status
curl -I https://api.example.com/api/v1/national

# Response headers will include:
# X-RateLimit-Limit: 100
# X-RateLimit-Remaining: 99

# When rate limited:
# HTTP/1.1 429 Too Many Requests
# X-RateLimit-Limit: 100
# X-RateLimit-Remaining: 0
# X-RateLimit-Reset: 1672531200
# Retry-After: 60
```

## Disabling Rate Limiting

To disable rate limiting (not recommended for production):

```bash
export RATE_LIMIT_ENABLED=false
```

Or set it in your `.env` file:

```
RATE_LIMIT_ENABLED=false
```