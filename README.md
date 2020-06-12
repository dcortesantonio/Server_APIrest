# Server_APIrest
This repository contains an APIrest that search information of a domain.
## Getting Started

* Language: Go
* Data Base: CockRoach -  Go pq Driver
* API Router: fasthttprouter

### Go
[Install Go](https://golang.org/dl/) - Golang

### Cockroach
1. [Install cockroach](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)
2. [Start a Local Cluster](https://www.cockroachlabs.com/docs/stable/start-a-local-cluster.html)
3. Install Go pq Driver
```
go get -u github.com/lib/pq
```
4. In the SQL shell, [[Folder: DataBase Statements]](https://github.com/dcortesantonio/Server_APIrest/tree/master/DataBase%20Statements)

### Run API
```
go run main.go
```