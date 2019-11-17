# Database 'users' Table
## cURL API
### Get all users
#### Description
Get all users' names and passwords from the database table `users`.
#### Response
An JSON object containing all users' names and passwords fetched from the database.
#### Example
`curl -X GET localhost:8080/users`

### Create an user
#### Description
Create an user to the database table `users`.
#### Response
An JSON object containing the requested user's name and password.
#### Example
`curl -X POST -d "{ \"username\": \"bill\", \"password\": \"1\" }" localhost:8080/users`

### Get an user
#### Description
Get an user from the database table `users`.
#### Response
An JSON object containing the user's name and password fetched from the database.
#### Example
`curl -X GET localhost:8080/users/bill`

### Update an user password
#### Description
Update an user password in the database table `users`.
#### Response
An JSON object containing the requested user's name and password.
#### Example
`curl -X PUT -d "{ \"username\": \"bill\", \"password\": \"666\" }" localhost:8080/users/bill`

### Delete an user
#### Description
Delete an user from the database table `users`.
#### Response
An JSON object containing the user's name given in the requested URL.
#### Example
`curl -X DELETE localhost:8080/users/bill`