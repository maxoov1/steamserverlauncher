# Steam Server Launcher

## Configuration

```sh
LAUNCHER_TEMPLATES_PATTERN = "templates/*.html"
LAUNCHER_ADDRESS = "127.0.0.1:8000"
LAUNCHER_SERVER_NAME = "Default Source Server"
LAUNCHER_SERVER_ADDRESS = "127.0.0.1:27015"
```

## Run

### Local
```sh
go mod download
CGO_ENABLED=0 go build -o steamserverlauncher ./cmd/main.go

./steamserverlauncher
```

### Docker
```sh
docker build -t steamserverlauncher .
docker run -d steamserverlauncher # don't forget about env
```