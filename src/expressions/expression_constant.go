// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package expressions

import (
	"datavalues"
)

type ConstantExpression struct {
	value *datavalues.Value
}

func CONST(v interface{}) IExpression {
	return NewConstantExpression(datavalues.ToValue(v))
}

func NewConstantExpression(v *datavalues.Value) *ConstantExpression {
	return &ConstantExpression{
		value: v,
	}
}

func (e *ConstantExpression) Eval(params IParams) (*datavalues.Value, error) {
	return e.value, nil
}

func (e *ConstantExpression) Walk(visit Visit) error {
	return nil
}