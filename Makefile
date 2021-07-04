all: prepare tool

prepare:
	rm -r ./bin || true
	mkdir ./bin
	mkdir ./bin/lib
	cp -r ./assets/lib ./bin/lib/v1.2.0
	cp ./assets/linux-install.sh ./bin
	cp LICENSE ./bin
	cp README.md ./bin
tool:
	cd ./rust/catrina/src && cargo build --release
	cp ./rust/catrina/target/release/catrina ./bin/
dev:
	cd ./rust/catrina/src && cargo run
	cp ./rust/catrina/target/debug/catrina ./bin/