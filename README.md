# warden
Self hosted IAM Server 


### Development

To setup this project locally, you need to have Go installed. We are using Go v1.23.4, but any version above v1.21 should work.

Alongside Go, you will need to have docker installed to run the postgres container.

*Setup Go and dependencies*

```bash
go mod download
```

*Run the postgres container*    

```bash
docker compose up -d 
```

*Run the tests*

```bash
go test -v ./...
```

