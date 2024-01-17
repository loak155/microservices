.PHONY: protoc
protoc:
	protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user/*.proto
	protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auth/*.proto
	protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative article/*.proto
	# protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt==paths=source_relative user/*.proto
	# protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt==paths=source_relative auth/*.proto
	# protoc -I. -I./third_party --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt==paths=source_relative article/*.proto
