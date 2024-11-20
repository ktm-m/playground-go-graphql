gen:
	go run github.com/99designs/gqlgen generate

run:
	go run main.go

run-compose:
	docker compose up -d

stop-compose:
	docker compose down
