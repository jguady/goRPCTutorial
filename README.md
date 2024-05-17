# A Todo CRUD app using gRPC

This is just a sample to help me learn gRPC. 

# Service Methods

- `GetTodoItems`
Gets a list of Items given their Id's
- `CreateTodoItem` Creates a TodoItem and returns the Id
- `DeleteTodoItem` Deletes a TodoItem given an Id
- `UpdateTodoItem` Updates a todo item


# Generating Protobuf Code
```shell
#cd into the proto folder and run the following
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/TodoItem.proto
```
### Purpose of Each Option
`--go_out=.`: Generates Go source files for the protocol buffers messages and outputs them to the current directory.

`--go_opt=paths=source_relative`: Ensures that the generated files are placed in the directory structure relative to the source .proto file. For example, if the .proto file is in routeguide, the generated files will also be placed in the routeguide directory.

`--go-grpc_out=.`: Generates Go source files for the gRPC services defined in the .proto file and outputs them to the current directory.

`--go-grpc_opt=paths=source_relative`: Ensures that the generated gRPC files are placed in the directory structure relative to the source .proto file.