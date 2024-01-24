.PHONY: protoc
protoc:
	for file in $$(find . -name '*.proto' -not -path './third_party/*'); do \
		protoc \
		-I $$(dirname $$file) \
		-I ./third_party \
		--go_out=:$$(dirname $$file) --go_opt=paths=source_relative \
		--go-grpc_out=:$$(dirname $$file) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=:$$(dirname $$file) --grpc-gateway_opt=paths=source_relative \
		$$file; \
	done