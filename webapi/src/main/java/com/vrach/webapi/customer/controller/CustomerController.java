package com.vrach.webapi.customer.controller;

import com.vrach.webapi.customer.model.Customer;
import com.vrach.webapi.customer.response.SuccessMessage;
import com.vrach.webapi.customer.service.CustomerService;
import lombok.AllArgsConstructor;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static org.springframework.http.HttpStatus.CREATED;

@RequiredArgsConstructor
@RestController
@RequestMapping("/v1")
public class CustomerController {

    private final CustomerService service;

    @GetMapping("/")
    public String hello() {
        return "Hello";
    }
    
    @GetMapping("/account/list")
    public List<Customer> get() {
        return service.list();
    }

    @PostMapping("/register")
    public Object create(@RequestBody Customer customer) {
        service.create(customer);
        Map<String, Object> message = new HashMap<>();
        message.put("success", true);
        message.put("message", "Account has been created successfully");
        return message;
    }

    @DeleteMapping("/delete/{id}")
    public void delete(@PathVariable("id") int id) {
        service.delete(id);
    }
}
