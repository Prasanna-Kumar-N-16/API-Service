# API-Service
API Service
API Service is an HTTP Server which does all the CRUD Operations with respect to mongoDB.

**LOGIN**



API URL - http://DOMAIN:PORT/login

METHOD:POST

API Server will have a login api which generates an JWT-Token as a cookie to the client service.

//TODO ENHANCEMENT--Can take login credentials of user and validate the credentials

//ONCE Success then generate a JWT Token

Using Cookie authenticate the API Service of other end points with POST,PUT,GET and DELETE Methods.

PORT is configured in .json file

**POST** 

- Posts the company name and employee Name to the DataBase via Server


API URL - http://DOMAIN:PORT/post

METHOD:POST

REQ BODY - RAW JSON

EX:

{
"_id":"123",
"companyName":"ABCD",
"employeeName":"XYZ"
}

**PUT** 

- Creates a record in DB if ID doesn't exist else updates the record with the data.

METHOD:PUT

API URL - http://DOMAIN:PORT/put

REQ BODY - RAW JSON

EX:

{
"_id":"123",
"companyName":"ABCD",
"employeeName":"XYZ"
}

**GET** 

- Creates a record in DB if ID doesn't exist else updates the record with the data.

API URL - http://DOMAIN:PORT/get

METHOD:GET

Query Params - "id":"123"

RESPOSE Body - 

{
"_id":"123",
"companyName":"ABCD",
"employeeName":"XYZ"
}

**DELETE**

- Creates a record in DB if ID doesn't exist else updates the record with the data.

API URL - http://DOMAIN:PORT/delete

METHOD:DELETE

Query Params - "id":"123"

RESPOSE Body - 

"Successfully deleted the Record"

//ENHANCEMENTS TODO

//ONLY Happiy flow is handled show test all the test cases while handling with API.
