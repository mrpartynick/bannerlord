init:
	docker network create bannerlord
	make postgres
	make

postgres:
	docker run --name postgr -e POSTGRES_USER=postgre -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres:16
	docker network connect bannerlord postgr
createdb:
	docker exec -it postgr createdb --username=postgre --owner=postgre bannerlord
dropdb:
	docker exec -it postgr dropdb bannerlord
migrateup:
	migrate -path migrations -database "postgresql://postgre:1234@localhost:5432/bannerlord?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://postgre:1234@localhost:5432/bannerlord?sslmode=disable" -verbose down
run:
	docker build -t bannerlord:latest .
	docker run --name bann --network bannerlord -p 8080:8080 bannerlord:latest

.PHONY: postgres createdb dropdb