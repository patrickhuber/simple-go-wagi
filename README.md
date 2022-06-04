# Simple Go WAGI example

Using the WAGI cli, to host a simple go application compiled to wasm using tinygo

## Prerequisites:

* [Install Tinygo](https://tinygo.org/getting-started/install/)
* [Install wagi](https://github.com/deislabs/wagi/releases)
* [Install binaryen](https://github.com/WebAssembly/binaryen/releases/tag/version_108) # tinygo should install this

## Compiling

```bash
tinygo build -o test.wasm --target wasi main.go
```

## Running

> bash

```bash
# this will start listening on port 3000
wagi -c ./modules.toml
```

> powershell

```powershell
# this will start listening on port 3000
wagi -c .\modules.toml
```

## Troubleshooting

#### error: could not find wasm-opt, set the WASMOPT environment variable to override

You need to install binaryen and make sure wasm-opt.exe|wasm-opt is in your PATH
