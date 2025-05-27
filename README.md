# Agritrak API

Service API for [AgriTrakPH web app](https://marketmonitor.tbdh.dev/)

Swagger URL: https://api-marketmonitor.tbdh.dev/swagger/index.html

## Development

Requires Go `>= 1.23.2`

- For running during development, use `air` for auto reload.

### Environment Variables

```
MONGODB_URI=mongodb-uri-instance
```

### Deployment

Using docker compose.

```
docker compose up -d --build
```
