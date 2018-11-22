# todo-grpc

this is a training project for understanding the concept and the feasibility of rest-grpc gateway.
Interfaces of rest-gateway and grpc will be generated from Protocol buffers.  
 

Install
========================================

first install go and protobuf 
```
brew install go
brew install protobuf
```
git clone this project in your go src folder. e.g. in ~/go/src
```
cd ~/go/src
git clone ....
```
install dependencies
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u moul.io/protoc-gen-gotemplate

go get github.com/gogo/protobuf/protoc-gen-gogofast 
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto

go get github.com/oklog/ulid
go get upper.io/db.v3
go get upper.io/db.v3/sqlite

```

then use regenerateProto.sh to generate code for gateway and gRPC
```
regenerateProto.sh
```

start gateway
========================================

```
go run /cmd/gateway/main.go
```

start gRPC server
========================================
there two way to start the gRPC server. either via go or via java

### start gRPC server via go
```
go run /cmd/grpcserer/main.go
```

### start gRPC server via java

```
./gradlew build
./gradlew idea
./gradlew runserver
```
test link
========================================
[http://localhost:8080/todos](http://localhost:8080/todos)
