# Closes running services
docker-compose down

# Generates Swag files
export PATH=$(go env GOPATH)/bin:$PATH # Required for swag init
( cd src ; swag init --parseDependency --parseInternal )


# Brings services up
docker-compose up -d --build glofox-task-db

sleep 5 # Database was not ready when gorm attempted to connect

docker-compose up -d --build glofox-task

# Clear Console
printf "\033c"

#Docker logging
docker logs --follow glofox-task