# Client-Server API

## First project for the GoExpert course from the [FullCycle](https://fullcycle.com.br/) community

That's a very basic project that aims to practice some important concepts and packages from de standard library of the Go programming language:

- Create and http server
- Make a request to a server
- Work with a dabase
- Work with contexts

## Run this project locally

### Requirements
- Go ^1.19

1. Clone this repo in the desired folder
```bash
git clone https://github.com/vctaragao/client-server-api.git
```

2. Install project dependencies
```bash
cd <path-to-project>/client-server-api
go mod tidy
```

3. Open a terminal on the project folder and start the server
```bash
cd <path-to-project>/client-server-api
go run server/main.go
```

4. Open a new terminal on the project folder run the client
```bash
cd <path-to-project>/client-server-api
go run client/main.go
```
