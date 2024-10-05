# go-ranking
Ranking system API written in Go

# Setting Up
Pre-requisites & Installations

```
1) Ensure Go is installed 
https://go.dev/doc/install
```
```
2) Ensure Docker is installed 
https://docs.docker.com/get-started/get-docker/
```
```
3) Ensure Goose is installed (used for data migration)
`go install github.com/pressly/goose/v3/cmd/goose@latest`
```
```
4) Ensure you have Make and can run Makefile
```

# Running the Application

<h3> 1) First, make sure the database is running: </h3>

```
make set-db
```
This will create a postgreSQL DB using Docker. The data is migrated using Goose.

To ensure data persists within the container, run the following for graceful shutdown:
```
make stop-db
```
To start the DB again, run :

```
make start-db
```
**Note: If you wish to "restart" the database with default data, you can still use `make set-db`**

<h3> 2. Running the Server on port 3000:</h3>
Run the following command. 

```
make run
```


# API endpoints

**1. Create New User**
- Endpoint: POST /users
- Request Body:
```
{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "score": 150
}
```
- Response:
  - Status Code 201 Created
  ```
  {
      "status_code": 201,
      "data": "created"
  }
  ```
  - Status Code 409 Conflict (email already in use)
  ```
  {
      "status_code": 409,
      "data": "Email is already in use"
  }
  ```

**2. Get All Users**
- Endpoint: GET /users
- Response:
  - Status Code 200 OK 
  - Response Body:
  ```
  {
      "statusCode": 200,
      "data": [
          {
              "id": 1,
              "name": "John Doe",
              "email": "johndoe@example.com",
              "score": 150
          },
          {
              "id": 2,
              "name": "Jane Smith",
              "email": "janesmith@example.com",
              "score": 200
          }
      ]
  }
  ```
  
**3. Get User by ID**
- Endpoint: GET /users/{id}
- Response:
  - Status Code 200 OK
  - Response Body:
  ```
  {
      "statusCode": 200,
      "data": {
          "id": 1,
          "name": "John Doe",
          "email": "johndoe@example.com",
          "score": 150
      }
  }
  ```
- Error Response (if user not found):
  - Status Code: 404 Not Found 
  ```
  {
    "status_code": 404,
    "error": "resource not found"
  }
  ```


**4. Update User by ID**
- Endpoint: POST /users/{id}
- Request Body:
  ```
  {
    "name": "John Doe Updated",
    "email": "johndoe.updated@example.com",
    "score": 200
  }
  ```
- Response:
  - Status Code 204 No Content (successful update)
  

- Error Response (if user not found):
  - Status Code: 404 Not Found
  ```
  {
    "status_code": 404,
    "error": "resource not found"
  }
  ```

**5. Delete User by ID**
- Endpoint: DELETE /users/{id}
- Response:
  - Status Code 204 No Content (successful update)


- Error Response (if user not found):
  - Status Code: 404 Not Found
  ```
  {
    "status_code": 404,
    "error": "resource not found"
  }
  ```
**6. Get All User Ranks**
- Endpoint: GET /users/rank
- Response:
  - Status Code 200 OK 
  - Response Body:
  ```
  {
    "status_code": 200,
    "data": [
        {
            "ID": 3,
            "Name": "Poh",
            "Email": "poh@gmail.com",
            "Score": 70,
            "Rank": 1
        },
        {
            "ID": 2,
            "Name": "Daniel",
            "Email": "daniel@gmail.com",
            "Score": 60,
            "Rank": 2
        },
        {
            "ID": 1,
            "Name": "John",
            "Email": "john@gmail.com",
            "Score": 50,
            "Rank": 3
        }
    ]
  }

**6. Get User Rank by ID**
- Endpoint: GET /users/rank/{id}
- Response:
  - Status Code 200 OK
  - Response Body:
  ```
  {
    "status_code": 200,
    "data": {
        "ID": 1,
        "Name": "John",
        "Email": "john@gmail.com",
        "Score": 50,
        "Rank": 3
    }
  }
  ```
