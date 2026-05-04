callback-server:
	@go run cmd/callback/main.go


build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o spotify-in-github cmd/app/main.go
