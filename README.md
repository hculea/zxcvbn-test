## Background

This repository exists in order to provide easy functional testing for https://github.com/shssoichiro/zxcvbn-rs/pull/83

## Prerequisites
* [Go toolchain](https://go.dev/doc/install)
* [Rust toolchain](https://www.rust-lang.org/tools/install) 
* `wasm-objdump` and `wasm-snip`: see [installation instructions](https://github.com/WebAssembly/wabt)

## Testing environment

This project is made of two components:
* `core`, a Rust WASM library that calls the `zxcvbn` function of the `zxcvbn` crate. It exposes a function `strength_for_password` consuming a string, and returns the `guesses_log10` attribute of its entropy.
* `client`, a Go app consuming the WASM-compiled `core`. It injects the `unix_time_milliseconds_imported` in the environment, which is called within the `zxcvbn` library, imported in `core`'s `Cargo.toml` with the `custom_wasm_env` feature on.

## Testing steps

From the root of the repository, run
```go
make run
```
This should output the entropy score of the `PASSWORD_TO_TEST` var defined in `client/main.go`.

Not injecting the function, or importing the crate without the `custom_wasm_env` will result in errors.