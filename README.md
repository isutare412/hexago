# Hexago

System architecture PoC tuned for Go. The name Hexago was taken from Hexagonal
Architecture.

## Considerations

- [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
  + `pkg/controller`: Driving adapters
  + `pkg/infrastructure`: Driven adapters
  + `pkg/core`: Business logics
- [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)
- [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection)
- [CQRS with Event Sourcing](https://docs.microsoft.com/en-us/azure/architecture/patterns/event-sourcing)
- [Go Standard Layout](https://github.com/golang-standards/project-layout)

## Overview

![Hexago CQRS diagram](assets/hexago-cqrs.drawio.png?raw=true)

## Modules

### [Hexago Gateway](./gateway)

API gateway module for hexago service.

### [Hexago Payment](./payment)

Payment module for hexago service.

### [Hexago Common](./common)

Includes common libraries of hexago modules.

## Infrastructure

1. Run infrastructures using docker compose.

```bash
make infra
# 1) up
# 2) down
# 3) ps
# 4) logs
# 5) quit
# Please enter your choice: 1
```

2. Add local dns to each insfrastructure for development.

```bash
sudo tee -a /etc/hosts > /dev/null <<EOT
127.0.0.1 mongodb
127.0.0.1 kafka1
127.0.0.1 kafka2
EOT
```

3. Create kafka topics using [kafka-ui](http://localhost:58080).
    - `donation-request`
