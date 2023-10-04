run-db:
	 docker run --name=BWG -e POSTGRES_PASSWORD='12345' -p 5432:5432 -d --rm postgres

migrations-up:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

migrations-down:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' down

run-cache:
	docker run --name=redis -e REDIS_HOST=redis -e REDIS_PORT=6379 -p 6379:6379 -d --rm redis

run-server:
	go run cmd/main.go