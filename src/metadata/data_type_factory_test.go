package metadata

import (
	"testing"

	"github.com/apache/arrow/go/arrow"
	"github.com/pkg/errors"
	"github.com/silverbk/projectraijin/src/common"
)

func TestInt8GenerationFromLower(t *testing.T) {
	dataTypeFactory := DataTypeFactory{}
	tpName := "int8"
	expected := arrow.INT8

	result, err := dataTypeFactory.create(tpName)

	common.AssertNil(t, err)
	common.AssertEqual(t, expected, result)
}

func TestInt8GenerationFromUpperCase(t *testing.T) {
	dataTypeFactory := DataTypeFactory{}
	tpName := "INT8"
	expected := arrow.INT8

	result, err := dataTypeFactory.create(tpName)

	common.AssertNil(t, err)
	common.AssertEqual(t, expected, result)
}

func TestInt32GenerationFromLowerCase(t *testing.T) {
	dataTypeFactory := DataTypeFactory{}
	tpName := "int32"
	expected := arrow.INT32

	result, err := dataTypeFactory.create(tpName)

	common.AssertNil(t, err)
	common.AssertEqual(t, expected, result)
}

func TestInt32GenerationFromUpperCase(t *testing.T) {
	dataTypeFactory := DataTypeFactory{}
	tpName := "INT32"
	expected := arrow.INT32

	result, err := dataTypeFactory.create(tpName)

	common.AssertNil(t, err)
	common.AssertEqual(t, expected, result)
}

func TestIntUnsupported(t *testing.T) {
	dataTypeFactory := DataTypeFactory{}
	tpName := "unkown"
	expectedMessage := errors.Wrapf(common.ParsingError, "unsupported type: %s", tpName)

	result, err := dataTypeFactory.create(tpName)

	common.AssertEqual(t, arrow.NULL, result)
	common.AssertNotNil(t, err)
	errors.Is(err, common.ParsingError)
	common.AssertEqual(t, expectedMessage.Error(), err.Error())
}
