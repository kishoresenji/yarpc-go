// Code generated by thriftrw v1.7.0. DO NOT EDIT.
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package kv

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ResourceDoesNotExist struct {
	Key     string  `json:"key,required"`
	Message *string `json:"message,omitempty"`
}

// ToWire translates a ResourceDoesNotExist struct into a Thrift-level intermediate
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
func (v *ResourceDoesNotExist) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Key), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Message != nil {
		w, err = wire.NewValueString(*(v.Message)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ResourceDoesNotExist struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ResourceDoesNotExist struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ResourceDoesNotExist
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ResourceDoesNotExist) FromWire(w wire.Value) error {
	var err error

	keyIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !keyIsSet {
		return errors.New("field Key of ResourceDoesNotExist is required")
	}

	return nil
}

// String returns a readable string representation of a ResourceDoesNotExist
// struct.
func (v *ResourceDoesNotExist) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}

	return fmt.Sprintf("ResourceDoesNotExist{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this ResourceDoesNotExist match the
// provided ResourceDoesNotExist.
//
// This function performs a deep comparison.
func (v *ResourceDoesNotExist) Equals(rhs *ResourceDoesNotExist) bool {
	if !(v.Key == rhs.Key) {
		return false
	}
	if !_String_EqualsPtr(v.Message, rhs.Message) {
		return false
	}

	return true
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *ResourceDoesNotExist) GetMessage() (o string) {
	if v.Message != nil {
		return *v.Message
	}

	return
}

func (v *ResourceDoesNotExist) Error() string {
	return v.String()
}
