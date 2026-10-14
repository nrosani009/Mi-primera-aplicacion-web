package logic

import (
	"errors"
)

type UsuarioInput struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	Telefono  string `json:"telefono"`
	PagoAlDia bool   `json:"pago_al_dia"`
}

func ValidateUsuario(u UsuarioInput) error {
	if u.Nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	if u.Apellido == "" {
		return errors.New("el apellido no puede estar vacío")
	}
	if u.Email == "" {
		return errors.New("el email no puede estar vacío")
	}
	return nil
}
