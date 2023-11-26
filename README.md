# Snapp Food Task Services
**Snapp Food Task** Services consists of several microservices that are connected to each other by **RabbitMQ**. the database used in the snapp-food-task services is **Postgresql**, also, order delay check requests are stored in **Redis**.

## Requirements
- [Golang](https://go.dev/dl/) (1.18 or higher)
- [Docker](https://docker.io/)
- [Make](https://en.wikipedia.org/wiki/Make_(software))

## Getting Started
To run the snapp-food-task services, after ensuring `docker` and `docker compose` are installed, enter the following command:
```bash
make up
```
## How to use
You can import **[Postman](./InfoService.postman_collection.json)** collection, after ensuring start services and send requests.

## Routes

| Route                                  | Method | Header | Body | Response                         |
|----------------------------------------|--------|--------|------|----------------------------------|
| {infoService}/order/append             | GET    | -      | -    | Order                            |
| {infoService}/delay/report/:orderID    | GET    | -      | -    | Delay report / New estimate time |
| {infoService}/delay/fetch              | GET    | -      | -    | Report result                    |
| {infoService}/delay/fetch              | GET    | -      | -    | Report result                    |
| {supportService}/agent/add/:agentID    | GET    | -      | -    |                                  |
| {supportService}/agent/remove/:agentID | GET    | -      | -    |                                  |

## How to Build
Make sure you have installed golang (1.18 or higher). to check your golang version run `go version`. Also, you need `make` to build snapp-food-task services
### Compile the source
you can compile snapp-food-task services by following these steps:
```bash
make go-build
```

### Build docker files
you can build docker files by following these steps:
```bash
make docker-build
```

