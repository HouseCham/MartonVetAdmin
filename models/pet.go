package models

import "time"

type Pet struct {
	Id              int       `json:"id" gorm:"primary_key"`
	Nombre          string    `json:"nombre" validate:"required" gorm:"not null"`
	FechaNacimiento time.Time `json:"fechaNacimiento" validate:"required" gorm:"not null"`
	Genero          string    `json:"genero" validate:"required" gorm:"not null"`
	ClienteId       int       `json:"clienteId" validate:"required" gorm:"not null"`
	Raza            string    `json:"raza" validate:"required" gorm:"not null"`
	ImgUrl          string    `json:"imgUrl" validate:"required" gorm:"not null"`
}