# ALBUMS API

- A basic crud API for music albums
- To Run the project you have to install Go in your machine and then
    run the command 
	```
	go get .
	go run .
	``` 
	inside the project folder

RUN with postgreSQL database


CREATE TABLE public.album (
	id serial4 NOT NULL,
	title varchar(128) NOT NULL,
	artist varchar(255) NOT NULL,
	price numeric(5, 2) NOT NULL,
	CONSTRAINT album_pkey PRIMARY KEY (id)
);