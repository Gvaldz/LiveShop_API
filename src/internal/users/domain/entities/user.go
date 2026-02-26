package entities

type User struct {
	IdUser   int32  `json:"id"`        // Se mapea a "id" en la respuesta
    Name     string `json:"name"`      // Se mapea a "name" en el JSON entrante
    Lastname string `json:"lastname"`  // Se mapea a "lastname"
    Email    string `json:"email"`     // Se mapea a "email"
    Password string `json:"password"`
}