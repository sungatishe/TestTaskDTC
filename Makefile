createdb:
	docker exec -it order_db createdb --username=db_user --owner=db_user order_db

dropdb:
	docker exec -it order_db dropdb --username=db_user order_db

migrateup:
	migrate -path db/migrations -database "postgresql://db_user:db_password@localhost:5432/order_db?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgresql://db_user:db_password@localhost:5432/order_db?sslmode=disable" down

build:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down


.PHONY: createdb dropdb migrateup migratedown build up down
