package com.vrach.webapi.customer.controller;

import com.vrach.webapi.customer.model.Customer;
import com.vrach.webapi.customer.service.CustomerService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

import static org.springframework.http.HttpStatus.CREATED;

@RequiredArgsConstructor
@RestController
@RequestMapping("/api/v1/account")
public class CustomerController {

    private final CustomerService service;

    @GetMapping("/")
    public String hello() {
        return "Hello";
    }
    
    @GetMapping("/list")
    public List<Customer> get() {
        return service.list();
    }

    @PostMapping("/create")
    public ResponseEntity<Customer> create(@RequestBody Customer customer) {
        return new ResponseEntity<>(service.create(customer), CREATED);
    }

    @DeleteMapping("/delete/{id}")
    public void delete(@PathVariable("id") int id) {
        service.delete(id);
    }
}
