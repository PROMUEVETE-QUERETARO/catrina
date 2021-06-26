all: prepare golang rustlang

prepare:
	rm -r ./bin || true
	mkdir ./bin
	mkdir ./bin/tools
	cp -r ./assets/lib ./bin
	cp ./assets/linux-install.sh ./bin
	cp LICENSE ./bin
	cp README.md ./bin
rustlang:
	cd ./rust/catrina-wizard/src && cargo build --release
	cp ./rust/catrina-wizard/target/release/catrina-wizard ./bin/tools/wizard
golang:
	cd cmd/catrina && CGO_ENABLED=0 go build -o ../../bin/catrina main.go
	cd cmd/catrina && CGO_ENABLED=0 go build -o ../../bin/catrina-update update.go