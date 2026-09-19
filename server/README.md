```
server/ # Root folder of the individual repo
├── api/ # OpenAPI/Swagger specs or Protobuf files
│ └── proto/
│ └── order.proto
│
├── cmd/ # Main entry point directory
│ └── order-api/
│ └── main.go # Application bootstrapper
│
├── configs/ # Configuration files (YAML, env vars)
│ └── config.yaml
│
├── internal/ # Private module code; cannot be imported by other services
│ ├── app/ # Application lifecycle and orchestration
│ ├── domain/ # Business domain entities and core logic
│ ├── handler/ # Adapters: REST controllers or gRPC servers
│ ├── service/ # Core use cases / orchestrators
│ └── repository/ # Data adapters (SQL, Mongo, Redis)
│
├── pkg/ # Code that is safe to be imported by other projects
│ └── client/ # Generated gRPC client SDK for other services to use
│
├── Dockerfile # Containerization script
├── Makefile # Tooling tasks (build, test, lint, proto generation)
├── go.mod
└── go.sum
```
