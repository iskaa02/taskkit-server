migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema down

migrate-drop:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema drop
