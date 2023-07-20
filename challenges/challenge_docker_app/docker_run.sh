#bin/bash

docker run -p 8000:8000 -it my-docker-app

# docker stop $(docker ps -a -q)
# docker rm -f $(docker ps -a -q)