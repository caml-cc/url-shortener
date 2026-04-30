# url-shortener

A lightweight self-hosted URL shortener written in Go.

Current implementation details:

- HTTP server built with Gorilla Mux
- SQLite storage (file at `internal/database/url-shortener.db`)
- Automatic SQL migrations on startup
- API-key protected create/delete endpoints
- Plain-text request bodies

## Requirements

- Go
- SQLite support via `github.com/mattn/go-sqlite3` (CGO-enabled build environment)

## Configuration

The app requires a `.env` file in the project root.

```env
PORT=5099
API_KEY=abc123
RAND_CHARS=6 # Optional, defaults to 8 when missing.
```

If either value is missing, startup fails.

## Run

Using make:

```bash
make run
```

Or directly:

```bash
go run ./cmd/url-shortener/main.go
```

Build binary:

```bash
make build
./bin/url-shortener
```

## API

Base URL examples below use `http://localhost:5099`.

### Create short URL

`POST /create`

Headers:

- `K: <API_KEY>`

Body format (plain text):

- `url`
- `url|alias`

Examples:

```bash
curl -H "K: <API_KEY>" http://localhost:5099/create -d 'example.com'
```

```bash
curl -H "K: <API_KEY>" http://localhost:5099/create -d 'https://example.com/docs|docs'
```

Behavior:

- If alias is omitted, an 8-character random alias is generated
- URLs without scheme are normalized to `http://...`
- Only `http` and `https` URLs are accepted

Success response:

- Status: `201 Created`
- Body: `<host>/<alias>`

Possible errors:

- `400 Bad Request` for invalid body/URL
- `401 Unauthorized` for missing or invalid API key
- `500 Internal Server Error` (for example, duplicate alias insertion)

### Delete short URL

`POST /delete`

Headers:

- `K: <API_KEY>`

Body format (plain text):

- `alias`

Example:

```bash
curl -H "K: <API_KEY>" http://localhost:5099/delete -d 'docs'
```

Success response:

- Status: `200 OK`
- Body: `<alias> deleted`

Possible errors:

- `401 Unauthorized` for missing or invalid API key
- `500 Internal Server Error` if alias does not exist or delete fails

### Redirect

`GET /{alias}`

Example:

```bash
curl -i http://localhost:5099/docs
```

Behavior:

- Redirects using `301 Permanent Redirect`
- Returns `404` if alias is not found

## Storage and Migrations

- Database file: `internal/database/url-shortener.db`
- Migrations dir: `internal/database/migrations`
- Migrations run automatically when the server starts

## Project Status

Current state is a working core service with create, delete, and redirect endpoints backed by SQLite.

Known limitations in current behavior:

- No structured JSON API
- Duplicate alias attempts currently surface as generic http 500 errors
- Delete for missing alias currently returns 500 rather than 404

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
