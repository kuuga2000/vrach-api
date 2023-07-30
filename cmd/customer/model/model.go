package model

import (
	"vrachapi/configs"
)

func CreateAccount(customer Customer) (Customer, error) {
	db := configs.GetDB()
	sqlStatement := `INSERT INTO customer (firstname, lastname, date_of_birth, email, gender) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(
		sqlStatement,
		customer.CustomerFirstname,
		customer.CustomerLastname,
		customer.CustomerDateOfBirth,
		customer.CustomerCustomerEmail,
		customer.CustomerGender).Scan(&customer.CustomerId)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
