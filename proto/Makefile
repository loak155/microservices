.PHONY: protoc
protoc:
	for file in $$(find . -name '*.proto' -not -path './third_party/*'); do \
		protoc \
		-I . \
		-I ./third_party \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
		$$file; \
	done
