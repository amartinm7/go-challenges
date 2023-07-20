# From your terminal run:
> docker exec -i -t -u root $(docker ps | grep challenge_1-kafka-1 | cut -d' ' -f1) /bin/bash
# $(docker ps | grep docker_kafka | cut -d' ' -f1) - Will return the docker process ID of the Kafka Docker running so you can acces it

# Create a topic
bash> /bin/kafka-topics --create --partitions 4 --bootstrap-server kafka:9092 --topic topic.test.1

# Create a consumer
bash> /bin/kafka-console-consumer --from-beginning --bootstrap-server kafka:9092 --topic=test

# Create a producer
bash> /bin/kafka-console-producer --broker-list kafka:9092 --topic=test