# Nemo
## Getting Started
CQRS Example

### Local Development

Fork this repo for your repo then clone in your local
```
git clone https://github.com/sofyan48/boilerplate.git
```

Get Project Moduls

```
go get github.com/sofyan48/svc_gateway
```

#### Environment Setup
For Development Mode Setup dotenv
```
cp .env.example .env
```
Setting up your local configuration see example
```
SERVER_ADDRESS=0.0.0.0
SERVER_PORT=3000
SERVER_TIMEZONE=Asia/Jakarta

DB_MYSQL_USERNAME=root
DB_MYSQL_PASSWORD=password
DB_MYSQL_HOST=localhost
DB_MYSQL_PORT=3306
DB_MYSQL_DATABASE=db
```

After environment setting then run your server

```
go run src/main.go
```

for building
```
go build src/main.go
```
#### Live Reload
To activate Live Reload install air 
##### on macOS

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
chmod +x /usr/local/bin/air
```

##### on Linux

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x /usr/local/bin/air
```

##### on Windows

```
curl -fLo ~/air.exe \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/windows/air.exe
```

see watcher.conf setting watcher file for air config now

##### Starting Live Reload
go to your project path
```
air -c watcher.conf
```

### Production Mode

#### Dockerizing
Building Image
```
docker build -t boiler
```
Edit Environment In docker-compose.yml then Run Compose
```
docker-compose up
```
Stop Container
```
docker-compose stop
```
Remove your container
```
docker-compose rm -f
```


## How To Use
See the list below for instructions on how to use this boilerplate:
**app/$version**
1. ***Entity*** are a form of modeling of the REST API that you will create
2. ***Repositories*** Is a connection model that you will use in a service that will be created later, the entity plays an important role for modeling data that will be stored in a database
3. ***Service*** Is a logic process that you will create services can also be combined with business logic to facilitate you in managing all services so pay attention to the layering version of the service
4. ***Controller*** an intermediary between routing and service controller regulates all input and input formats of the REST API that you will create
5. ***Routing*** Routing is a mapping of the paths of the REST API that you have designed, routing is available on each controller version layering, this routing will be called on the router that has been created

Plugins and utils are in the ***util*** folder all third-party packages that help you should be stored in this folder, you can choose whether the package is a middleware of your REST API or as a pure supporting utility


## Documentation Format
### Setup Swagger Docs
See Documentation 
[Swag Docs](https://github.com/swaggo/swag)

## How To Contribute
Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.
 1. ***Fork*** the repo on GitHub
 2. ***Clone*** the project to your own machine
 3. ***Commit*** changes to your own branch
 4. ***Push*** your work back up to your fork
 5. Submit a ***Pull request*** so that we can review your changes
