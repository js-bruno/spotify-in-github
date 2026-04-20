callback-server:
	@go run cmd/callback/main.go


build:
	@go build cmd/app/main.go
	@mv ./.env.local ~/.
	@mv ./main ~/.
