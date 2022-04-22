# Requirements
- Go 
- Protobuf compiler
- Angular
- Docker

# Running

```sh
cd backend
./proto-generate.sh
go run main.go
```

```sh
cd frontend
./proto-generate.sh
ng serve
```

```sh
cd frontend
./proxy.sh
```

Browse to `http://localhost:4200`
