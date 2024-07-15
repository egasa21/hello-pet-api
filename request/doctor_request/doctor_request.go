package doctor_request

type DoctorRequest struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}
