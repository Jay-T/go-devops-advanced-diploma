postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret  -d postgres:12.3-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root gophkeeper

dropdb:
	docker exec -it postgres12 dropdb gophkeeper

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/gophkeeper?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/gophkeeper?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/Jay-T/go-devops-advanced-diploma/db/sqlc Store

test:
	go test -v -cover ./...

proto:
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
	--go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	internal/proto/*.proto 

server-run:
	ENVIRONMENT=development go run ./cmd/server/. -c ./cmd/server/config.json

client-run:
	ENVIRONMENT=development go run ./cmd/client/.

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mockgen test client-run server-run