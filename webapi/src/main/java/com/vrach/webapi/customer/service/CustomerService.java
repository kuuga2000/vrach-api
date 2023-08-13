package com.vrach.webapi.customer.service;

import com.vrach.webapi.customer.model.Customer;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class CustomerService implements CrudService<Customer>{
    @Override
    public List<Customer> list() {
        return null;
    }

    @Override
    public Customer create(Customer customer) {
        return null;
    }

    @Override
    public Optional<Customer> get(int id) {
        return Optional.empty();
    }

    @Override
    public void update(Customer customer, int id) {

    }

    @Override
    public void delete(int id) {

    }
}
