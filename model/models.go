package model

type Response struct {
	Persons []Person `json:"persons"`
}

type Organization struct {

}

type Season struct {

}

type Team struct {

}

type Meet struct {

}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Swimmer struct {

}

type Coach struct {

}

type Event struct {

}

type Entry struct {

}

type Pool struct {
	
}
