# Генерация API кода
gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# Запуск приложения
run:
	go run cmd/app/main.go

# Применение миграций
migrate-up:
	migrate -path ./migrations -database "postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable" up

# Создание новой миграции
migrate-new:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-new NAME=<migration_name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir ./migrations -format "20060102150405" -seq $(NAME)

# Проверка кода
lint:
	golangci-lint run --out-format=colored-line-number
