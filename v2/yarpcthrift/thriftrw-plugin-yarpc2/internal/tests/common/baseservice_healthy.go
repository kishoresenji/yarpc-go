// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package common

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// BaseService_Healthy_Args represents the arguments for the BaseService.healthy function.
//
// The arguments for healthy are sent and received over the wire as this struct.
type BaseService_Healthy_Args struct {
}

// ToWire translates a BaseService_Healthy_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BaseService_Healthy_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a BaseService_Healthy_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BaseService_Healthy_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BaseService_Healthy_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BaseService_Healthy_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a BaseService_Healthy_Args
// struct.
func (v *BaseService_Healthy_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("BaseService_Healthy_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this BaseService_Healthy_Args match the
// provided BaseService_Healthy_Args.
//
// This function performs a deep comparison.
func (v *BaseService_Healthy_Args) Equals(rhs *BaseService_Healthy_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "healthy" for this struct.
func (v *BaseService_Healthy_Args) MethodName() string {
	return "healthy"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *BaseService_Healthy_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// BaseService_Healthy_Helper provides functions that aid in handling the
// parameters and return values of the BaseService.healthy
// function.
var BaseService_Healthy_Helper = struct {
	// Args accepts the parameters of healthy in-order and returns
	// the arguments struct for the function.
	Args func() *BaseService_Healthy_Args

	// IsException returns true if the given error can be thrown
	// by healthy.
	//
	// An error can be thrown by healthy only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for healthy
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// healthy into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by healthy
	//
	//   value, err := healthy(args)
	//   result, err := BaseService_Healthy_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from healthy: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(bool, error) (*BaseService_Healthy_Result, error)

	// UnwrapResponse takes the result struct for healthy
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if healthy threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := BaseService_Healthy_Helper.UnwrapResponse(result)
	UnwrapResponse func(*BaseService_Healthy_Result) (bool, error)
}{}

func init() {
	BaseService_Healthy_Helper.Args = func() *BaseService_Healthy_Args {
		return &BaseService_Healthy_Args{}
	}

	BaseService_Healthy_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	BaseService_Healthy_Helper.WrapResponse = func(success bool, err error) (*BaseService_Healthy_Result, error) {
		if err == nil {
			return &BaseService_Healthy_Result{Success: &success}, nil
		}

		return nil, err
	}
	BaseService_Healthy_Helper.UnwrapResponse = func(result *BaseService_Healthy_Result) (success bool, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// BaseService_Healthy_Result represents the result of a BaseService.healthy function call.
//
// The result of a healthy execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type BaseService_Healthy_Result struct {
	// Value returned by healthy after a successful execution.
	Success *bool `json:"success,omitempty"`
}

// ToWire translates a BaseService_Healthy_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BaseService_Healthy_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueBool(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("BaseService_Healthy_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a BaseService_Healthy_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BaseService_Healthy_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BaseService_Healthy_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BaseService_Healthy_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("BaseService_Healthy_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a BaseService_Healthy_Result
// struct.
func (v *BaseService_Healthy_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("BaseService_Healthy_Result{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this BaseService_Healthy_Result match the
// provided BaseService_Healthy_Result.
//
// This function performs a deep comparison.
func (v *BaseService_Healthy_Result) Equals(rhs *BaseService_Healthy_Result) bool {
	if !_Bool_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *BaseService_Healthy_Result) GetSuccess() (o bool) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "healthy" for this struct.
func (v *BaseService_Healthy_Result) MethodName() string {
	return "healthy"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *BaseService_Healthy_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
