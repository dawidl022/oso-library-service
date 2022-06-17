package models

type Book struct {
	Title             string
	GloballyAvailable bool
	Regions           []string
}
