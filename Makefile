postgres:
	docker run --name fintechAPI -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ethereum -d postgres:16-alpine

createdb:
	docker exec -it fintechAPI createdb --username=root --owner=root fintechAPI

dropdb: 
	docker exec -it fintechAPI dropdb fintechAPI

startdoc:
	docker start fintechAPI

stopdoc:
	docker stop fintechAPI

migrateup:
	migrate -path db/migration -database "postgresql://root:ethereum@localhost:5432/fintechAPI?sslmode=disable" -verbose up 

migrateup1:
	migrate -path db/migration -database "postgresql://root:ethereum@localhost:5432/fintechAPI?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:ethereum@localhost:5432/fintechAPI?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:ethereum@localhost:5432/fintechAPI?sslmode=disable" -verbose down 1 

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/atanda0x/FintechConnect/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc startdoc stopdoc test server mock