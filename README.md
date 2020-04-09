# ORCHESTRATION SAGA PATTERN
Orchestration saga pattern example, this project written in golang

## Concept
![concept](https://github.com/sofyan48/orchestration-pattern-example/raw/master/docs/concept.png)

## Service
- cimol (Notification Service for sms(infobip, wavecell, twilio), email (sendgrid) and firebase)
- cinlog (Logger History, support storage mongo, elasticsearch and AWS S3)
- svc_user (User Service)
- svc_gateway (API Layer Gateway)
- svc_order (Order Service)
- svc_payment (Payment Service)

## Dependecies
- broker (kafka)
- database (cockroachdb)
- mongodb
- elasticsearch

## Getting Started
### Setting environment
See docker-compose.yml and search object environment in the service service block then setup environment
### Running
```bash
docker-compose up
```
## Documentation
This rest documentation using insomnia
```
├── api
│   └── api_layer.json
```
import api_layer.json to insomnia workspace
## Migration Tool
Using golang migrate
### Instalation
See this [mattes/migrate](https://github.com/mattes/migrate) for install instructions

## Service
This project using cockroachdb same driver in postgres for go migrate
### User
#### Migration
```
migrate -source file://path/to/migrations -database postgres://localhost:26257/user up 2
```
#### Testing
**Create User**
```
curl --request POST \
  --url http://localhost/v1/user \
  --header 'content-type: application/json' \
  --data '{
	"first_name":"test",
	"last_name":"ting",
	"email":"mail11@testing.com",
	"handphone":"6281247930699",
	"address":"Address",
	"city":"city",
	"province":"province",
	"district":"district"
}'
```
Figure 1.
![concept](https://github.com/sofyan48/orchestration-pattern-example/raw/master/docs/user/create.png)

**Get User**
```
curl --request GET \
  --url http://localhost/v1/user/89137028-0be0-466b-b97f-1b104ab8e092
```
Figure 2.
![concept](https://github.com/sofyan48/orchestration-pattern-example/raw/master/docs/user/get.png)

### Order
#### Migration
```
migrate -source file://path/to/migrations -database postgres://localhost:26259/order up 2
```
#### Testing
**Create Order**
```
curl --request POST \
  --url http://localhost/v1/order \
  --header 'content-type: application/json' \
  --data '{
	"order_number":"001/ord/sku/IV/2020",
	"uuid_user":"89137028-0be0-466b-b97f-1b104ab8e092",
	"id_order_type": "545204885836038145",
	"id_payment_model": "545204998214516737"
}'
```
Figure 1.
![concept](https://github.com/sofyan48/orchestration-pattern-example/raw/master/docs/order/create.png)

**Get Order**
```
curl --request GET \
  --url http://localhost/v1/order/091a9f65-d56c-4224-b3d7-a160a604ba4d
```
Figure 2.
![concept](https://github.com/sofyan48/orchestration-pattern-example/raw/master/docs/order/get.png)