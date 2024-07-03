# A Library Management System 
## written in Go using the MVC architecture

## To run
- Set up MariaDB on your device (only Arch based systems will work, else you will have to replicate the startup script for your OS)
- run `./init.sh` as root (ONLY FOR ARCH)
- If not Arch, then configure virtual proxy according to your distro and map it to port 8080 with a domain name of your choice.
- Then add the domain name to the `/etc/hosts` file
- Also install `go migrate` and add connection strings to your MySQL / MariaDB database in the format given below and add it to the `Makefile` as shown below:
```bash
migration_up:
		@read -p "Enter version: " v; \
		migrate -path database/migration/ -database "mysql://username:password@tcp(localhost:port<usually 3306>)/databasename?" -verbose up	$$v
migration_down:
		@read -p "Enter version: " v; \
		migrate -path database/migration/ -database "mysql://username:password@tcp(localhost:port<usually 3306>)/databasename?" -verbose down $$v
migration_fix:
		@read -p "Enter version: " v; \
		migrate -path database/migration/ -database "mysql://username:password@tcp(localhost:port<usually 3306>)/databasename?" force $$v
```
- Install Golang-Migrate from the AUR and provide a connection string to the database in the `./init.sh` script
- run `make migration_up` to run the migrations and press enter to up all versions.
- rename `.env.example` file to `.env` and paste the contents there and add required credentials
- run the following command to install `air` (hot reload support)
```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
- run `make dev` to start in development mode
- run `make build` to get a binary in `./target/` directory
- in `.env`, use mode `dev` to get Logs, use mode `prod` to get only system critical/spotlit logs

## Docker Specific instructions

- `docker build -t xeonlib1 .`
- `docker run -p 8080:8080 xeonlib1`