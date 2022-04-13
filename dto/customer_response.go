package dto

type CustomerResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	ZIPCode   string `json:"zipCode"`
	BirthDate string `json:"birthDate"`
	Status    string `json:"status"`
}
