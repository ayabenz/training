package model

type User struct {
	ID int `json:"ID"`
	Name string  `json:"Name"`
	Address string  `json:"Address"`
	Dept Department  `json:"Dept"`
}

type Department struct {
	ID int `json:"ID"`
	Name string  `json:"Name"`
}
