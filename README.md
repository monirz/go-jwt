# GoJWT - JSON Web Token Authentication in Go

GoJWT is a simple Go (Golang) application that provides authentication using JSON Web Tokens (JWT). It exposes two API endpoints for user login and signup.

## Prerequisites

## Getting Started

### Clone the Repository

```
$ git clone https://github.com/monirz/gojwt.git
$ cd gojwt
```

Create JWT Keys
To generate the JWT keys, you can use the following commands:

# Generate a private key (jwtRS256.key)
ssh-keygen -t rsa -b 4096 -m PEM -f keys/jwtRS256.key

# Don't add a passphrase when prompted

# Extract the public key (jwtRS256.key.pub)
openssl rsa -in keys/jwtRS256.key -pubout -outform PEM -out keys/jwtRS256.key.pub
Make sure to keep these keys secure and do not share them publicly.

Configure Environment Variables
Create a .env file in the project root directory and configure the following environment variables:

```
Copy code
PORT=8090
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=gojwt
DB_PORT=5432
```
Build and Run the Application
Use Docker Compose to build and run the application and PostgreSQL database:


docker-compose up --build
API Endpoints
Login
Endpoint: POST /api/v1/login

## Input (JSON)
json

{
  "email": "user@example.com",
  "password": "your_password"
}
Output (JSON)

## Curl example 

```
 curl -X POST \
  http://localhost:8090/api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "admin@example.com",
    "password": "password"
  }'
```

```
{
  "message": "Login successful",
  "access_token": "your_access_token",
  "refresh_token": "your_refresh_token"
}
```
Signup
 Endpoint: POST /api/v1/signup 

Input (JSON)

```
{
  "email": "newuser@example.com",
  "password": "new_password",
  "username": "newusername"
}
```


Output (JSON)
```
{
  "message": "Signup successful"
}
```


Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.





