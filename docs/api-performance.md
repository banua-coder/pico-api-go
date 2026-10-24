# API performance and compatibility notes

Beads: pico-ko0.6 (gzip and cache headers), pico-ko0.7 (A2 fixes), pico-ko0.8 (latest case per regency). All changes are additive.

## Transport: gzip, Cache-Control, ETag (pico-ko0.6)

Applies to `GET`/`HEAD` under `/api/v1` only. Swagger and `/admin/*` are untouched.
Middleware order: Recovery, Logging, RateLimit, CORS, Gzip, Cache.

| Behaviour | Detail |
|---|---|
| gzip | Sent when `Accept-Encoding` allows gzip (`gzip;q=0` is honoured). Bodies under 1 KB and responses that already set `Content-Encoding` stay as they are. `Vary: Accept-Encoding` is always set. |
| Cache-Control | 200 responses: `public, max-age=300, stale-while-revalidate=600`. `/api/v1/health`: `no-store`. Errors carry no cache header. |
| ETag | Weak ETag (`W/"<sha256 prefix>"`) of the uncompressed body. A matching `If-None-Match` returns `304` with no body. |

Config (Go durations, see `.env.example`):

- `HTTP_CACHE_MAX_AGE` (default `5m`). `0` sends `no-cache`, so clients revalidate through the ETag.
- `HTTP_CACHE_STALE_WHILE_REVALIDATE` (default `10m`). `0` omits the directive.

Note: `POST /admin/cache/clear` clears the server cache only. Browsers and CDNs may serve a response for up to `max-age` plus `stale-while-revalidate` (15 minutes with the defaults) after a data change. If you front the API with a CDN, set a short `s-maxage` there or purge on clear. Lower `HTTP_CACHE_STALE_WHILE_REVALIDATE` to shrink the window.

## Fixes (pico-ko0.7)

| Change | Before | After |
|---|---|---|
| `sort` on `/provinces/{id}/cases` | Ignored. Always date descending | Honoured (`sort=date:asc`, `positive:desc`, ...). With no `sort`, still date descending |
| `end_date` | Compared as midnight, so rows stamped later that day were dropped | Inclusive of the whole day (`end_date=2021-07-01` includes `2021-07-01T16:14:07`). Applies to `/national` and `/provinces[/{id}]/cases` |
| Invalid `start_date` / `end_date` | Ignored for `/provinces*`, `500` for `/national` | `400` with `Invalid <param>: expected format YYYY-MM-DD` |
| Unknown province in `/provinces/{id}/cases` | `200` with an empty list | `404` |

Not changed: `/provinces/cases` (all provinces) with a date range still ignores `sort`. Its default order differs from the single-province one, so honouring `sort` there needs its own decision.

### Skipped, with reasons

- **ODP/PDP `null` for provinces after the source stopped.** It changes a field type from number to nullable. That breaks clients that treat the field as a number. It also needs a per-province "last non-zero date" rule. Needs a versioned or opt-in change.
- **`date` on `/stats/gender*`.** `province_gender_cases.day` has no reliable join key. It is not unique (809 rows, 797 distinct values, duplicates such as 613 and 625, gaps at 656 and 733). It lines up with `national_cases.day` for the latest row (799 is 2022-05-09) but with `province_cases.day`/`national_cases.id` one day earlier (2022-05-08). The two keys diverge after the out-of-order inserts, so any derived date can be wrong by a day or more. Fix the data first (pico-0pk.6, day renumbering), then add `date`.

## Latest case per regency (pico-ko0.8)

`GET /api/v1/regencies?include_latest_case=true` (works with `load_all=true` and with pagination) adds `latest_case` to each regency, mirroring `/provinces`. It is the flat regency case row (daily `positive`/`recovered`/`deceased`, `cumulative_*`, `rt`, `rt_upper`, `rt_lower`, `date`), or `null` when a regency has none. The latest row is the one with the highest `day` that has any non-zero daily count. Served from the cached `GetLatestRegencyCases`, which now covers every regency that has `regency_cases` rows, in any province (earlier it joined only province 72, Sulawesi Tengah; other provinces got `null`).

## Review follow-ups (pico-ko0.14)

| Change | Detail |
|---|---|
| `Accept-Encoding` | Parsed per RFC 9110: `q` read with `ParseFloat` (`q=0`, `q=0.0`, `q=0.000` refuse gzip), coding names are case-insensitive, `*` allows gzip unless an explicit `gzip` entry says otherwise, an unparsable `q` refuses gzip. |
| Paging order | Province case queries end `ORDER BY ... , pc.id DESC`, so sorts on non-unique columns (`positive`, `province_id`, ...) page without repeats or gaps. |
| Province 404 check | `GetProvinceByID` is cached for found provinces only. Unknown ids are never stored, so random ids cannot grow the cache. |
| Rate-limit headers on `304` | `X-RateLimit-*` are removed from `304` responses. They describe the request quota at the time of the `200` and would be stale. They are advisory on `200` responses. A `304` still counts against the quota. |
