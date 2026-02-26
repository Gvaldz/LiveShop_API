package entities

type User struct {
	IdUser   int32  `json:"id"`        
    Name     string `json:"name"`      
    Number    string `json:"numebr"`     
    Password string `json:"password"`
}