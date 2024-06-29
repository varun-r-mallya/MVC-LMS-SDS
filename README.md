# A Library Management System 
## written in Go using the MVC architecture

## To run
- Set up MariaDB on your device (only Arch based systems will work, else you will have to replicate the startup script for your OS)
- run `./init.sh` as root
- Install Golang-Migrate from the AUR and provide a connection string to the database in the `./init.sh` script
- rename `.env.example` file to `.env` and paste the contents there and add required credentials
- run the following command to install `air` (hot reload support)
```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
- run `make dev` to start in development mode
- run `make build` to get a binary in `./target/` directory
