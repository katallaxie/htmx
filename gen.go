//go:build generate
// +build generate

package htmx

//go:generate go run cmd/heroicons/main.go
//go:generate go run cmd/sidekickicons/main.go
//go:generate protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/htmx.proto
