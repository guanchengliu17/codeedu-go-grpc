gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:framework/pb

clean:
	rm framework/pb/*.go