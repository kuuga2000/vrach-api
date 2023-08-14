package com.vrach.webapi.customer.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.*;
import lombok.*;

import java.sql.Time;
import java.sql.Timestamp;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "customer")
public class Customer {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String firstname;

    private String lastname;

    private String nickname;

    @JsonProperty("date_of_birth")
    private String dateOfBirth;

    private String email;

    @JsonProperty("contact_number")
    private String contactNumber;

    private int gender;

    private String username;

    private String password;

}
