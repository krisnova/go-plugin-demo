
# Run the official golang:1.8 docker container.
docker:
	docker run -i -t -v ${GOPATH}/src/github.com/kris-nova/go-plugin-demo:/go/src/github.com/kris-nova/go-plugin-demo -w /go/src/github.com/kris-nova/go-plugin-demo --rm golang:1.8

# Build the Go plugins for the demonstration, as well as compile the main program.
build:
	@echo ""
	@echo "---[Compiling]-----------------------------------------------------------"
	@rm -f plugins/plugin1.so
	go build -buildmode=plugin -o plugins/plugin1.so plugins/plugin1.go
	@rm -f plugins/plugin2.so
	go build -buildmode=plugin -o plugins/plugin2.so plugins/plugin2.go
	go build -o demo main.go
	@touch *
	@echo "Now you can change your plugin at runtime by setting PLUGIN_NUMBER=n"
	@echo "-------------------------------------------------------------------------"
	@echo ""

# Run the main program.
run:
	@echo ""
	@echo "---[Demonstrating]-------------------------------------------------------"
	./demo
	@echo "-------------------------------------------------------------------------"
	@echo ""

docker-ubuntu:
	docker run -i -t -v ${GOPATH}/src/github.com/kris-nova/go-plugin-demo:/go/src/github.com/kris-nova/go-plugin-demo -w /go/src/github.com/kris-nova/go-plugin-demo --rm ubuntu:16.04
