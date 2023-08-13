package com.vrach.webapi.customer.controller;

import com.vrach.webapi.customer.model.Customer;
import com.vrach.webapi.customer.repository.CustomerRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RequiredArgsConstructor
@RestController
@RequestMapping("/api/v1")
public class CustomerController {

    private final CustomerRepository customerRepository;

    @GetMapping("/")
    public String hello() {
        return "Hello";
    }

    @GetMapping("/account/list")
    public List<Customer> get() {
        return customerRepository.findAll();
    }
}
