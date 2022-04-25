# Requirements
- Go 
- Protobuf compiler
- Angular
- Docker
- Valid gcloud credentials in `backend/gcloud-credentials.json` with access to gcloud project `prj-nxh-moapr-spanner-dev-8104`
    - Run `gcloud auth application-default login`, note location of `application_default_credentials.json`
    - Run `gcloud config set project prj-nxh-moapr-spanner-dev-8104`
    - Copy `application_default_credentials.json` to `backend/gcloud-credentials.json`

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
