REGISTRY?=consensys/
VERSION?=dev

# User `make dev arg=--build` to rebuild
dev:
	docker-compose -f docker-compose.dev.yml up $(arg)

# stop:
# 	docker-compose stop

build:
	docker build -f Dockerfile -t ${REGISTRY}fc-retrieval-register:${VERSION} .

build-local:
	docker build -f Dockerfile.dev -t ${REGISTRY}fc-retrieval-register:${VERSION} .

build-dev:
	go build -v cmd/register-server/main.go

uselocal:
	cd scripts; bash use-local-repos.sh

useremote:
	cd scripts; bash use-remote-repos.sh

# Have a clean target to match the other repos. This will be called by the integration test 
# system when building the register
clean:
	echo Does nothing
