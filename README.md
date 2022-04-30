# Hexago

Practice of hexagonal architecture tuned for Go.

## Considerations

- [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
  + `pkg/controller`: Primary adapters
  + `pkg/infrastructure`: Secondary adapters
  + `pkg/core`: Cores
- Dependency Injection
- [CQRS with Event Sourcing](https://docs.microsoft.com/en-us/azure/architecture/patterns/event-sourcing)
- [Go standard layout](https://github.com/golang-standards/project-layout)

## Modules

### [Hexago Gateway](./gateway)

API gateway module for hexago service.

## Infrastructure

```bash
# Start infrastructures using docker-compose
make infra
```
