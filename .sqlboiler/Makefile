include .env

DATABASE_URL=postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

# Run Up Migration
migration-up:
	migrate -path ../db/migration -database "postgresql://root:password@127.0.0.1:5433/gostra_migrate_db?sslmode=disable" -verbose up
	
# Run Down Migration
migration-down:
	migrate -path ../db/migration -database "${DATABASE_URL}" -verbose down


gen-model:
	sqlboiler psql

.PHONY: migration-up migration-down gen-model