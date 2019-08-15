# Help
.PHONY: default
default:
	@echo "Please specify a build target. The choices are:"
	@echo "    check:            Run lint checks"
	@echo "    compose-up:       Start the service locally using docker-compose"
	@echo "    compose-down:     Stop the service locally using docker-compose"
	@echo "    compose-logs:     Output logs from docker-compose"
	@echo "    format:           Format code using gofmt"
	@echo "    image:            Create Docker image"
	@echo "    test:             Run unit tests"

.PHONY: image
image:
	@echo "============= Creating container image ============="
	./infra/makecmd image

.PHONY: compose-up
compose-up:
	@echo "============= Spinning up services ============="
	./infra/makecmd compose-up

.PHONY: compose-down
compose-down:
	@echo "============= Tearing down services ============="
	./infra/makecmd compose-down

.PHONY: compose-logs
compose-logs:
	@echo "============= Getting compose logs ============="
	./infra/makecmd compose-logs

.PHONY: test
test:
	@echo "============= Running unit tests ============="
	./infra/makecmd test

.PHONY: test-integration
test-integration:
	@echo "============= Running integration tests ============="
	./infra/makecmd test-integration

.PHONY: format
format:
	@echo "============= Formatting project code ============="
	./infra/makecmd fmt

.PHONY: check
check:
	@echo "============= Linting project code ============="
	./infra/makecmd vet

