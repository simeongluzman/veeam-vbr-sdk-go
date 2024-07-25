golang_spec1.1-rev0 = spec/V1.1-rev0/openapi_spec.yaml
vbr_spec1.1-rev0 = spec/V1.1-rev0/swagger.json

golang_spec1.0-rev1 = spec/V1.0-rev1/openapi_spec.yaml
vbr_spec1.0-rev1 = spec/V1.0-rev1/swagger.json

default: generate 
cleanup:
	@echo "Cleaning up..."
	rm -f ./pkg/client/types.go
	rm -f ./pkg/client/client.go

generate: cleanup
	@echo "Generating types..."
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate types -o ./pkg/client/types.go -package client ${golang_spec}
	@echo "Generating client..."
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate client -o ./pkg/client/client.go -package client ${golang_spec}

convert:
	@echo "Converting spec..."
	cd tools/oapifixer && go build
	./tools/oapifixer/oapifixer -input ${vbr_spec} -output ${golang_spec}
