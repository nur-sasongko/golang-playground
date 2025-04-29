# Scripts
1. Install dependencies.
```bash
go mod tidy 
```

2. Run the app.
```bash
go run .
```

# After running the app, you can use the following endpoints:

- **POST** `http://localhost:8080/shorten` - Create a short URL
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url":"https://google.com"}'
```

Response
```json
{
  "short_url": "http://localhost:8080/aZ1kKq"
}
```


# What You'll Build
- Shorten long URLs like `https://example.com/article/1234` → `http://localhost:8080/abc123`
- Redirect when short URL is visited
- Optional:
    - Custom alias (e.g. `/my-cv`)
    - Expiration time
    - Click tracking

# **Basic features:**
- `POST /shorten` – Create a short URL
- `GET /{shortCode}` – Redirect to original URL