package main

import (
	"database/sql"
	"strconv"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Conn() {
	connStr := "postgresql://" + conf.DataBaseConfig.User + ":" + conf.DataBaseConfig.Password + 
        "@" + conf.DataBaseConfig.Host + ":" + conf.DataBaseConfig.Port + "/" + 
        conf.DataBaseConfig.Name + "?sslmode=disable"
	// Get a database handle.
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}


func GetAll() ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album")
    if err != nil {
        return nil, fmt.Errorf("error fetching albums")
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("error assigning album to the slice")
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("weird error")
    }
    return albums, nil
}

func GetById(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("SELECT * FROM album WHERE id = $1;", id)
    if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
        if err == sql.ErrNoRows {
            return album, fmt.Errorf("albumsById %d: no such album", id)
        }
        return album, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return album, nil
}

func Save(newAlbum AlbumDTO) (Album, error) {
	var alb Album = Album{}
	var lastInsertId int64
	err := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id;", newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&lastInsertId)
    if err != nil {
        return alb, fmt.Errorf("addAlbum: %v", err)
    }
	var savedAlbum = Album{strconv.FormatInt(lastInsertId, 10), newAlbum.Title, newAlbum.Artist, newAlbum.Price}
	return savedAlbum, nil
}

func Update(id int64, newAlbum AlbumDTO) (Album, error) {
	var alb Album = Album{}

	_, err := db.Exec("UPDATE album SET title = $1, artist = $2, price = $3 WHERE id = $4;", newAlbum.Title, newAlbum.Artist, newAlbum.Price, id)
    if err != nil {
        fmt.Println(err)
        return alb, fmt.Errorf("updateAlbum: %v", err)
    }
    
	var updatedAlbum = Album{strconv.FormatInt(id, 10), newAlbum.Title, newAlbum.Artist, newAlbum.Price}
	return updatedAlbum, nil
}

func Delete(id int64) (int64, error) {
    var lastAffectedId int64
	err := db.QueryRow("DELETE FROM album WHERE id = $1 RETURNING id;", id).Scan(&lastAffectedId)
    if err != nil {
        return 0,fmt.Errorf("deleteAlbum: %v", err)
    }
	return lastAffectedId, nil
}