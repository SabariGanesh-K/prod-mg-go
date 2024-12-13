
# Tech Stack and App details

- Go Gin  [https://gin-gonic.com](https://gin-gonic.com)

- PostgreSQL  - DOCKER - postgres:16.3-alpine3.20

- Redis - DOCKER  -  redis:7-alpine

- RabbitMQ - Docker composed - rabbitmq:3-management

- Used golang-migrate for Migrations 

- Used go mock for Mocking in API testing on DB Store 

- [Zerolog](github.com/rs/zerolog)  for Advanced Logging

-  Amqp library for RabbitMQ interaction  [github.com/streadway/amqp](github.com/streadway/amqp)

- AWS SDK for S3

# Documentation



- Fork the repo
- Ensure you have go installed
- Port 8083 is used for API Backend. Adjust accordingly if being altered in imageprocessor_microservice as well .

## PRE SETUP

Install make to utilize make file commands.

Postgres image 

```bash
make postgres

````
Redis Local 

```bash
make postgres

````

RabbitMQ docker script

```bash
make rabbitmq

````

Open a Public S3 Bucket  from AWS


Setup environment  [appp.env]
```bash
ENVIRONMENT=development
DB_SOURCE=postgresql://root:secret@localhost:5432/prod-mgm?sslmode=disable
DB_DRIVER=postgres
MIGRATION_URL=file://db/migration
HTTP_SERVER_ADDRESS=0.0.0.0:8083
AWS_ACCESS_KEY_ID= <access key >
AWS_SECRET_ACCESS_KEY= <secret access >
AWS_REGION=us-east-1
AWS_BUCKET_NAME= <bucket name >
REDIS_ADDRESS=0.0.0.0:6379
RABBITMQ_URL="amqp://guest:guest@localhost:5672/"
````

## DB and Migrations
Ensure Postgresql container is running

```bash
make createdb

````
```bash
make migrateup

````

## Execution

Ensure you have go installed
```bash
go mod tidy

````

NOW OPEN 2 TERMINALS 

Terminal 1:- Main Backend code
```bash
go run main.go

````
![image](https://github.com/user-attachments/assets/399233ae-4ca6-4fb6-9b6b-9ffb646f0939)

Terminal 2:- Image Processing Microservice
```bash
cd ./imageprocessor_microservice

````
```bash
go run main.go

````
![image](https://github.com/user-attachments/assets/1dcf191b-21c1-4694-b5c5-bfe109a31a26)


## Execution
```bash
make testfull

````
