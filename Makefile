gen-tasks:
	oapi-codegen -config openapi/.openapi.tasks -include-tags tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi.users -include-tags users openapi/openapi.yaml > ./internal/web/users/api.gen.go

gen: gen-tasks gen-users

run:
	go run cmd/app/main.go

migrate-up:
	migrate -path ./migrations -database "postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable" up

lint:
	golangci-lint run --out-format=colored-line-number
