postgres:   # Run PostgreSQL container with specified user and password
	docker run --name motionless_DB -p 5432:5432 -e POSTGRES_USER=maalem -e POSTGRES_PASSWORD=secretugtk -d postgres

createdb:   # Create a database named 'Motionless' within the PostgreSQL container
	docker exec -it motionless_DB createdb --username=maalem --owner=maalem Motionless

dropdb:     # Drop the 'Motionless' database from the PostgreSQL container
	docker exec -it motionless_DB dropdb Motionless

migrate_up: # Apply database migrations to the 'Motionless' database
	migrate -path migrations -database "postgresql://maalem:secretugtk@localhost:5432/Motionless?sslmode=disable" -verbose up

migrate_down:   # Roll back the last applied database migration for 'Motionless'
	migrate -path migrations -database "postgresql://maalem:secretugtk@localhost:5432/Motionless?sslmode=disable" -verbose down

sqlc:   # Generate Go code using the SQLC tool
	sqlc generate

test:   # Run tests for the application
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate_up migrate_down   # Declare specified targets as phony