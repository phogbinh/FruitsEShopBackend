# Database `users` table
## cURL API
### Get all users
#### Description
Get all users' information from the database table `users`.
#### Authorization
N/A.
#### Response
A JSON object containing all users' information fetched from the database.
#### Example
`curl -X GET localhost:8080/users`
#### Expected response format
```json
[
    {
        "mail": "anna@hotmail.com",
        "password": "A@mmX",
        "userName": "anna",
        "nickname": "annychan",
        "fname": "Anna",
        "lname": "Carter",
        "phone": "0111222333",
        "location": "Texas, US",
        "money": "1000000.90",
        "introduction": "Girl.",
        "staffFlag": "1"
    },
    {
        "mail": "bill@gmail.com",
        "password": "1111",
        "userName": "bill",
        "nickname": "kyo",
        "fname": "Phong Binh",
        "lname": "Tran",
        "phone": "0987654321",
        "location": "Taipei, Taiwan",
        "money": "1000.00",
        "introduction": "Programming geek.",
        "staffFlag": "0"
    },
    ...
]
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Get a user by user name
#### Description
Get a user by user name from the database table `users`.
#### Authorization
N/A.
#### Response
A JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X GET localhost:8080/users/bill`
#### Expected response format
```json
{
    "mail": "bill@gmail.com",
    "password": "1111",
    "userName": "bill",
    "nickname": "kyo",
    "fname": "Phong Binh",
    "lname": "Tran",
    "phone": "0987654321",
    "location": "Taipei, Taiwan",
    "money": "1000.00",
    "introduction": "Programming geek.",
    "staffFlag": "0"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Get a user by mail
#### Description
Get a user by mail from the database table `users`.
#### Authorization
N/A.
#### Response
A JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X GET localhost:8080/user?Mail=bill@gmail.com`
#### Expected response format
```json
{
    "mail": "bill@gmail.com",
    "password": "1111",
    "userName": "bill",
    "nickname": "kyo",
    "fname": "Phong Binh",
    "lname": "Tran",
    "phone": "0987654321",
    "location": "Taipei, Taiwan",
    "money": "1000.00",
    "introduction": "Programming geek.",
    "staffFlag": "0"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Delete a user
#### Description
Delete a user from the database table `users`.
#### Authorization
N/A.
#### Response
N/A.
#### Example
`curl -X DELETE localhost:8080/users/bill`
#### Error response format
```json
{
    "error": "error message."
}
```

### Login
#### Description
Authenticate a user login mail and password with that fetched from the database table `users`.
#### Authorization
N/A.
#### Response
A JSON object containing the authentication information, consisting of a JWT token and its expiry date.
#### Example
`curl -X POST -d "{ \"mail\": \"bill@gmail.com\", \"password\": \"1111\" }" localhost:8080/login`
#### Expected response format
```json
{
    "code": "200",
    "expire": "<token expiry date>",
    "token": "<token>"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```
or
```json
{
    "code": "error http status code.",
    "message": "error message."
}
```

### Sign up
#### Description
Create a user to the database table `users`.
#### Authorization
N/A.
#### Response
A JSON object containing the newly created user's information fetched from the database.
#### Example
##### Create user `bill`
`curl -X POST -d "{ \"mail\": \"bill@gmail.com\", \"password\": \"1111\", \"userName\": \"bill\", \"nickname\": \"kyo\", \"fname\": \"Phong Binh\", \"lname\": \"Tran\", \"phone\": \"0987654321\", \"location\": \"Taipei, Taiwan\", \"money\": \"1000.00\", \"introduction\": \"Programming geek.\", \"staffFlag\": \"0\" }" localhost:8080/sign-up`
##### Create user `anna`
`curl -X POST -d "{ \"mail\": \"anna@hotmail.com\", \"password\": \"A@mmX\", \"userName\": \"anna\", \"nickname\": \"annychan\", \"fname\": \"Anna\", \"lname\": \"Carter\", \"phone\": \"0111222333\", \"location\": \"Texas, US\", \"money\": \"1000000.90\", \"introduction\": \"Girl.\", \"staffFlag\": \"1\" }" localhost:8080/sign-up`
##### Create user `mathew`
`curl -X POST -d "{ \"mail\": \"mathew@yahoo.com\", \"password\": \"MostHandsomePersonInTheWorld\", \"userName\": \"mathew\", \"nickname\": \"mat\", \"fname\": \"Mathew\", \"lname\": \"Brown\", \"phone\": \"0920655185\", \"location\": \"Houston, US\", \"money\": \"888.88\", \"introduction\": \"Bruh.\", \"staffFlag\": \"0\" }" localhost:8080/sign-up`
##### Create user `john`
`curl -X POST -d "{ \"mail\": \"john@gmail.com\", \"password\": \"JohnnyNeverDies\", \"userName\": \"john\", \"nickname\": \"johnny\", \"fname\": \"John\", \"lname\": \"Butler\", \"phone\": \"0999666999\", \"location\": \"California, US\", \"money\": \"10.01\", \"introduction\": \"Poor.\", \"staffFlag\": \"0\" }" localhost:8080/sign-up`
##### Create user `duke`
`curl -X POST -d "{ \"mail\": \"duke@hotmail.com\", \"password\": \"Mr.Duke\", \"userName\": \"duke\", \"nickname\": \"duker\", \"fname\": \"Duke\", \"lname\": \"Bennett\", \"phone\": \"0888222555\", \"location\": \"London, UK\", \"money\": \"99999999.99\", \"introduction\": \"I am rich.\", \"staffFlag\": \"0\" }" localhost:8080/sign-up`
#### Expected response format
```json
{
    "mail": "bill@gmail.com",
    "password": "1111",
    "userName": "bill",
    "nickname": "kyo",
    "fname": "Phong Binh",
    "lname": "Tran",
    "phone": "0987654321",
    "location": "Taipei, Taiwan",
    "money": "1000.00",
    "introduction": "Programming geek.",
    "staffFlag": "0"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Update user password
#### Description
Update a user password in the database table `users`.
#### Authorization
Required.
#### Response
A JSON object containing the updated user's information fetched from the database.
#### Example
`curl -X PUT -H "Authorization: Bearer <token>" -d "{ \"password\": \"666\" }" localhost:8080/auth/users/bill`
#### Expected response format
```json
{
    "mail": "bill@gmail.com",
    "password": "666",
    "userName": "bill",
    "nickname": "kyo",
    "fname": "Phong Binh",
    "lname": "Tran",
    "phone": "0987654321",
    "location": "Taipei, Taiwan",
    "money": "1000.00",
    "introduction": "Programming geek.",
    "staffFlag": "0"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Register staff
#### Description
Set a user's staff flag to `true` in the database table `users`.
#### Authorization
Required.
#### Response
A JSON object containing the updated user's information fetched from the database.
#### Example
`curl -X PUT -H "Authorization: Bearer <token>" localhost:8080/auth/users/bill/register-staff`
#### Expected response format
```json
{
    "mail": "bill@gmail.com",
    "password": "666",
    "userName": "bill",
    "nickname": "kyo",
    "fname": "Phong Binh",
    "lname": "Tran",
    "phone": "0987654321",
    "location": "Taipei, Taiwan",
    "money": "1000.00",
    "introduction": "Programming geek.",
    "staffFlag": "1"
}
```
#### Error response format
```json
{
    "error": "error message."
}
```