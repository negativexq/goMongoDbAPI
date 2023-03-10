# CRUD API

This is a simple API for performing CRUD operations on components stored in a MongoDB database, built with the Golang Fiber v2 framework.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Golang](https://golang.org/)
- [Fiber v2](https://gofiber.io/)
- [MongoDB](https://www.mongodb.com/)

### Installing

1.Clone the repository to your local machine:

git clone https://github.com/negativexq/goMongoDbAPI.git

2.Install the dependencies:

go get -d

3.Start the MongoDB server.

4.Update the .env file with your MongoDB URI and database name.

5.Run the API:

go run main.go


## API Endpoints

The following endpoints are available:

- `GET /components`: Retrieve a list of all components
- `GET /components/:id`: Retrieve a specific component by ID
- `POST /components`: Create a new component
- `PUT /components/:id`: Update a specific component by ID
- `DELETE /components/:id`: Delete a specific component by ID

## Built With

- [Golang](https://golang.org/)
- [Fiber v2](https://gofiber.io/)
- [MongoDB](https://www.mongodb.com/)
