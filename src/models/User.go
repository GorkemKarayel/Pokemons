package models

type User struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
	Age       int
	Pokemons  []Pokemon
}
