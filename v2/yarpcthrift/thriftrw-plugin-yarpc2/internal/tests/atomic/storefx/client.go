// Code generated by thriftrw-plugin-yarpc
// @generated

package storefx

import (
	"fmt"
	"go.uber.org/fx"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/atomic/storeclient"
)

// Params defines the dependencies for the Store client.
type Params struct {
	fx.In

	Provider yarpc.ClientProvider
}

// Result defines the output of the Store client module. It provides a
// Store client to an Fx application.
type Result struct {
	fx.Out

	Client storeclient.Interface

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// Client provides a Store client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		storefx.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...yarpcthrift.ClientOption) interface{} {
	return func(p Params) (Result, error) {
		client, ok := p.Provider.Client(name)
		if !ok {
			return Result{}, fmt.Errorf("generated code could not retrieve client for %q", name)
		}
		return Result{Client: storeclient.New(client, opts...)}, nil
	}
}
