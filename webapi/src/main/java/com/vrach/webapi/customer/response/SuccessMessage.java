package com.vrach.webapi.customer.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.vrach.webapi.customer.model.Customer;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

@Getter
@Setter
@JsonSerialize
public class SuccessMessage {
        @JsonProperty("success")
        private String Success;
        @JsonProperty("message")
        private String Message;
}
