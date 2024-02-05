# split-it-backend
backend for split it written in Go

## Running locally
1. Fill in the necessary ENV vars and create a file named `.env` in the root of the repo
```dotenv
SERVER_PORT= # eg. 8080
PG_CONN_STRING= # eg. "postgres://postgres:postgres@localhost:5435/splitthat"
FRONTEND_REDIRECT_URL= # eg. "http://localhost:5173"
```
2. Run the following commands
```shell
$ git clone git@github.com:ericmiranda7/split-it-frontend.git
$ go mod tidy
$ go install github.com/cosmtrek/air@latest # a lovely livereload server
$ air
```

## Running in the cloud
Google cloud's Google Cloud Run is used for running the app as a container.  
Google Cloud Build is triggered on a push to the `prod` branch.  
The app is available at https://split-it-backend-6vxkdevw6q-de.a.run.app