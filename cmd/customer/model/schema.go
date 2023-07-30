package model

type Customer struct {
	CustomerId            int    `json:"id"`
	CustomerFirstname     string `json:"firstname"`
	CustomerLastname      string `json:"lastname"`
	CustomerDateOfBirth   string `json:"date_of_birth"`
	CustomerCustomerEmail string `json:"email"`
	CustomerGender        int    `json:"gender"`
}

type Customers struct {
	Customers []Customer `json:"customers"`
}
