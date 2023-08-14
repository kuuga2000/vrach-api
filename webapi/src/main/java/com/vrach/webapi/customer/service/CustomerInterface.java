package com.vrach.webapi.customer.service;

import java.util.List;
import java.util.Optional;

public interface CustomerInterface<T> {

    List<T> list();

    T create(T t);

    Optional<T> get(long id);

    void update(T t, int id);

    void delete(int id);
}
