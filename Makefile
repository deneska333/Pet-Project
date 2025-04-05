gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

run:
	go run cmd/app/main.go

migrate-up:
	migrate -path ./migrations -database "postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable" up

lint:
	golangci-lint run --out-format=colored-line-number