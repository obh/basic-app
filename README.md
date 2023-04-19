# Prerequisites
 - MySQL >= 8.x.y 
 - Golang 1.19

# Design 
* The application uses a clean architecture to separate concerns. 
* This application uses MySQL as its persistence layer. The interaction with MySQL is handled through the entity framework `ent`. We define the entities in an `ent` schema and use that to generate the code. 
* The application makes a distinction between the database entity and the data transfer object shared with the client. We transform the database entities to these DTOs in the services layer. 
* The application builds on the `services` entity and exposes endpoints to fetch all `services` and to fetch a particular service using `services/:id`. 
* The application responds with standard HTTP response codes back to the client. 
* The application also has integration tests written to ensure correctness.

# Assumptions
* The response of a `services` API represents a single page in a reverse chronological stream of objects.
* The response of a `services/:id` API represents all versions of a service in a reverse chronological stream of objects.
* The `services` API is a GET method, taking at least these four parameters: limit, created_after, created_before and filter_by. 
  * One specifies the `created_after` equal to the object ID of an item to retrieve the newer services created aftr this object. Similarly you specify `created_before` to retrieve the older services created before this service.
  * `limit` A limit on the number of objects to be returned, between 1 and 50. If you do not provide any limit, a default limit of `10` is applied.
  * `filter_by` The filter_by clause is used to filter the services based on their `name` and `description`. The filtering is restricted by the choices offered by MySQL (our persistence layer). We use a `contains` query on MySQL and our results are limited to it. 

* The Response of the `services` API is the body containing the services in a reverse chronological list. 
  * To help with pagination we also respond back with the `size` of the overall result. 
  * We also respond back with the original request which was sent in a separate sub-object called `request`.
  
* The `services/:id` API is a GET method, taking only the service id as a parameter. 
  * This API responds with the Service DTO and the service versions in a reverse chronological stream. 

# Other Assumptions 
* We assume that a service will always have atleast 1 version in serviceversions. 
* We create a mysql index on `id` and the `created_on` field to ensure correctness of data.
* We also create an index on the `service_id` column in serviceversions table to improve query performance.

# Setup

## Step 1
Create and setup the database to run tests. Please provide your hostname, username and db name. I

```
mysql -h localhost -u username -p database_name < dump.sql
```

## Step 2
This service uses MySQL, please provide the mysql dsn on mac using (replace the {parameter} with actual values):
```bash
export MYSQL_DSN="{username}:{password}@tcp({hostname}:3306)/{dbname}?parseTime=true"
```

## Step 3
Run the integration tests

```go
go test ./test/ -v
```

# Run
This service uses the port 1323. Run this service as (ensure you have the env variable `MYSQL_DSN` set):
```go
go run main.go
```

