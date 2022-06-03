# Simple Go WAGI example

Using the WAGI cli, to host a simple go application compiled to wasm using tinygo

Prerequisites:

* Install Tinygo
* Install wagi
* Install binaryen # tinygo should install this

## Troubleshooting

#### error: could not find wasm-opt, set the WASMOPT environment variable to override

You need to install binaryen and make sure wasm-opt.exe|wasm-opt is in your PATH

