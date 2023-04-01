package models

import (
	"time"
)

type Mahasiswa struct {
	ID        int       `json:"id"`
	NIM       int       `json:"nim"`
	Name      string    `name:"fullname"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
