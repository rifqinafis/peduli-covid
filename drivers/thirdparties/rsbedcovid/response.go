package rsbedcovid

type ProvinceRespon struct {
	Province []Province `json:"provinces"`
}

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityRespon struct {
	City []City `json:"cities"`
}

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HospitalRespon struct {
	Status   int        `json:"status"`
	Hospital []Hospital `json:"hospitals"`
}

type Hospital struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Queue           int    `json:"queue"`
	BedAvailability int    `json:"bed_availability"`
	Info            string `json:"info"`
}

type BedDetailRespon struct {
	Status int       `json:"status"`
	Data   BedDetail `json:"data"`
}

type BedDetail struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	Phone     string       `json:"phone"`
	BedDetail []BedDetails `json:"bedDetail"`
}

type BedDetails struct {
	Time  string         `json:"time"`
	Stats BedDescription `json:"stats"`
}

type BedDescription struct {
	Title        string `json:"title"`
	BedAvailable int    `json:"bed_available"`
	BedEmpty     int    `json:"bed_empty"`
	Queue        int    `json:"queue"`
}

type HospitalLocationRespon struct {
	Status int              `json:"status"`
	Data   HospitalLocation `json:"data"`
}

type HospitalLocation struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Long    string `json:"long"`
	Gmaps   string `json:"gmaps"`
}
