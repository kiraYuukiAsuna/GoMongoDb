rmdir /s /q "DBMS/Generated"
mkdir "DBMS/Generated"
mkdir "DBMS/Generated/Message"
mkdir "DBMS/Generated/Service"

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto --go_out=DBMS/Generated --go-grpc_out=DBMS/Generated proto/Message/*
"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto --go_out=DBMS/Generated --go-grpc_out=DBMS/Generated proto/Service/*

move DBMS/Generated/DBMS/Generated/proto DBMS/Generated/proto

rmdir /s /q "DBMS/Generated/DBMS"
rmdir /s /q "DBMS/Generated/Message"
rmdir /s /q "DBMS/Generated/Service"

"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto --cpp_out=Generated --plugin=protoc-gen-grpc=Bin/protobuf/bin/grpc_cpp_plugin.exe --grpc_out=Generated proto/Message/*
"Bin/protobuf/bin/protoc.exe" -I=Bin/protobuf/include -I=proto --cpp_out=Generated --plugin=protoc-gen-grpc=Bin/protobuf/bin/grpc_cpp_plugin.exe --grpc_out=Generated proto/Service/*
