#!/bin/bash

sleep 120
echo "MySQL is up - executing command"
migrate -path database/migration/ -database "mysql://user:password@tcp(db:3306)/LMS" up

exec "./build/MVC-LMS-SDS"