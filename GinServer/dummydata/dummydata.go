package dummydata

type Doctor struct {
	Id   int    `json:"doctor_id"`
	Name string `json:"name"`
}

var Doctors = []Doctor{
	{1, "Uzumaki Naruto"},
	{2, "Uchiha Sasuke"},
	{3, "Uchiha Itachi"},
	{4, "Haruno Sakura"},
}
