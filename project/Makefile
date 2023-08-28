MQTTCLIENT_BINARY=mqttClientApp
MQTTSERVER_BINARY=mqttServerApp
STORAGESERVER_BINARY=storageServerApp

up:
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_mqtt_client build_mqtt_server build_storage_server
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_mqtt_client:
	@echo "Building mqtt client binary..."
	cd ../mqtt-client && env GOOS=linux CGO_ENABLED=0 go build -o ${MQTTCLIENT_BINARY} ./cmd/api
	@echo "Done!"

build_mqtt_server:
	@echo "Building mqtt server binary..."
	cd ../mqtt-server && env GOOS=linux CGO_ENABLED=0 go build -o ${MQTTSERVER_BINARY} ./cmd/api
	@echo "Done!"

build_storage_server:
	@echo "Building storage server binary..."
	cd ../storage-server && env GOOS=linux CGO_ENABLED=0 go build -o ${STORAGESERVER_BINARY} ./cmd/api
	@echo "Done!"