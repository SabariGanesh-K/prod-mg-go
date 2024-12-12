DB_URL=postgresql://root:secret@localhost:5432/prod-mgm?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createmigrateinitschema:
	 migrate create -ext sql --dir db/migrations --seq init_schema  

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

createdb:
	docker exec -it postgres createdb --username=root --owner=root prod-mgm

dropdb:
	docker exec -it postgres dropdb prod-mgm
sqlc:
	sqlc generate
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/SabariGanesh-K/prod-mgm-go/db/sqlc Store

testfull:
	go test -v -cover ./db/sqlc ./api ./token  
testmail:
	go test -v -cover ./email

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

rabbitmq:
	docker-compose up -d

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
test:
	go test -v -cover  ./api
	
launch:
	docker start postgres
	docker start redis
.PHONY: postgres new_migration createdb dropdb migrateup migratedown createmigrateinitschema sqlc mock test  redis testmail launch test rabbitmq