package models

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"movielist-app/pkg/configuration"


)

func connectionString() string {
	cfg := configuration.GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", 
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Address,
		cfg.Database.Port,
		cfg.Database.Name)
}

func ListMovies() ([]Movie, error) {
	db, err := sql.Open("mysql", connectionString())

	if err != nil {
	 return nil, err
	}
   
	defer db.Close()
   
	results, err := db.Query("SELECT * FROM movies")
   
	if err != nil {
	 return nil, err
   
	}
	movies1 := []Movie{}
	for results.Next() {
	 var mov Movie
	 err = results.Scan(&mov.ID, &mov.Title, &mov.Director)
   
	 if err != nil {
	  return nil, err
	 }
	 movies1 = append(movies1, mov)
	}
	return movies1, nil
}


func GetMovie(id string) (*Movie, error) {
	db, err := sql.Open("mysql", connectionString())

	if err != nil {
	 return nil, err
	}
   
	defer db.Close()
   
	results, err := db.Query("SELECT * FROM movies where id=?", id)
   
	if err != nil {
	 return nil, err
	}

	mov := &Movie{}

	if results.Next() {
	 err = results.Scan(&mov.ID, &mov.Title, &mov.Director)
   
	 if err != nil {
	 	return nil, err
	 }
	} else {
		return nil, nil
	}

	return mov, nil
}

func CreateMovie(movie Movie) error {
	db, err := sql.Open("mysql", connectionString())
	if err != nil {
		return err
	}
   
	defer db.Close()
	insert, err := db.Query(
	 "INSERT INTO movies (title,director) VALUES (?,?)",
	 movie.Title, movie.Director)
 	defer insert.Close()
  
	if err != nil {
		return err
	}

	return nil
}
