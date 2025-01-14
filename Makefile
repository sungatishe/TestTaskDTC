createdb:
	docker exec -it order_db createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it order_db dropdb --username=${POSTGRES_USER} ${POSTGRES_DB}

migrateup:
	migrate -path db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" down

build:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

.PHONY: createdb dropdb migrateup migratedown build up down
