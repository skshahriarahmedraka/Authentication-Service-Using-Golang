**Title:**  an authentication service using Golang and the Gin web framework.

**User Registration:**

- Implemention user registration, allowing users to sign up with their email address and password.
  
  ***solution :***
  
  POST  `/register`
  - request 
  
  - ```
    {
        "firstname": "sk",
        "lastname": "ssar",
        "email": "ssar@gmail.com",
        "password": "111187" 
    }
    ```
  
  - response 
  
  - ```
    {
        "message": "successfully signed up"
    }
    ```
    
    

**User Authentication:**

- Develop user login functionality using email and password
  
  ***solution :***
  
  POST `/login`
  - request 
  
  - ```
    {
        "email": "skshahra@gmail.com",
        "password": "111187" 
    }
    ```
  
  - response 
  
  - ```
    {
        "id": "654d046d5b494b47d7705481",
        "message": "Login Successfull"
    }
    ```
- Implement token-based authentication using JWT for secure user authentication. Provide a unique token upon successful authentication.
- Ensure handling of JWT refresh tokens.
  
  ***Solution :*** 
  - **Auth** and **Refresh**  ,  JWT  token is used for  authentication and authorization
  - **Auth** token used for  user authentication and **Refresh** token is used for keep the cookie up to date  



**User Roles and Permissions:**

- Implement role-based access control (RBAC) to assign predefined roles (e.g., user, admin) to users. Note that CRUD operations are not required for this task.
  
  ***solution:***
  
  GET `/:id`
  
  user can access their profile info using their **ID** , 
  
  **ONLY ADMIN**  can retrive any users data using **ID**
  - request 
  
  - ```
    GET localhost:8080/654d046d5b494b47d7705481
    ```
  
  - response 
  
  - ```
    {
        "_id": "654d046d5b494b47d7705481",
        "firstname": "sk",
        "lastname": "raka",
        "email": "skshahra@gmail.com",
        "password": "",
        "telephone": "",
        "address": "",
        "accounttype": "admin"
    }
    ```
  
  - 

**Another** **Demo Service:**

- Create a demo service for authenticated users that allows them to retrieve lists of data arrays.
  
  ***solution :***
  
  GET `/alluser` 
  
  ONLY ADMIN can access this route 
  - request 
  
  - ```
    GET  localhost:8080/alluser
    ```
  
  - Response 
  
  - ```
    [
        {
            "_id": "654d929b76141a0ab1c236e1",
            "firstname": "sk",
            "lastname": "raka",
            "email": "skshahra@gmail.com",
            "password": "$2a$10$uGTE6IRREQHT2frMeJEb8eTF9JsdZMMysnzRotlIcZcbHU4dGfYrW",
            "telephone": "",
            "address": "",
            "accounttype": "admin"
        },
        {
            "_id": "654db17d46147f79d73f079b",
            "firstname": "sk",
            "lastname": "ssar",
            "email": "ssar@gmail.com",
            "password": "$2a$10$WZ2r91gqxEdfGL6A7cfq6.E34tBp129qOqFtetcW1w5HXbl1rtQjm",
            "telephone": "",
            "address": "",
            "accounttype": "normal"
        }
    ]
    ```
  
  - 

**Instructions:**

- For Loggin   `github.com/rs/zerolog`   is used 
- testing code for all route  `handler/handler_test.go`


