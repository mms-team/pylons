package types

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("Can't convert %v to float64", v.Type())
		}
	}
}

// Actualize creates a (key, value) list from ParamList
func (dpm DoubleParamList) Actualize(ec CelEnvCollection) (DoubleKeyValueList, error) {
	// We don't have the ability to do random numbers in a verifiable way rn, so don't worry about it
	var m []DoubleKeyValue
	for _, param := range dpm {
		var valDec sdk.Dec
		var err error

		if len(param.Program) > 0 {
			var val float64
			val, err = ec.EvalFloat64(param.Program)
			valDec, _ = sdk.NewDecFromStr(fmt.Sprintf("%v", val))
		} else {
			valDec, err = param.WeightTable.Generate()
		}
		if err != nil {
			return m, err
		}
		m = append(m, DoubleKeyValue{
			Key:   param.Key,
			Value: valDec,
		})
	}
	return m, nil
}
