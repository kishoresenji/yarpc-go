// Code generated by thriftrw-plugin-yarpc
// @generated

package baseservicefx

import (
	"go.uber.org/fx"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/common/baseserviceserver"
)

// ServerParams defines the dependencies for the BaseService server.
type ServerParams struct {
	fx.In

	Handler baseserviceserver.Interface
}

// ServerResult defines the output of BaseService server module. It provides the
// procedures of a BaseService handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []yarpc.TransportProcedure `group:"yarpcfx"`
}

// Server provides procedures for BaseService to an Fx application. It expects a
// baseservicefx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyBaseServiceHandler) baseserviceserver.Interface {
// 			return h
// 		},
// 		baseservicefx.Server(),
// 	)
func Server(opts ...yarpcthrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := baseserviceserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}
