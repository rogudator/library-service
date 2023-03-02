include db.env

build-db:
	docker build -t library-service-mysql .

db:
	docker run \
	-d --rm \
	--name=library-service-db \
	-e MYSQL_ROOT_PASSWORD='library-service' \
	-e MYSQL_USER='library-service' \
	-e MYSQL_PASSWORD='${MYSQL_PASSWORD}' \
	-e MYSQL_DATABASE='library-service' \
	-p 3306:3306 \
	library-service-mysql

local:
	go run cmd/main.go
protobuf:
	protoc -I ./proto/ --go_out ./ --go-grpc_out ./ ./proto/libraryServicePb/libraryService.proto