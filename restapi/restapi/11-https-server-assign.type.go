package restapi

type AssignResponseStruct struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Route   string `json:"route"`
}

type AssignBodyStruct struct {
	Pid              string `json:"production_identifier"`
	Project_name     string `json:"project_name"`
	Fleet_name       string `json:"fleet_name"`
	Fleet_id         string `json:"fleet_id"`
	Fleet_assignment bool   `json:"fleet_assignment"`
}

