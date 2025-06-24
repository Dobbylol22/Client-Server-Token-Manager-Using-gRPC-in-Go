go build -o ./bin/tokenclient ./client
./bin/tokenclient -create -id 1234 -host localhost -port 50051
./bin/tokenclient -write -id 1234 -name preths -low 0 -mid 10 -high 100 -host localhost -port 50051 
./bin/tokenclient -read -id 1234 -host localhost -port 50051
./bin/tokenclient -drop 1234 -host localhost -port 50051