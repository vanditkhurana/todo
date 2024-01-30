# Todos API

Golang API to create, fetch, update and delete todos from ScyllaDb

## Installation

Clone the repo and install the dependencies.

```bash
go get -u github.com/vanditkhurana/todo_api
```

## Start
To start the server, run the following

```bash
go run main.go
```

Open [http://localhost:8080](http://localhost:8080) and take a look around.

## Establish container of scylla-local using docker 
Considering you have docker installed and the engine is started. run below commands

1. Pull the ScyllaDB Docker Image
```bash
docker pull scylladb/scylla
```

2. Create a ScyllaDB Container
```bash
docker run --name scylla-local -p 9042:9042 -d scylladb/scylla
```

3. Access the CQL Shell
```bash
docker exec -it scylla-local cqlsh
```

After accessing CQL Shell you can create your own keyspace along with table and data. Do mention the same in .env if you have some other names for your keyspace and table name as that of what is mentioned in the repo

## APIs
The REST API to the users app is described below.

### 1. To fetch sorted todos sorted on recent created date along with the implementation of pagination and access to filter using status
#### Request

`GET /todos?page={page}&limit={limit}&status={status}`


### 2. To fetch todo data using id
#### Request

`GET /todos/{id}`


### 3. To create the todo using the todo data from body
#### Request

`POST /todos`


### 4. To update the todo using id and data from body
#### Request

`PUT /todos/{id}`


 ### 5. To delete the todo using id
#### Request

`DEL /todos/{id}`
