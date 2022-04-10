# Hexago

Practice of hexagonal architecture tuned for Go.

## Considerations

- [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
  + `pkg/controller`: Primary adapters
  + `pkg/infrastructure`: Secondary adapters
  + `pkg/core`: Cores

- Dependency Injection
  + Compile-time DI using [google wire](https://github.com/google/wire)

- [Go standard layout](https://github.com/golang-standards/project-layout)

## Development

```bash
# Start infrastructures using docker-compose
make infra
```

```bash
# Generate dependency injection code
make wire
```

```bash
# Run hexago server locally
make run-local
```
