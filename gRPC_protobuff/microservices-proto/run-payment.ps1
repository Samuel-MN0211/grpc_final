# PowerShell script to generate Go code from shipping.proto
$GITHUB_USERNAME = "ruandg"
$GITHUB_EMAIL = "smerson0211@gmail.com"
$SERVICE_NAME = "shipping"
$RELEASE_VERSION = "v1.2.3"

Write-Host "Installing protobuf Go plugins..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Write-Host "Generating Go source code for $SERVICE_NAME"
if (!(Test-Path "microservices-proto/golang")) {
  New-Item -ItemType Directory -Path "microservices-proto/golang"
}
if (!(Test-Path "microservices-proto/golang/$SERVICE_NAME")) {
  New-Item -ItemType Directory -Path "microservices-proto/golang/$SERVICE_NAME"
}

protoc --go_out=./microservices-proto/golang/$SERVICE_NAME `
  --go_opt=paths=source_relative `
  --go-grpc_out=./microservices-proto/golang/$SERVICE_NAME `
  --go-grpc_opt=paths=source_relative `
  ./microservices-proto/$SERVICE_NAME/*.proto

Write-Host "Generated Go source code files for $SERVICE_NAME"
Get-ChildItem -Path "microservices-proto/golang/$SERVICE_NAME"

Set-Location "microservices-proto/golang/$SERVICE_NAME"
go mod init "github.com/$GITHUB_USERNAME/microservices-proto/golang/$SERVICE_NAME"
go mod tidy

Write-Host "Shipping service Go code generation completed!" -ForegroundColor Green