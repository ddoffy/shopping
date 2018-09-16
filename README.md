
# Shopping API using Gin, Golang, JWT, Dependency Injection, Unit testing

## Description
This is an shopping cart service implementation in Go (Golang)+ Gin projects.

This project has  4 Domain layer :
 * Models Layer 
 * Repository Layer
 * Business logic/Usecase Layer  
 * API/Delivery Layer

### How To Run This Project

```bash
#move to directory
cd $GOPATH/src/github.com/karuppaiah

# Clone into YOUR $GOPATH/src
git clone https://github.com/karuppaiah/shopping.git

#move to project
cd shopping

# Run the script
sh execute.sh

# Data populate and setup DB
sh createData.sh

At this time, you will have a new data.db created in root directory. Change the DB if needed.

Site runs at http://127.0.0.1:8080/ping

Postman V2 request: https://github.com/karuppaiah/shopping/blob/master/golang%20shopping.postman_collection

Always get the JWT token and use them in authorization header for response.

```


# TODO :
- [ ]  User management
- [ ] Docker image and publish in hub.docker.com - In progress
- [ ] GOlang IOC
- [ ] redis storage for JWT token
- [ ] Move to Mongo DB
- [ ] Write more unit testing for more coverage
- [ ] React frontend
- [ ] Utilites for message Q like kafka(producer and consumer)
- [ ] Tree structure for promotion to enable double source product rule
- [ ] Stress testing scripts





