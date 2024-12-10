package domains

import (
	"errors"
	"strconv"
)

type Dict map[string]interface{}

var DictErrInvalidValue = errors.New("dict value must be a string or a Dict")
var DictErrKeyNotFound = errors.New("no such key in this dict")
var DictErrInvalidType = errors.New("the value obtained is not of the expected type")

func NewDict(values ...map[string]interface{}) (Dict, error) {
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

func (dict Dict) Exists(key string) bool {
	_, ok := dict[key]
	return ok
}

func (dict Dict) Set(key string, value interface{}) error {
	switch v := value.(type) {
	case string:
		dict[key] = v
	case Dict:
		dict[key] = v
	default:
		return DictErrInvalidValue
	}
	return nil
}

func (dict Dict) GetString(key string) (string, error) {
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

func (dict Dict) GetInt16(key string) (int16, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(strValue, 10, 16)
	if err != nil {
		return 0, DictErrInvalidType
	}
	return int16(i), nil
}

func (dict Dict) GetInt32(key string) (int32, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return 0, DictErrInvalidType
	}
	return int32(i), nil
}

func (dict Dict) GetInt64(key string) (int64, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, DictErrInvalidType
	}
	return i, nil
}

func (dict Dict) GetFloat32(key string) (float32, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(strValue, 32)
	if err != nil {
		return 0, DictErrInvalidType
	}
	return float32(f), nil
}

func (dict Dict) GetFloat64(key string) (float64, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return 0, DictErrInvalidType
	}
	return f, nil
}

func (dict Dict) GetBool(key string) (bool, error) {
	strValue, err := dict.GetString(key)
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(strValue)
	if err != nil {
		return false, DictErrInvalidType
	}
	return b, nil
}

func (dict Dict) GetDict(key string) (Dict, error) {
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
