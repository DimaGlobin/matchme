postgresinit:
	docker run --name matchme_db -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=matchme_db -p 5436:5432 -d postgres 

postgresrm:
	docker rm -f matchme_db

run:
	go run cmd/matchme/main.go

migrate-db-up:
	migrate -path ./internal/schema/migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate-db-down:
	migrate -path ./internal/schema/migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

.PHONY: postgresinit, postgresrm, run, migrate-db-up, migrate-db-down 