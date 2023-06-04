# ALBUMS API

- A basic crud API for music albums
- To Run the project you have to install Docker in your machine and then
    run the command 
	```
	docker-compose up -d
	``` 
	inside the project folder

The API will run in the port 3000 with some default data

- In order to connect to the database you need to create a local.env file with:
 > DB_HOST="DB_HOST" \
 DB_USER="DB_USER" \
 DB_PASSWORD="DB_PASSWORD" \
 DB_PORT="DB_PORT" \
 DB_NAME="DB_NAME"