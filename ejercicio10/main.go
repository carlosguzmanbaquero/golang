package main

import (
	"fmt"
	"time"

	us "./user"
)

type persona struct {
	us.Usuario
}

func main() {
	user := new(persona)
	user.AltaUsuario(1, "Carlos Guzman", time.Now(), true, "Admin")
	fmt.Printf("nombre: %s id: %d\n", user.Usuario.Nombre, user.Usuario.Id)
}
