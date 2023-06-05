package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primary_key"`
	Nombre    string `json:"nombre" validate:"required" gorm:"not null"`
	ApellidoP string `json:"apellidoP" validate:"required" gorm:"not null"`
	ApellidoM string `json:"apellidoM" validate:"required" gorm:"not null"`
	Telefono  string `json:"telefono" validate:"required" gorm:"not null"`
	Email     string `json:"email" validate:"required" gorm:"not null"`

	Calle        string `json:"calle" validate:"required" gorm:"not null"`
	NumeroExt    string `json:"numero" validate:"required" gorm:"not null"`
	NumeroInt    string `json:"numeroInt"`
	Colonia      string `json:"colonia" validate:"required" gorm:"not null"`
	CodigoPostal string `json:"codigoPostal" validate:"required" gorm:"not null"`
	Ciudad       string `json:"ciudad" validate:"required" gorm:"not null"`
	Estado       string `json:"estado" validate:"required" gorm:"not null"`
}

func (c *Client) ParseClientParams(client Client) {
	c.Nombre = client.Nombre
	c.ApellidoP = client.ApellidoP
	c.ApellidoM = client.ApellidoM
	c.Telefono = client.Telefono
	c.Email = client.Email
	c.Calle = client.Calle
	c.NumeroExt = client.NumeroExt
	c.NumeroInt = client.NumeroInt
	c.Colonia = client.Colonia
	c.CodigoPostal = client.CodigoPostal
	c.Ciudad = client.Ciudad
	c.Estado = client.Estado
}