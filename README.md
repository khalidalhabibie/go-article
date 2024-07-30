# GoArticle

GoArticle is a web-based application designed to facilitate article management. It allows users to create, update, delete, and view articles through a secure, token-based authentication system. Users can register and login to manage their articles, making it a versatile platform for content creation and management.



## Approach
Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

 ![golang clean architecture](clean-arch.png)

### This app has :

- RESTful endpoint for asset's CR operation
- JSON formatted response

&nbsp;

### Tech Stack used to build this app :

- [GoFiber](https://gofiber.io/) web framework<br/>
- [PostgreSQL](https://www.postgresql.org/) as database<br/>
- [gorm](https://gorm.io/index.html) as ORM library <br/>
- [Redis](https://redis.io/) as cache <br/>
- [Docker](https://www.docker.com/)as containerization platform <br/>
- [Docker Compose](https://docs.docker.com/compose/) as container orchestration frameworks.<br/>
- [Prometheus](https://prometheus.io/) monitoring and alerting<br/>
- [Grafana](https://grafana.com/) for to compose observability dashboards with everything from Prometheus<br/>
- [swagger](https://swagger.io/) as Documentation<br/>

&nbsp;

### Requirement :

- [go](https://go.dev/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [go migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) (CLI)
- [swag-go](https://github.com/swaggo/swag)(CLI)(for development process, to regenerate)(optional)
- [mockery](https://github.com/vektra/mockery)(CLI) (for development process, to mock new interface)(optional)
  &nbsp;

## Install and Run 🙌👨‍💻🚀

### my os : Ubuntu 22.04 LTS

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
# install swagger cli
$ go install github.com/swaggo/swag/cmd/swag@latest

#install go migrate
# if failed, please check( https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md ) for your device
$ go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#move to directory
$ cd workspace

# Clone
$ git clone git@github.com:khalidalhabibie/GoArticle.git

#move to project
$ cd GoArticle

# rename file env.example
$ mv env.example .env

# Build the docker image first
$ make build

# Run the application
$ make run

# check if the containers are running
$ sudo docker ps

# migration database and seed data
$ make migrate.up

# Stop the application
$ make stop

```

### Prometheus UI:

http://localhost:9090

### Grafana UI

http://localhost:3000 \
username: admin \
password: admin

### Swagger UI:

http://localhost:8081/swagger/index.html
