up:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./schema up

down:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./schema down

drop:
	migrate -database "postgres://postgres:password@localhost:5432?sslmode=disable" -path ./schema drop
ent:
	go generate ./ent
