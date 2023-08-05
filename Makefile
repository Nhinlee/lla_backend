postgres:
	docker run --name llapp -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nhin123456 -d postgres:14.8-alpine

createdb:
	docker exec -it llapp createdb --username=root --owner=root lla

dropdb:
	docker exec -it llapp dropdb lla

migrate-up:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/lla?sslmode=disable" -verbose up
	
migrate-up1:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/lla?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/lla?sslmode=disable" -verbose down

migrate-down1:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/lla?sslmode=disable" -verbose down 1

run-local-server:
	IS_LOCAL=true go run main.go