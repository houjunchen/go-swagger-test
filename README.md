#### Generate model from proto

```Shell
protoc --go_out=./models/ --proto_path=./protos/ ./protos/*.proto
```

#### Generate `swagger.json` from code

```shell
go generate
```

