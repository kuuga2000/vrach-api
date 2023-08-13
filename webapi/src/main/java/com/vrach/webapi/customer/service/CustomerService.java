package com.vrach.webapi.customer.service;

import com.vrach.webapi.customer.exception.CustomerNotFoundException;
import com.vrach.webapi.customer.model.Customer;
import com.vrach.webapi.customer.repository.CustomerRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class CustomerService implements CustomerInterface<Customer> {

    private final CustomerRepository customerRepository;

    @Override
    public List<Customer> list() {
        return customerRepository.findAll();
    }

    @Override
    public Customer create(Customer customer) {
        return customerRepository.save(customer);
    }

    @Override
    public Optional<Customer> get(long id) {
        return Optional.ofNullable(customerRepository.findById(id)
                .orElseThrow(() -> new CustomerNotFoundException("No customer found with the id : " + id)));
    }

    @Override
    public void update(Customer customer, int id) {

    }

    @Override
    public void delete(int id) {
        Optional<Customer> customer = get(id);
        customer.ifPresent(cst -> customerRepository.deleteById(cst.getId()));
    }
}
