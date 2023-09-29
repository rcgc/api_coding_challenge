import http from 'k6/http';
import { check } from 'k6';

const url = 'http://localhost:8080/cars/'

function testGetCars_WhenStatusOk_Response200(){
  const res = http.get(url);
  check(res, { 'Test GetCars when status OK response 200': (r) => r.status == 200 });
}

function testGetCarById_WhenStatusOk_Response200(){  
  const id = 'JHk290Xj'

  const res = http.get(url+id);
  check(res, { 'Test GetCarById when status OK response 200': (r) => r.status == 200 });
}

function testGetCarById_WhenStatusNotFound_Response404(){
  const id = 'zzzzzzzzz0'

  const res = http.get(url+id)
  check(res, {'Test GetCarById when status Not Found response status 404': (r) => r.status == 404})
}

function testCreateCar_WhenStatusCreated_Response201(){
  const body = {
    Id: "abcdefghij",
    Make: "Nissan",
    Model: "March",
    Package: "XX",
    Color: "Gray",
    Year: 2013,
    Category: "SUV",
    Mileage: 799,
    Price: 2499000
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, JSON.stringify(body), params);
  check(res, {'Test CreateCar when status OK response 201': (r) => r.status == 201});
}

function testCreateCar_WhenStatusBadRequest_Response400(){
  const body = {
    Id: "",
    Make: "",
    Model: "",
    Package: "",
    Color: "",
    Year: 0,
    Category: "",
    Mileage: 0,
    Price: 0
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, JSON.stringify(body), params);
  check(res, {'Test CreateCar when status Bad Request response 400': (r) => r.status == 400});
}

function testUpdateCar_WhenStatusOk_Response200(){
  const body = {
    Id: "abcdefghij",
    Make: "Nissan",
    Model: "March",
    Package: "XX",
    Color: "Gray",
    Year: 2014,
    Category: "SUV",
    Mileage: 799,
    Price: 2499000
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.put(url, JSON.stringify(body), params)
  check(res, {'Test UpdateCar when status OK response 200': (r) => r.status == 200});
}

function testUpdateCar_WhenStatusBadRequest_Response400(){
  const body = {
    Id: "",
    Make: "",
    Model: "",
    Package: "",
    Color: "",
    Year: 0,
    Category: "",
    Mileage: 0,
    Price: 0
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.put(url, JSON.stringify(body), params)
  check(res, {'Test UpdateCar when status Bad Request response 400': (r) => r.status == 400});
}

function testUpdateCar_WhenStatusNotFound_Response404(){
  const body = {
    Id: "stuvwxyz0",
    Make: "Nissan",
    Model: "March",
    Package: "XX",
    Color: "Gray",
    Year: 2014,
    Category: "SUV",
    Mileage: 799,
    Price: 2499000
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.put(url, JSON.stringify(body), params)
  check(res, {'Test UpdateCar when status Not Found response 404': (r) => r.status == 404});
}

function testDeleteCar_WhenStatusNoContent_Response204(){
  const id = 'abcdefghij'

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.del(url+id, params)
  check(res, {'Test DeleteCar when status No Content response 204': (r) => r.status == 204});
}

function testDeleteCar_WhenStatusNotFound_Response404(){
  const id = '11zxyabcdfg'

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.del(url+id, params)
  check(res, {'Test DeleteCar when status Not Found response 404': (r) => r.status == 404});
}

export default function(){
  testGetCars_WhenStatusOk_Response200()

  testGetCarById_WhenStatusOk_Response200()
  testGetCarById_WhenStatusNotFound_Response404()

  testCreateCar_WhenStatusCreated_Response201()
  testCreateCar_WhenStatusBadRequest_Response400()

  testUpdateCar_WhenStatusOk_Response200()
  testUpdateCar_WhenStatusBadRequest_Response400()
  testUpdateCar_WhenStatusNotFound_Response404()

  testDeleteCar_WhenStatusNoContent_Response204()
  testDeleteCar_WhenStatusNotFound_Response404()
}