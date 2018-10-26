// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package weather

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Weather_Check_Args represents the arguments for the Weather.check function.
//
// The arguments for check are sent and received over the wire as this struct.
type Weather_Check_Args struct {
}

// ToWire translates a Weather_Check_Args struct into a Thrift-level intermediate
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
func (v *Weather_Check_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Weather_Check_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Weather_Check_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Weather_Check_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Weather_Check_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Weather_Check_Args
// struct.
func (v *Weather_Check_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Weather_Check_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Weather_Check_Args match the
// provided Weather_Check_Args.
//
// This function performs a deep comparison.
func (v *Weather_Check_Args) Equals(rhs *Weather_Check_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "check" for this struct.
func (v *Weather_Check_Args) MethodName() string {
	return "check"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Weather_Check_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Weather_Check_Helper provides functions that aid in handling the
// parameters and return values of the Weather.check
// function.
var Weather_Check_Helper = struct {
	// Args accepts the parameters of check in-order and returns
	// the arguments struct for the function.
	Args func() *Weather_Check_Args

	// IsException returns true if the given error can be thrown
	// by check.
	//
	// An error can be thrown by check only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for check
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// check into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by check
	//
	//   value, err := check(args)
	//   result, err := Weather_Check_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from check: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*Weather_Check_Result, error)

	// UnwrapResponse takes the result struct for check
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if check threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Weather_Check_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Weather_Check_Result) (string, error)
}{}

func init() {
	Weather_Check_Helper.Args = func() *Weather_Check_Args {
		return &Weather_Check_Args{}
	}

	Weather_Check_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Weather_Check_Helper.WrapResponse = func(success string, err error) (*Weather_Check_Result, error) {
		if err == nil {
			return &Weather_Check_Result{Success: &success}, nil
		}

		return nil, err
	}
	Weather_Check_Helper.UnwrapResponse = func(result *Weather_Check_Result) (success string, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Weather_Check_Result represents the result of a Weather.check function call.
//
// The result of a check execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Weather_Check_Result struct {
	// Value returned by check after a successful execution.
	Success *string `json:"success,omitempty"`
}

// ToWire translates a Weather_Check_Result struct into a Thrift-level intermediate
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
func (v *Weather_Check_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Weather_Check_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Weather_Check_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Weather_Check_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Weather_Check_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Weather_Check_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
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
		return fmt.Errorf("Weather_Check_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Weather_Check_Result
// struct.
func (v *Weather_Check_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("Weather_Check_Result{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Weather_Check_Result match the
// provided Weather_Check_Result.
//
// This function performs a deep comparison.
func (v *Weather_Check_Result) Equals(rhs *Weather_Check_Result) bool {
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Weather_Check_Result) GetSuccess() (o string) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "check" for this struct.
func (v *Weather_Check_Result) MethodName() string {
	return "check"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Weather_Check_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
