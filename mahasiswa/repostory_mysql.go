package mahasiswa

import (
	"context"
	"fmt"
	"log"
	"time"

	"crud_mysql/config"
	"crud_mysql/models"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {

	var mahasiswas []models.Mahasiswa

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		fmt.Println(mahasiswa)
		fmt.Println("batas")
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.NIM,
			&mahasiswa.Name,
			&mahasiswa.Semester,
			&createdAt,
			&updatedAt); err != nil {
			fmt.Println("masuk")
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas, nil
}

// GetAll
func DeleteById(id string, ctx context.Context) ([]models.Mahasiswa, error) {

	var mahasiswas []models.Mahasiswa

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id=%v", table, id)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		fmt.Println(mahasiswa)
		fmt.Println("batas")
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.NIM,
			&mahasiswa.Name,
			&mahasiswa.Semester,
			&createdAt,
			&updatedAt); err != nil {
			fmt.Println("masuk")
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas, nil
}
