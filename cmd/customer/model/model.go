package model

import (
	"vrachapi/configs"
)

func CreateAccount(customer Customer) (Customer, error) {
	db := configs.GetDB()
	sqlStatement := `INSERT INTO customer (firstname, lastname, nickname, date_of_birth, email, gender, username, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.QueryRow(
		sqlStatement,
		customer.CustomerFirstname,
		customer.CustomerLastname,
		customer.CustomerNickname,
		customer.CustomerDateOfBirth,
		customer.CustomerCustomerEmail,
		customer.CustomerGender,
		customer.CustomerUsername,
		customer.CustomerPassword).Scan(&customer.CustomerId)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
