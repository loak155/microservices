module github.com/loak155/microservices/services/auth

go 1.21.5

require (
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/joho/godotenv v1.5.1
	github.com/loak155/microservices/proto v0.0.0-20240117044419-dc53153e6de1
	github.com/loak155/microservices/services/user v0.0.0-20240126050935-07fdc1202360
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
	google.golang.org/grpc v1.60.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
