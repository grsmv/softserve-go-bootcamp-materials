package inmem

import (
	"strconv"
	"strings"
	"time"
)

type inMemDB struct {
	dbs map[string]*inMem
}

func (memDB *inMemDB) Save(dbname string, key string, value interface{}, expire ...int) error {
	db, ok := memDB.dbs[dbname]
	if !ok {
		return errNotExists
	}
	db.set(key, value)
	if len(expire) > 0 {
		delay := time.Duration(expire[0]) * time.Second
		db.expire(key, delay)
	}
	return nil
}

func (memDB *inMemDB) Get(dbname string, key string, receiver interface{}) error {
	db, ok := memDB.dbs[dbname]
	if !ok {
		return errNotExists
	}
	value, err := db.get(key) //race condition
	if err != nil {
		return err
	}

	switch receiver.(type) {
	case *int:
		*(receiver.(*int)), err = intFromRecord(value)
	case *string:
		*(receiver.(*string)), err = stringFromRecord(value)
	case *bool:
		*(receiver.(*bool)), err = boolFromRecord(value)
	case *float64:
		*(receiver.(*float64)), err = floatFromRecord(value)
	default:
		err = errNotImplemented
	}
	return err
}

func (memDB *inMemDB) GetMulti(dbname string, pattern string) ([]string, error) {
	db, ok := memDB.dbs[dbname]
	if !ok {
		return nil, errNotExists
	}
	response, err := db.getMulti(pattern)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(response))
	var str string
	for _, val := range response {
		str, err = stringFromRecord(val)
		if err == nil {
			result = append(result, str)
		}
	}
	return result, nil
}

func (memDB *inMemDB) CloseDB(dbname string) error {
	db, ok := memDB.dbs[dbname]
	if !ok {
		return errNotExists
	}
	waitDB := make(chan struct{})
	db.close(waitDB)
	<-waitDB
	delete(memDB.dbs, dbname)
	return nil
}

func intFromRecord(val interface{}) (int, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return val.(int), nil
	case float32, float64:
		return int(val.(float64)), nil
	case string:
		return strconv.Atoi(val.(string))
	case bool:
		if val.(bool) {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, errNotImplemented
	}
}

func boolFromRecord(val interface{}) (bool, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return val.(int) != 0, nil
	case float32, float64:
		return val.(float64) != 0.0, nil
	case string:
		return strings.ToLower(strings.Trim(val.(string), "\r\n\t ")) == "true", nil
	case bool:
		return val.(bool), nil
	default:
		return false, errNotImplemented
	}
}

func floatFromRecord(val interface{}) (float64, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return float64(val.(int)), nil
	case float32, float64:
		return val.(float64), nil
	case string:
		return strconv.ParseFloat(val.(string), 64)
	case bool:
		if val.(bool) {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		return 0.0, errNotImplemented
	}
}

func stringFromRecord(val interface{}) (string, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(val.(int)), nil
	case float32, float64:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64), nil
	case string:
		return val.(string), nil
	case bool:
		return strconv.FormatBool(val.(bool)), nil
	default:
		return "", errNotImplemented
	}
}
