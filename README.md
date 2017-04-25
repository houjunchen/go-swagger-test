go generate
protoc --go_out=./models/ --proto_path=./protos/ ./protos/*.proto
