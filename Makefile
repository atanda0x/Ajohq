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

.PHONY: postgres createdb dropdb startdoc stopdoc