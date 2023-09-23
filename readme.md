### Lending Service

#### Description
Simple service to record debt from another person

#### Docs endpoint
[postman](https://documenter.getpostman.com/view/11693224/2s9Y5SW5ro)

#### How to run
1. go mod tidy
2. go run main.go

#### Dependency
1. postgres database

#### Expose endpoint to public
1. [ngrok](https://ngrok.com/)
2. [localtunnel](https://theboroer.github.io/localtunnel-www/ | lt --port 8080 --subdomain lending-service