NETWORK_NAME = mynetwork
DEV_CONTAINER_NAME = dev_kontti
MONGODB_IMAGE = mymongoimage
MONGODB_CONTAINER_NAME = local_mongodb

.PHONY: clean start_db

create_network:
	@echo "Creating network: $(NETWORK_NAME)"
	-sudo docker network create -d bridge $(NETWORK_NAME)
	-sudo docker network connect $(NETWORK_NAME) $(DEV_CONTAINER_NAME)

start_db:
	-sudo docker start $(MONGODB_CONTAINER_NAME) ||  make create_db
	-sudo docker network connect $(NETWORK_NAME) $(DEV_CONTAINER_NAME)

stop:
	-sudo docker stop $(MONGODB_CONTAINER_NAME)

create_db: create_network
	@echo "Creating mongodb container"
	-sudo docker build -t $(MONGODB_IMAGE) -f db.Dockerfile .
	-sudo docker run -d --name $(MONGODB_CONTAINER_NAME) $(MONGODB_IMAGE)
	-sudo docker network connect $(NETWORK_NAME) $(MONGODB_CONTAINER_NAME)


clean:
	-@docker stop $$(docker ps -aq) && docker rm $$(docker ps -aq)
	-@docker rmi $$(docker images -aq)
	-@docker volume rm $$(docker volume ls -q)
	-@docker system prune -f
	-@docker volume prune -f
	@echo "Removing Docker networks..."
	-@docker network prune -f
	@echo "Removing Docker images..."
	@echo "Cleaning completed."