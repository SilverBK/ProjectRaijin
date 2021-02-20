package metadata

import (
	"strings"

	"github.com/apache/arrow/go/arrow"
	"github.com/pkg/errors"
	"github.com/silverbk/projectraijin/src/common"
)

type DataTypeFactory struct{}

func (factory *DataTypeFactory) create(tp string) (arrow.Type, error) {
	if strings.ToLower(tp) == "int8" {
		return arrow.INT8, nil
	}

	if strings.ToLower(tp) == "int32" {
		return arrow.INT32, nil
	}

	return arrow.NULL, errors.Wrapf(common.ParsingError, "unsupported type: %s", tp)
}
