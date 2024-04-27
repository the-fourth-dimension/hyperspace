# hyperspace

the medium, the server

## Running locally

```sh
# setup env vars
cp .env.example .env

# run the hyperspace dev server
go run cmd/main.go

# build the hyperspace server
go build -o hyperspace-bin cmd/main.go

# run the build
./hyperspace-bin
```
