include .env
start:
	mkdir -p pkg/generate/library
	mkdir -p docs/swagger
	rm -f docs/swagger/*.swagger.json
	protoc --go_out=pkg/generate/library \
 	--go_opt=paths=source_relative \
 	--go-grpc_out=pkg/generate/library \
 	--go-grpc_opt=paths=source_relative \
 	--openapiv2_out=docs/swagger \
 	library.proto
	mkdir -p .data
	docker-compose up -d
	go run cmd/app/main.go
migrate:
	docker exec db sh -c "mysql -uroot -p${MYSQL_PASSWORD} < /migrations/18-07-2023/create_authors_table.sql && mysql -uroot -p${MYSQL_PASSWORD} < /migrations/18-07-2023/create_books_table.sql && mysql -uroot -p${MYSQL_PASSWORD} < /migrations/18-07-2023/create_author_book_table.sql"
db-seed:
	docker exec db sh -c "mysql -uroot -p${MYSQL_PASSWORD} < /migrations/test_data.sql"
run-test:
	cd internals/repository/ && go generate && go test github.com/chernyshevvladislav/test-grpc/internals/service