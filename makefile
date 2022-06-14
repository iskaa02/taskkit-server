up:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema up

down:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema down

drop:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./db/sql/schema drop
