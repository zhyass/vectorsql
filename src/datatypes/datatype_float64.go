// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package datatypes

import (
	"fmt"
	"io"
	"math"

	"base/binary"
	"base/errors"
	"datavalues"
)

const (
	DataTypeFloat64Name = "Float64"
)

type Float64DataType struct {
}

func NewFloat64DataType() IDataType {
	return &Float64DataType{}
}

func (datatype *Float64DataType) Name() string {
	return DataTypeFloat64Name
}

func (datatype *Float64DataType) Serialize(writer *binary.Writer, v datavalues.IDataValue) error {
	round := math.Round(datavalues.AsFloat(v)*10000) / 10000
	if err := writer.Float64(round); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (datatype *Float64DataType) SerializeText(writer io.Writer, v datavalues.IDataValue) error {
	if _, err := writer.Write([]byte(fmt.Sprintf("%.4f", datavalues.AsFloat(v)))); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (datatype *Float64DataType) Deserialize(reader *binary.Reader) (datavalues.IDataValue, error) {
	if res, err := reader.Float64(); err != nil {
		return nil, errors.Wrap(err)
	} else {
		return datavalues.MakeFloat(res), nil
	}
}
