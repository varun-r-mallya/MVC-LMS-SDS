#!/bin/bash

ls -la
pwd
migrate -path database/migration/ -database "mysql://root:password@tcp(db:3307)/LMS" up

# Start the main application
chmod +x MVC-LMS-SDS
exec ./MVC-LMS-SDS