OPENAPI_SPEC=decta.json

generate-decta:
ifndef VERSION
	$(error VERSION is required. Usage: make generate-decta VERSION=v1)
endif
	curl -o $(OPENAPI_SPEC) https://dapi.decta.com/decta-iss-api.json
	mkdir -p $(VERSION)
	openapi-generator generate \
	  -i $(OPENAPI_SPEC) \
	  -g go \
	  -o $(VERSION) \
	  --additional-properties=packageName=decta$(VERSION),isGoSubmodule=true \
	  --global-property=apiTests=false,modelTests=false,apiExamples=false,modelExamples=false,apiDocs=true,modelDocs=true \
	  --skip-validate-spec
