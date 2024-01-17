protoc -I. --include_imports --include_source_info --descriptor_set_out=protos/helloworld.pb ./protos/helloworld.proto
protoc -I. --include_imports --include_source_info --descriptor_set_out=protos/book.pb ./protos/book.proto

protoc --go_out=. --go-grpc_out=. protos/book.proto