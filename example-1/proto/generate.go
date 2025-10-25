//go:generate sh -c "protoc -I=. --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. *.proto"

package proto
