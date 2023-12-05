postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ethereumsolana -d postgres:16-alpine

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ethereumsolana -d mysql

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root Ajohq

dropdb:
	docker exec -it postgres16 dropdb -U root Ajohq

migrateup:
	migrate -path db/migration  -database "postgresql://root:ethereumsolana@localhost:5432/Ajohq?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration  -database "postgresql://root:ethereumsolana@localhost:5432/Ajohq?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:ethereumsolana@localhost:5432/Ajohq?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:ethereumsolana@localhost:5432/Ajohq?sslmode=disable" -verbose down 1

sqlc: 
	docker run --rm -v "${CURDIR}:/src" -w /src kjconroy/sqlc generate

init:
	sqlc init

test:
	go test -v -cover ./...

server:
	go run main.go


mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/atanda0x/Ajohq/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown1 migratedown  sqlc init test server mock 