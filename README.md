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