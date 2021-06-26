all: prepare golang rustlang

prepare:
	rm -r ./bin || true
	mkdir ./bin
	mkdir ./bin/tools
rustlang:
	cd ./rust/catrina-wizard/src && cargo build --release
	cp ./rust/catrina-wizard/target/release/catrina-wizard ./bin/tools/wizard
golang:
	cd cmd/catrina && CGO_ENABLED=0 go build -o ../../bin/catrina main.go
	cd cmd/catrina && CGO_ENABLED=0 go build -o ../../bin/catrina-update update.go