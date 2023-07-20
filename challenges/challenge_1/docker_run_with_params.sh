#bin/bash

docker run -p 8000:8000 \
-e POSTGRES_ENABLE_DATABASE=true \
-e POSTGRES_USER=postgres_user \
-e POSTGRES_PASSWORD=postgres_pass \
-e POSTGRES_DB=learning_go_db \
-e POSTGRES_DB_HOST=database \
--network=challenge_1_default \
-it ms-ma-myads

# docker stop $(docker ps -a -q)
# docker rm -f $(docker ps -a -q)
# docker-compose -f docker-compose-all.yml up
# docker-compose -f docker-compose-all.yml down
# docker volume prune -f
# docker volume ls
# docker network prune -f
# docker network ls

# curl -v --location 'http://localhost:8000/health' --header 'Content-Type: application/json'

# everytime we execute a docker-compose we create a network
# to add a docker container to a previous docker-compose network we have to discover that network
# use this command to fill the --network param on the previous script --network=challenge_1_default \
#  docker network ls
