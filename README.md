# tinyurl

Petit raccourcisseur d'URL (net/http).

```bash
go run .
go test ./...
```

- `GET /health`
- `POST /shorten` — `{"url": "..."}` renvoie un code
- `GET /{code}` — redirige vers l'URL d'origine
