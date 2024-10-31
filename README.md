# My-Flat Oauth2 Server

This project will manage access tokens (generation and validation) for the My-Flat project. Those tokens will grant access to the the services a user is able to use.

## Data Model

The following is the data model used for the server, which can be seen at the models package.
![Data Model Image Representation](server_data_mdel.png)


## Cloud Firestore

This project uses Google Cloud Firestore as the database, using a NoSQL database with a cloud native approach.

## Installation

This project is built using Go. To install dependencies, run the following command:

```bash
go mod tidy
```

## Things to do

- Use TTL policies for the tokens, and remove revoke field from the token.
- Blacklist possible cliend ids or IPs which may be intruders.
- Use CI/CD tools to automate the deployment of the server.
- Use a load balancer to distribute the load of the server.