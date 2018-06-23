build:
	go build -o ./bin/server ./cmd/server 

init-db:
	cd ./db && docker build -t users-db .
	docker run -p 5432:5432 -d users-db

run:
	DEBUG=0 PORT=8888 DATABASE_URL=postgres://postgres:@127.0.0.1:5432/users?sslmode=disable \
		./bin/server
