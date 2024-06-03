build-core:
	cd core && cargo build --target wasm32-unknown-unknown --release
	wasm-snip core/target/wasm32-unknown-unknown/release/core.wasm -o client/core.wasm $$(wasm-objdump -x --section=import core/target/wasm32-unknown-unknown/release/core.wasm | grep func | grep wbindgen | cut -d "<" -f2 | cut -d ">" -f1)

run-client:
	cd client && go run main.go

run:
	$(MAKE) build-core
	$(MAKE) run-client
