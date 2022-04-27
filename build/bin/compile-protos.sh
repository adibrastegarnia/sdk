#!/bin/sh

proto_path="./api:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src"

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor"
go_import_paths="${go_import_paths},Matomix/controller/v1/controller.proto=github.com/atomix/runtime-api/api/atomix/controller/v1"
go_import_paths="${go_import_paths},Matomix/runtime/v1/runtime.proto=github.com/atomix/runtime-api/api/atomix/runtime/v1"
go_import_paths="${go_import_paths},Matomix/atom/v1/descriptor.proto=github.com/atomix/runtime-api/api/atomix/atom/v1"
go_import_paths="${go_import_paths},Matomix/atom/v1/atom.proto=github.com/atomix/runtime-api/api/atomix/atom/v1"
go_import_paths="${go_import_paths},Matomix/atom/meta/v1/headers.proto=github.com/atomix/runtime-api/api/atomix/atom/meta/v1"
go_import_paths="${go_import_paths},Matomix/atom/meta/v1/object.proto=github.com/atomix/runtime-api/api/atomix/atom/meta/v1"
go_import_paths="${go_import_paths},Matomix/atom/meta/v1/timestamp.proto=github.com/atomix/runtime-api/api/atomix/atom/meta/v1"

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/v1 \
  --doc_opt=markdown,controller.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/runtime/v1,plugins=grpc:api \
  api/atomix/runtime/v1/controller.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/atom/v1 \
  --doc_opt=markdown,descriptor.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/atom/v1,plugins=grpc:api \
  api/atomix/atom/v1/descriptor.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/atom/v1 \
  --doc_opt=markdown,atom.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/atom/v1,plugins=grpc:api \
  api/atomix/atom/v1/atom.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/atom/meta/v1 \
  --doc_opt=markdown,headers.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/atom/meta/v1,plugins=grpc:api \
  api/atomix/atom/meta/v1/headers.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/atom/meta/v1 \
  --doc_opt=markdown,object.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/atom/meta/v1,plugins=grpc:api \
  api/atomix/atom/meta/v1/object.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/atom/meta/v1 \
  --doc_opt=markdown,timestamp.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/atom/meta/v1,plugins=grpc:api \
  api/atomix/atom/meta/v1/timestamp.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/counter/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/counter/v1,plugins=grpc:api \
  api/atomix/counter/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/counter/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/counter/v1,plugins=grpc:api \
  api/atomix/counter/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/election/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/election/v1,plugins=grpc:api \
  api/atomix/election/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/election/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/election/v1,plugins=grpc:api \
  api/atomix/election/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/indexed_map/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/indexed_map/v1,plugins=grpc:api \
  api/atomix/indexed_map/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/indexed_map/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/indexed_map/v1,plugins=grpc:api \
  api/atomix/indexed_map/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/list/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/list/v1,plugins=grpc:api \
  api/atomix/list/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/list/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/list/v1,plugins=grpc:api \
  api/atomix/list/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/lock/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/lock/v1,plugins=grpc:api \
  api/atomix/lock/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/lock/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/lock/v1,plugins=grpc:api \
  api/atomix/lock/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/map/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/map/v1,plugins=grpc:api \
  api/atomix/map/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/map/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/map/v1,plugins=grpc:api \
  api/atomix/map/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/set/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/set/v1,plugins=grpc:api \
  api/atomix/set/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/set/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/set/v1,plugins=grpc:api \
  api/atomix/set/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/topic/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/topic/v1,plugins=grpc:api \
  api/atomix/topic/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/topic/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/topic/v1,plugins=grpc:api \
  api/atomix/topic/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/value/v1 \
  --doc_opt=markdown,manager.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/value/v1,plugins=grpc:api \
  api/atomix/value/v1/manager.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/value/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/runtime-api/api/atomix/value/v1,plugins=grpc:api \
  api/atomix/value/v1/primitive.proto
