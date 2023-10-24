# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOCLEAN = $(GOCMD) clean
GOFMT = $(GOCMD) fmt
GOGET = $(GOCMD) get

# Main API directory
API_DIR = .

# API binary name
API_BINARY = api

all: test build

build:
	$(GOBUILD) -o $(API_BINARY) $(API_DIR)

test:
	$(GOTEST) -v $(API_DIR)/...

clean:
	$(GOCLEAN)
	rm -f $(API_BINARY)

fmt:
	$(GOFMT) $(API_DIR)

docker-build:
	docker build -t todoapi .

docker-run:
	docker run -p 8080:8080 todoapi

# ... (previous Makefile content)

# Docker Compose settings
DOCKER_COMPOSE = docker-compose

docker-compose-build:
	$(DOCKER_COMPOSE) build

docker-compose-up:
	$(DOCKER_COMPOSE) up

docker-compose-down:
	$(DOCKER_COMPOSE) down
