package model

type Customer struct {
	CustomerId            int    `json:"id"`
	CustomerFirstname     string `json:"firstname"`
	CustomerLastname      string `json:"lastname"`
	CustomerNickname      string `json:"nickname"`
	CustomerDateOfBirth   string `json:"date_of_birth"`
	CustomerCustomerEmail string `json:"email"`
	CustomerGender        int    `json:"gender"`
	CustomerUsername      string `json:"username"`
	CustomerPassword      string `json:"password"`
}

type Customers struct {
	Customers []Customer `json:"customers"`
}
