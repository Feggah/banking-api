db:
	docker run -v dbdata:/var/lib/mysql --network host -e MYSQL_ROOT_PASSWORD=admin -d mysql:8

local:
	SERVER_ADDRESS=localhost SERVER_PORT=5000 DB_USER=root DB_PASSWORD=admin DB_ADDRESS=localhost DB_PORT=3306 DB_NAME=banking go run main.go
