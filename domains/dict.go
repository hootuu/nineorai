package domains

import (
	"database/sql/driver"
	"encoding/json"
	nerr "errors"
	"fmt"
	"github.com/hootuu/gelato/errors"
	"strconv"
)

type Dict map[string]interface{}

var DictErrInvalidValue = errors.Verify("dict value must be a string or a Dict")
var DictErrKeyNotFound = errors.Verify("no such key in this dict")
var DictErrInvalidType = errors.Verify("the value obtained is not of the expected type")

func NewDict(values ...map[string]interface{}) (Dict, *errors.Error) {
	d := make(Dict)
	if len(values) > 0 {
		value := values[0]
		for key, val := range value {
			if err := d.Set(key, val); err != nil {
				return nil, err
			}
		}
	}

	return d, nil
}

func MustNewDict(values ...map[string]interface{}) Dict {
	dict, _ := NewDict(values...)
	return dict
}

func (dict Dict) Exists(key string) bool {
	_, ok := dict[key]
	return ok
}

func (dict Dict) Set(key string, value interface{}) *errors.Error {
	switch v := value.(type) {
	case string:
		dict[key] = v
	case int, int8, int16, int32, int64:
		dict[key] = v
	case uint, uint8, uint16, uint32, uint64:
		dict[key] = v
	case Dict:
		dict[key] = v
	default:
		return DictErrInvalidValue
	}
	return nil
}

func (dict Dict) MustSet(key string, value interface{}) Dict {
	_ = dict.Set(key, value)
	return dict
}

func (dict Dict) MustExists(key ...string) *errors.Error {
	for _, i := range key {
		if !dict.Exists(i) {
			return errors.Verify(fmt.Sprintf("key %s does not exist in meta", i))
		}
	}
	return nil
}

func (dict Dict) GetString(key string) (string, *errors.Error) {
	value, exists := dict[key]
	if !exists {
		return "", DictErrKeyNotFound
	}
	strValue, ok := value.(string)
	if !ok {
		return "", DictErrInvalidType
	}
	return strValue, nil
}

func (dict Dict) GetInt16(key string) (int16, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, nErr := strconv.ParseInt(strValue, 10, 16)
	if nErr != nil {
		return 0, DictErrInvalidType
	}
	return int16(i), nil
}

func (dict Dict) GetInt32(key string) (int32, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, nErr := strconv.ParseInt(strValue, 10, 32)
	if nErr != nil {
		return 0, DictErrInvalidType
	}
	return int32(i), nil
}

func (dict Dict) GetInt64(key string) (int64, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, nErr := strconv.ParseInt(strValue, 10, 64)
	if nErr != nil {
		return 0, DictErrInvalidType
	}
	return i, nil
}

func (dict Dict) GetFloat32(key string) (float32, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	f, nErr := strconv.ParseFloat(strValue, 32)
	if nErr != nil {
		return 0, DictErrInvalidType
	}
	return float32(f), nil
}

func (dict Dict) GetFloat64(key string) (float64, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	f, nErr := strconv.ParseFloat(strValue, 64)
	if nErr != nil {
		return 0, DictErrInvalidType
	}
	return f, nil
}

func (dict Dict) GetBool(key string) (bool, *errors.Error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return false, err
	}
	b, nErr := strconv.ParseBool(strValue)
	if nErr != nil {
		return false, DictErrInvalidType
	}
	return b, nil
}

func (dict Dict) GetDict(key string) (Dict, *errors.Error) {
	value, exists := dict[key]
	if !exists {
		return nil, DictErrKeyNotFound
	}
	dictValue, ok := value.(Dict)
	if !ok {
		return nil, DictErrInvalidType
	}
	return dictValue, nil
}

func (dict *Dict) Scan(value interface{}) error {
	if value == nil {
		dict = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, dict)
	default:
		return nerr.New("invalid type for Dict")
	}
}

func (dict Dict) Value() (driver.Value, error) {
	return json.Marshal(dict)
}
