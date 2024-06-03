package main

import (
	"context"
	_ "embed"
	"time"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero/api"
)

//go:embed core.wasm
var coreWASM []byte

var PASSWORD_TO_TEST = "RFYGWFN*$XW*9i4385c48uBEFB_!"

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmData{
				Data: coreWASM,
			},
		},
	}

	extismConfig := extism.PluginConfig{}

	getTimeFunc := extism.NewHostFunctionWithStack("unix_time_milliseconds_imported", func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
		stack[0] = uint64(time.Now().UnixMilli())
	}, []api.ValueType{}, []api.ValueType{api.ValueTypeI64})
	getTimeFunc.SetNamespace("zxcvbn")

	plugin, err := extism.NewPlugin(context.Background(), manifest, extismConfig, []extism.HostFunction{getTimeFunc})
	if err != nil {
		panic(err)
	}

	_, resp, err := plugin.CallWithContext(context.Background(), "strength_for_password", []byte(PASSWORD_TO_TEST))
	if err != nil {
		panic(err)
	}

	print(string(resp))
}
