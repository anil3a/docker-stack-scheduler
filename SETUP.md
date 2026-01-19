# Configurations

## Setup initial go modules

```
docker compose -f docker-compose.dev.yml run --rm dev
```

You now have a shell inside container:
```
/app # pwd
/app # go version
go version go1.25 linux/amd64
```

Then run:

### Only inside container
```
go mod init stack-scheduler
go get github.com/robfig/cron/v3
go get github.com/docker/docker/client
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get modernc.org/sqlite
```

go.mod and go.sum are created inside the mounted folder, visible on host.

## Running the development server

```
docker compose -f docker-compose.stage.yml up
```

The server will be accessible at `http://localhost:8044`.

## Database file

The SQLite database file will be created inside the `data/` folder in the project root. Make sure this folder exists and is writable.

## Stopping the development server
To stop the server, press `CTRL+C` in the terminal where the Docker Compose command is running. This will gracefully shut down the server and any running containers.

## Notes
- Make sure you have Docker and Docker Compose installed on your machine.
- You can modify the `docker-compose.stage.yml` file to change environment variables or other settings as needed.


