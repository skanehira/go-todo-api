.PHONY: generate
generate:
	@go generate ./...

.PHONY: doc
doc:
	@docker run --rm -d -p 8888:8080 -v $$PWD:/tmp -e SWAGGER_FILE=/tmp/ent/openapi.json swaggerapi/swagger-editor
