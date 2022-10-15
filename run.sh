# Closes running services
docker-compose down

# Brings services up
docker-compose up -d --build glofox-task-db

sleep 5 # Database was not ready when gorm attempted to connect

docker-compose up -d --build glofox-task


#Docker logging
docker logs --follow glofox-task