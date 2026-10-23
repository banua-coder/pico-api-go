# Content endpoints (spec)

Read-only `GET` endpoints under `/api/v1` for site content that already lives in the production MySQL tables. No schema changes, no writes.

## Conventions

- Envelope: `{"status":"success","data":{"data":[...],"pagination":{...}}}` (same as `/hospitals`).
- Pagination: `?page=1&per_page=10` (max 100). `?load_all=true` returns a plain list in `data` with no pagination block.
- Cache: in-memory via `cachedContentService`, TTL 1 hour (`ttlDefault`), one key per filter and page. Cleared by `POST /admin/cache/clear`.
- Image and URL fields are returned exactly as stored. Many `image` values point at the expired `banuacoders.com` host, some are relative paths (`donations/July2021/x.jpg`) and partner images are bare file names. Clients must tolerate broken images.
- Fields are English snake_case. Only columns that matter to clients are exposed (no `created_at` / `updated_at`).

## Endpoints

| Path | Filters | Source tables | Order |
|------|---------|---------------|-------|
| `GET /banners` | `include_inactive=true` (default: `status=1` only, as the old site did) | `banners` | id asc |
| `GET /infographics` | - | `infographics`, `infographic_images` | id desc (old site), images by `queue` |
| `GET /donations` | - | `donations`, `bank_accounts`, `banks`, `contacts`, `contact_types` | id asc |
| `GET /telemedicines` | - | `telemedicines`, `telemedicine_schedules`, `contacts`, `contact_types` | id asc |
| `GET /partners` | `in_footer=true\|false` | `partners` | id asc |
| `GET /contacts` | `owner_type=hospital\|task_force\|donation\|telemedicine`, `owner_id=<int>` (400 on unknown owner_type) | `contacts`, `contact_types` | id asc |

No `?province=` filter: none of these tables has a province column.

Not added (already covered): hospital beds and bed types are returned inside `GET /hospitals` and `GET /hospitals/{code}` (`beds[]` with `bed_type_name`). Contact types are embedded in every contact, so there is no separate `/contact-types`.

## Response shapes

Banner (`link` = `url_type` + `url`, because the table splits the URL prefix from the rest):
```json
{"id":1,"title":"Panduan Isolasi Mandiri","link":"https://www.instagram.com/p/CR5q6ObHWJP","image":"https://...png","is_active":true}
```

Infographic:
```json
{"id":146,"title":"...","source":"https://www.instagram.com/p/...","images":[{"id":446,"queue":1,"image_url":"https://i.ibb.co/..."}]}
```

Donation (`delivery_address` is `null` when unset, see notes):
```json
{"id":2,"title":"...","initiator":"@x","url":"https://...","description":"...","image":"donations/July2021/x.jpg",
 "accept_noncash":true,"delivery_address":null,
 "bank_accounts":[{"id":1,"account_number":"7921024685","account_name":"Yaumil","bank":{"id":3,"name":"Bank Central Asia, Tbk","nickname":"BCA","logo":"https://..."}}],
 "contacts":[{"id":89,"contact_type":{"id":2,"name":"Telepon","prefix":"tel:","icon":"fas fa-phone","bg_color":"bg-blue-500","label":"Telpon"},"contact":"0812..."}]}
```

Telemedicine:
```json
{"id":1,"name":"dr. Evans Suryani","is_free":true,
 "schedules":[{"id":1,"operational_day":"Senin-Jumat","operational_time":"07:00-20:00","timezone":"WIB"}],
 "contacts":[{"id":56,"contact_type":{...},"contact":"evansevansevans"}]}
```

Partner:
```json
{"id":1,"name":"Detexi","link":"https://detexi.id","image":"detexi.png","in_footer":true,"show_anywhere":false}
```

Contact (from `/contacts`; nested contacts omit `owner_type` / `owner_id`). A clickable link is `contact_type.prefix + contact`:
```json
{"id":1,"owner_type":"hospital","owner_id":10,"contact_type":{...},"contact":"045022251"}
```

## Notes on the data (checked read-only against production)

- `contacts` is polymorphic: `contactable_type` is `App\Models\{Hospital,TaskForce,Donation,Telemedicine}`; mapped to the short `owner_type` names above. Same for `bank_accounts.accountable_type` (only `Donation` in use).
- `donations.delivery_address` holds the literal string `'NULL'` on 5 of 8 rows; the query uses `NULLIF(..., 'NULL')` so the API returns JSON `null`. The old Laravel transformer returned `''` for null.
- Banners: 7 rows, 3 active (ids 1, 2, 8); ids 3 and 7 do not exist.
- Infographics: 146 rows, 448 images, queue 1..10, no infographic without images, no orphan images. Donations: 8 rows (ids 2..9), 9 bank accounts over 7 donations, 5 banks. Telemedicines: 33 rows, 47 schedules (at least one each). Partners: 12 rows, 9 with `in_footer=1`. Contacts: 167 rows (hospital 55, task_force 50, donation 28, telemedicine 34), no orphans, no duplicates.
- Children are loaded with one `IN (...)` query per relation, not per row.
