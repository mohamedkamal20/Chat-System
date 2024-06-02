# Chat System

### Overview
This is a creation of messages and authentication endpoints in goLang.

### Goals
* Creating a microservice that simulates a simplified chat platform. 
* This platform will handle user messages and Store them in a distributed database.
* Ensure efficient retrieval and caching mechanisms using redis.
* Use nginx as an entry point for the API request.
* Use Docker to containerize the application and the database.

### Requirements
In order to run the application please follow the steps:
- Ubuntu running OS.
- Docker installed.
- Make sure these ports are not used by another process:
    * port `8080` for App application.
    * port `8085` for Nginx.
    * port `9042` for default Cassandra port.
    * port `6379` for default redis port.
    * port `9090` for default Prometheus port.
    * port `3000` for default grafana port.

### Quick start
* Clone the project .
* Run `sudo docker compose up --build -d`.

### Test services
* ##### Postman or curl command line
  *`{localhost:8085}/api/v1/login/ [POST] {"email": "example@gmail.com", "password":"password"}`\
  *`{localhost:8085}/api/v1/register/ [POST] {"email": "example@gmail.com", "password":"password"}`
  
    ##### Hint : Use Auth Token from Login Endpoint for those APIs 
  *`{localhost:8085}/api/v1/send/ [POST] {"sender": "example@gmail.com", "recipient":"example2@gmail.com", "content":"Hello"}`\
  *`{localhost:8085}/api/v1/messages/{email} [GET] {"email": "example@gmail.com", "password":"password"}`