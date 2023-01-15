postgres:
	sudo docker run --name picsaverdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
	
createdb:
	sudo docker exec -it picsaverdb createdb --username=root --owner=root picsaverdb

dropdb:
	sudo docker exec -it picsaverdb dropdb picsaverdb

migrateup:
	migrate -path db/schema -database "postgresql://root:secret@localhost:5432/picsaverdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/schema -database "postgresql://root:secret@localhost:5432/picsaverdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server 
