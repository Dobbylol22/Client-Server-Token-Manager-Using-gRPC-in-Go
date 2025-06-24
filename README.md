# Client-Server-Token-Manager-Using-gRPC-in-Go
Implemented and demonstrated the full workflow for building, launching, and interacting with both server and client components, including step-by-step instructions and sample commands for each operation. 

Give the command in the git bash...
//Also attached output_steps file with screenshots and description for it.

Command to build server and client

`./build.sh`

command to Run server:

`./bin/tokenserver`

command to Run client:

`./bin/tokenclient`


command running client- 

create -    `./bin/tokenclient -create -id 1235 -host localhost -port 50051'
write -     `./bin/tokenclient -write -id 1235 -name abcd -low 0 -mid 10 -high 100 -host localhost -port 50051'
read -     `./bin/tokenclient -read -id 1235 -host localhost -port 50051'
drop -     `./bin/tokenclient -drop -id 1235 -host localhost -port 50051'

# Reference code- 
In client >> main.go >>line 15 to 26 used the reference from https://gobyexample.com/command-line-flags  (ratio = 11/75 = 14%)

In client >> main.go >> line  34, 45 used the reference from https://github.com/grpc/grpc-go/blob/0066bf69deb33b0e5bee4de69090c3ef8f6991aa/examples/helloworld/greeter_client/main.go'  (ratio = 12/75 = 15%)

In client >> main.go >> line 48 used the reference from https://github.com/grpc/grpc-go/blob/0066bf69deb33b0e5bee4de69090c3ef8f6991aa/examples/helloworld/greeter_client/main.go' (overall ratio = 29%)

In server >> main.go >> line 1 to 73 used the reference from https://github.com/grpc/grpc-go/blob/0066bf69deb33b0e5bee4de69090c3ef8f6991aa/examples/helloworld/greeter_server/main.go  (Haven't reused the code to calculate the ratio, just used it as for reference.)

In token.go >> line 13 to 17 used the refernce from the project description. (4/172 = 2%)
In token.go used the concepts of Map and RWMutex

The tokens.proto file is referred following the concept of https://github.com/grpc/grpc-go/blob/0066bf69deb33b0e5bee4de69090c3ef8f6991aa/examples/helloworld/helloworld/helloworld.proto.

Please note-  There is no reused code, just used the references and have written as per the requirements.
