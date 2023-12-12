rmdir /s /q "DBMS/Generated"
mkdir "DBMS/Generated"
mkdir "DBMS/Generated/Message"
mkdir "DBMS/Generated/Service"

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Message/* --plugin=protoc-gen-go=Bin/protobuf/bin/protoc-gen-go.exe --plugin=protoc-gen-go-grpc=Bin/protobuf/bin/protoc-gen-go-grpc.exe  --go_out=DBMS/Generated --go-grpc_out=DBMS/Generated

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Service/* --plugin=protoc-gen-go=Bin/protobuf/bin/protoc-gen-go.exe --plugin=protoc-gen-go-grpc=Bin/protobuf/bin/protoc-gen-go-grpc.exe  --go_out=DBMS/Generated --go-grpc_out=DBMS/Generated

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Service/* --plugin=protoc-gen-grpc-gateway=Bin/protobuf/bin/protoc-gen-grpc-gateway.exe --grpc-gateway_out DBMS/Generated --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Service/* -I=protoc-gen-openapiv2/options --plugin=protoc-gen-openapiv2=Bin/protobuf/bin/protoc-gen-openapiv2.exe --openapiv2_out apiref/openapiv2 --openapiv2_opt generate_unbound_methods=true

rmdir /s /q "DBMS\Generated\proto
move "DBMS\Generated\DBMS\Generated\proto" "DBMS\Generated\proto"
move "DBMS\Generated\Service\Service.pb.gw.go" "DBMS\Generated\proto\service"
rmdir /s /q "DBMS\Generated\DBMS
rmdir /s /q "DBMS\Generated\Message
rmdir /s /q "DBMS\Generated\Service"

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Message/* --cpp_out=Generated --plugin=protoc-gen-grpc=Bin/protobuf/bin/grpc_cpp_plugin.exe --grpc_out=Generated

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto proto/Service/* --cpp_out=Generated --plugin=protoc-gen-grpc=Bin/protobuf/bin/grpc_cpp_plugin.exe --grpc_out=Generated

