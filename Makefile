.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/video/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/video/cache.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/video/task.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/video/build.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/video/notification.proto
