package models

//Smartphone model struct
type Smartphone struct {
	ID            int64 //similar to a long in Java - autoincremental ID in SQL
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

//CreateSmartphoneCMD command to create a new smartphone
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	OS            string `json:"os"`
}
