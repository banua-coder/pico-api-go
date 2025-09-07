# Breaking Changes - National COVID-19 API Response Structure

## Version: 2.0.0
## Date: 2025-09-07

### Overview
The JSON response structure for national COVID-19 endpoints has been completely redesigned to provide a more organized and intuitive data format. All keys are now in English, and related data is grouped into logical nested structures.

### Affected Endpoints
- `GET /api/national`
- `GET /api/national/latest`
- `GET /api/national?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`

### Migration Guide

#### Old Structure (v1.x)
```json
{
  "id": 1,
  "day": 1,
  "date": "2020-03-01T17:00:00Z",
  "positive": 2,
  "recovered": 0,
  "deceased": 0,
  "cumulative_positive": 2,
  "cumulative_recovered": 0,
  "cumulative_deceased": 0,
  "rt": 1.2,
  "rt_upper": 1.5,
  "rt_lower": 0.9
}
```

#### New Structure (v2.0)
```json
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
```

### Key Changes

1. **Removed Fields**
   - `id` - Internal database ID no longer exposed

2. **Restructured Fields**
   - Daily new cases moved to `daily` object
   - Cumulative totals moved to `cumulative` object
   - Reproduction rate (Rt) values moved to `statistics.reproduction_rate`
   - Added calculated percentages in `statistics.percentages`

3. **New Fields**
   - `daily.active` - Daily active cases (calculated as positive - recovered - deceased)
   - `cumulative.active` - Total active cases
   - `statistics.percentages` - Percentage distribution of cases

4. **Field Mapping**
   | Old Field | New Field |
   |-----------|-----------|
   | `positive` | `daily.positive` |
   | `recovered` | `daily.recovered` |
   | `deceased` | `daily.deceased` |
   | `cumulative_positive` | `cumulative.positive` |
   | `cumulative_recovered` | `cumulative.recovered` |
   | `cumulative_deceased` | `cumulative.deceased` |
   | `rt` | `statistics.reproduction_rate.value` |
   | `rt_upper` | `statistics.reproduction_rate.upper_bound` |
   | `rt_lower` | `statistics.reproduction_rate.lower_bound` |

### Client Migration Example

#### JavaScript/TypeScript
```javascript
// Old way
const activeCases = data.cumulative_positive - data.cumulative_recovered - data.cumulative_deceased;
const rtValue = data.rt;

// New way
const activeCases = data.cumulative.active; // Already calculated
const rtValue = data.statistics.reproduction_rate?.value;
```

#### Go
```go
// Old way
type OldResponse struct {
    CumulativePositive int64   `json:"cumulative_positive"`
    Rt                *float64 `json:"rt"`
}

// New way
type NewResponse struct {
    Cumulative struct {
        Positive int64 `json:"positive"`
        Active   int64 `json:"active"`
    } `json:"cumulative"`
    Statistics struct {
        ReproductionRate *struct {
            Value float64 `json:"value"`
        } `json:"reproduction_rate,omitempty"`
    } `json:"statistics"`
}
```

### Benefits of the New Structure

1. **Better Organization**: Related data is grouped together logically
2. **Calculated Fields**: Active cases and percentages are pre-calculated
3. **Clearer Naming**: English keys are more universally understood
4. **Nested Structure**: Easier to access related data without parsing flat fields
5. **Future-Proof**: Structure allows for easy addition of new statistics

### Backwards Compatibility

This is a **breaking change** and is not backwards compatible. Clients will need to update their code to handle the new response structure. Consider implementing API versioning if you need to support both old and new formats simultaneously.