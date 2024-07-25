golang_spec1.1-rev0 = spec/V1.1-rev0/openapi_spec.yaml
vbr_spec1.1-rev0 = spec/V1.1-rev0/swagger.json

golang_spec1.0-rev1 = spec/V1.0-rev1/openapi_spec.yaml
vbr_spec1.0-rev1 = spec/V1.0-rev1/swagger.json

default: generate 
cleanup:
	@echo "Cleaning up..."
	rm -f ./pkg/client1.0-rev1/types.go
	rm -f ./pkg/client1.0-rev1/client.go
	rm -f ./pkg/client1.1-rev0/types.go
	rm -f ./pkg/client1.1-rev0/client.go


generate: cleanup
	@echo "Generating types..."
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate types -o ./pkg/client1.1-rev0/types.go -package client ${golang_spec1.1-rev0}
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate types -o ./pkg/client1.0-rev1/types.go -package client ${golang_spec1.0-rev1}
	@echo "Generating client..."
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate client -o ./pkg/client1.1-rev0/client.go -package client ${golang_spec1.1-rev0}
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate client -o ./pkg/client1.0-rev1/client.go -package client ${golang_spec1.0-rev1}

convert:
	@echo "Converting spec..."
	cd tools/oapifixer && go build
	./tools/oapifixer/oapifixer -input ${vbr_spec1.1-rev0} -output ${golang_spec1.1-rev0}
	./tools/oapifixer/oapifixer -input ${vbr_spec1.0-rev1} -output ${golang_spec1.0-rev1}
