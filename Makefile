run:
	go run cmd/main.go

db: postgres-image postgres-run wait-for-postgres migrate-up

# pulls postgres image
postgres-image:
	docker pull postgres:16-alpine

# runs postgres within docker container, allows connection via port :5432
postgres-run:
	docker run \
	--name ranking_container \
	-p 5432:5432 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=secret \
	-e POSTGRES_DB=ranking_db \
	-v pgdata:/var/lib/postgresql/data \
	-d postgres:16-alpine \

# start the docker container which contains the postgresql image
start-db:
	-docker start ranking_container

# gracefully shut down docker container which allows data to persists in postgresql
stop-db:
	-docker stop ranking_container

# waits for postgres container to be ready - without this, migrate will run immediately and fails
wait-for-postgres:
	@echo "Waiting for PostgreSQL to start..."
	@sleep 3;
	@until docker exec ranking_container pg_isready -U root -h localhost -p 5432; do \
		sleep 5; \
	done
	@echo "PostgreSQL is ready"

# uses goose to migrate data into postgres database
migrate-up: wait-for-postgres
	goose -dir db postgres "postgresql://root:secret@localhost:5432/ranking_container?sslmode=disable" up;

# delete and create a new container - resets the database to default (for testing purposes)
set-db: stop-db remove-db remove-volume db