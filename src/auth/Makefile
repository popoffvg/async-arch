

build: 
	docker build -f src/auth/Dockerfile .

generate-server:
	oapi-codegen -config api/oapi/config.yaml api/oapi/swagger.yaml

.PHONY: ent
ent:
	go generate ./ent

proto:
	
