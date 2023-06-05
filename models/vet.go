package models

type Vet struct {
	Id        int    `json:"id" validate:"required" gorm:"primaryKey"`
	Nombre    string `json:"nombre" validate:"required" gorm:"not null"`
	ApellidoP string `json:"apellidoP" validate:"required" gorm:"not null"`
	ApellidoM string `json:"apellidoM" validate:"required" gorm:"not null"`
	Telefono  string `json:"telefono" validate:"required" gorm:"not null"`
	Email     string `json:"email" validate:"required" gorm:"not null"`
	Password  string `json:"password" validate:"required" gorm:"not null"`
}
