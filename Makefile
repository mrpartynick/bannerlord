init:
	docker network create bannerlord
	make postgres
	make builtapp

postgres:
	docker run --name postgr --network bannerlord -e POSTGRES_USER=postgre -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres:16
createdb:
	docker exec -it postgr createdb --username=postgre --owner=postgre bannerlord
dropdb:
	docker exec -it postgr dropdb bannerlord
migrateup:
	migrate -path migrations -database "postgresql://postgre:1234@localhost:5432/bannerlord?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://postgre:1234@localhost:5432/bannerlord?sslmode=disable" -verbose down
stoppostgres:
	docker stop postgr

built:
	docker build -t bannerlord:latest .
runserver:
	docker run --name bann --network bannerlord -p 8080:8080 bannerlord:latest
stopserver:
	docker stop bann

deinit:
	docker rm bann
	docker network remove bannerlord