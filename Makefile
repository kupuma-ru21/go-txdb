run:
	go run server.go

generate:
	gqlgen generate
	sqlboiler mysql
	go generate ./...


testing:
	go test ./test