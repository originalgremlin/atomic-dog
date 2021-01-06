## devops: create a dockerized build environment
devops: $(shell find devops/*)
	docker build -t atomic-dog/devops devops
	touch $(shell find devops)

# makefile
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
