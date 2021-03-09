package repository

import (
	"fibertest/models"
)

func GetCats() []models.Cat {

	config := Config{Host: "localhost", Port: "5432", Username: "postgres", Password: "123456", DBname: "catdb", SSLMode: "disable"}
	db, err := NewPostgresDB(config)

	if err != nil {
		println("VSE PROPALO: " + err.Error())
	}

	if db == nil {
		println("DB IS NIL")
	}

	rows, err := db.Queryx("SELECT * FROM cats")
	var cats []models.Cat
	for rows.Next() {
		var cat models.Cat
		err = rows.StructScan(&cat)
		cats = append(cats, cat)
	}

	return cats
}
