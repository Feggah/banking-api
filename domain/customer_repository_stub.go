package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (r CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return r.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:        "1001",
			Name:      "Ashish",
			City:      "New Delhi",
			ZIPCode:   "110011",
			BirthDate: "2000-01-01",
			Status:    "1",
		},
		{
			ID:        "1002",
			Name:      "Rob",
			City:      "New Delhi",
			ZIPCode:   "110011",
			BirthDate: "2000-01-01",
			Status:    "1",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
