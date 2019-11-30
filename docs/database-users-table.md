# Database 'users' Table
## cURL API
### Get all users
#### Description
Get all users' information from the database table `users`.
#### Response
An JSON object containing all users' information fetched from the database.
#### Example
`curl -X GET localhost:8080/users`

### Create an user
#### Description
Create an user to the database table `users`.
#### Response
An JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X POST -d "{ \"mail\": \"bill@mail.com\", \"password\": \"1111\", \"userName\": \"bill\", \"nickname\": \"kyo\", \"fname\": \"Phong Binh\", \"lname\": \"Tran\", \"phone\": \"0987654321\", \"location\": \"Taipei, Taiwan\", \"money\": \"1000\", \"introduction\": \"Programming geek.\" }" localhost:8080/users`

### Get an user
#### Description
Get an user from the database table `users`.
#### Response
An JSON object containing the user's information fetched from the database.
#### Example
`curl -X GET localhost:8080/users/bill`

### Update an user password
#### Description
Update an user password in the database table `users`.
#### Response
An JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X PUT -d "{ \"password\": \"666\" }" localhost:8080/users/bill`

### Delete an user
#### Description
Delete an user from the database table `users`.
#### Response
An JSON object containing the requested user's information fetched from the database.
#### Example
`curl -X DELETE localhost:8080/users/bill`