This project is written purely in Go and used PostgreSQL for its database needs .Entirely Dockerized this project using dockerfile and docker-compose.yml
Frameworks used

Gin-Gonic:This whole project is fully completed using gin-gonic which is a popular go framework for rapid web development

 go get -u github.com/gin-gonic/gin

Database used:


PostgreSQL:This project mainly used PostgreSQL as Database with the help of ORM tool named GORM.It provides better and simplified forms of queries for better understanding

go get -u gorm.io/gorm

go get -u gorm.io/driver/postgres

commands to run using go run:

go run main.go

Technologies and tools used

Server : GO Framework : GIN , Database : PSQL , Authentication : JWT , Payment Gateway : stripe , Container : Docker
Run On local machine

clone this project

git clone https://github.com/mohamedmuhsinJ/shopifykart

open shopifykart Directory
cd shopifykart


download dependencies

go get

Run

go run main.go

app is listening on port 3000


For [api documentation](https://documenter.getpostman.com/view/24747045/2s93RQTZS7).
