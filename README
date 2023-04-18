# Assumptions
> 1. The API responds with a sorted response for services and versions based on the created_on field?

> 2. can there be a service without a serviceVersion? 
No, assuming that serviceversion will always be present. If not, that service won't show up in the /services API and other API results will be invalid.

# Prerequisites
 - MySQL >= 8.x.y 
 - Golang 1.19

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

