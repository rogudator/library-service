db:
	docker run -d --rm --name=library-service-db -e MYSQL_ROOT_PASSWORD='library-service' -e MYSQL_USER='library-service' -e MYSQL_PASSWORD='library-service' -e MYSQL_DATABASE='library-service' -v ${CURDIR}/init:/docker-entrypoint-initdb.d -p 3306:3306 -p 33060:33060 mysql:8.0
protobuf:
	protoc -I ./proto/ --go_out ./ --go-grpc_out ./ ./proto/libraryServicePb/libraryService.proto