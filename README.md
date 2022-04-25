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

# Reference material
- https://angular.io/cli
- https://angular.io/guide/forms
- https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
- https://www.youtube.com/watch?v=Y92WWaZJl24
- https://intuting.medium.com/using-protocol-buffers-in-angular-96b4f8ab18d4
- https://cloud.google.com/spanner/docs/getting-started/go
