ifneq (,$(wildcard ./.env))
    include .env
    export
endif


createdb:
	docker exec -it order_db createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it order_db dropdb --username=${DB_USER} ${DB_NAME}

migrateup:
	migrate -path db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" down

build:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

swag_init:
	swag init -d ./cmd/order,./

test:
	go test ./...

create_topic:
	docker exec ${KAFKA_CONTAINER_NAME} kafka-topics --create --topic ${TOPIC_NAME} --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092


.PHONY: createdb dropdb migrateup migratedown build up down swag_init test create_topic
