This repo is just me keeping track of what I have done learning Go and reminding me what I intend to do

[![hgaard](https://circleci.com/gh/hgaard/simple-http-go.svg?style=shield)](https://circleci.com/gh/hgaard/workflows/simple-http-go)

 - [x] Simple web server (helloserver.go)
 - [x] Api server
    - [x] Api that returns something in JSON (apiserver.go)
    - [x] Add logs (Logrus and seq hooks)
    - [x] Add test (apiserver_test.go)
 - [x] Request logger + log endpoint
    - [x] Simple middleware
    - [x] Capture all calls to all endpoints in "list" (Request recorder holds all requests)
    - [x] Add endpoint to return "list" of all calls (logs endpoint added)
 - [x] CI/CD
    - [x] CircleCI
    - [x] Add CircleCI badge
 - [x] Queue + worker role (rabbitMQ + worker directory)
    - [ ] update docker compose 
    >```docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management```
 - [ ] DB 
    - [ ] Store messages from worker
    - [ ] Read messages from db in log server
 - [ ] Add auth
    - [ ] Google/Facebook/Microsoft
 - [ ] Add Swagger
 - [ ] CLI