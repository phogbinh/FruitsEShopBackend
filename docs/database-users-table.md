# Database 'users' Table
## cURL API
### Create an user
#### Description
Create an user to the database table `users`.
#### Response
An JSON object containing the requested user's information fetched from the database.
#### Example
##### Create user `bill`
`curl -X POST -d "{ \"mail\": \"bill@gmail.com\", \"password\": \"1111\", \"userName\": \"bill\", \"nickname\": \"kyo\", \"fname\": \"Phong Binh\", \"lname\": \"Tran\", \"phone\": \"0987654321\", \"location\": \"Taipei, Taiwan\", \"money\": \"1000.00\", \"introduction\": \"Programming geek.\" }" localhost:8080/users`
##### Create user `anna`
`curl -X POST -d "{ \"mail\": \"anna@hotmail.com\", \"password\": \"A@mmX\", \"userName\": \"anna\", \"nickname\": \"annychan\", \"fname\": \"Anna\", \"lname\": \"Carter\", \"phone\": \"0111222333\", \"location\": \"Texas, US\", \"money\": \"1000000.90\", \"introduction\": \"Girl.\" }" localhost:8080/users`
##### Create user `mathew`
`curl -X POST -d "{ \"mail\": \"mathew@yahoo.com\", \"password\": \"MostHandsomePersonInTheWorld\", \"userName\": \"mathew\", \"nickname\": \"mat\", \"fname\": \"Mathew\", \"lname\": \"Brown\", \"phone\": \"0920655185\", \"location\": \"Houston, US\", \"money\": \"888.88\", \"introduction\": \"Bruh.\" }" localhost:8080/users`
##### Create user `john`
`curl -X POST -d "{ \"mail\": \"john@gmail.com\", \"password\": \"JohnnyNeverDies\", \"userName\": \"john\", \"nickname\": \"johnny\", \"fname\": \"John\", \"lname\": \"Butler\", \"phone\": \"0999666999\", \"location\": \"California, US\", \"money\": \"10.01\", \"introduction\": \"Poor.\" }" localhost:8080/users`
##### Create user `duke`
`curl -X POST -d "{ \"mail\": \"duke@hotmail.com\", \"password\": \"Mr.Duke\", \"userName\": \"duke\", \"nickname\": \"duker\", \"fname\": \"Duke\", \"lname\": \"Bennett\", \"phone\": \"0888222555\", \"location\": \"London, UK\", \"money\": \"99999999.99\", \"introduction\": \"I am rich.\" }" localhost:8080/users`
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
    "introduction": "Programming geek."
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Get all users
#### Description
Get all users' information from the database table `users`.
#### Response
An JSON object containing all users' information fetched from the database.
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
        "introduction": "Girl."
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
        "introduction": "Programming geek."
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

### Get an user
#### Description
Get an user from the database table `users`.
#### Response
An JSON object containing the user's information fetched from the database.
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
    "introduction": "Programming geek."
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Update an user password
#### Description
Update an user password in the database table `users`.
#### Response
An JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X PUT -d "{ \"password\": \"666\" }" localhost:8080/users/bill`
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
    "introduction": "Programming geek."
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Delete an user
#### Description
Delete an user from the database table `users`.
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