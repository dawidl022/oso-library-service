package models

type Role struct {
	Name string
}

// type Region struct {
// 	Name string
// }

type User struct {
	Name string
	Role Role
	// Regions []Region
}
