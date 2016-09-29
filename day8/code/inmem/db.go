package inmem

import "time"

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

func intFromRecord(val interface{}) (int, error) {
	return 0, errNotImplemented
}

func boolFromRecord(val interface{}) (bool, error) {
	return false, errNotImplemented
}

func floatFromRecord(val interface{}) (float64, error) {
	return 0.0, errNotImplemented
}

func stringFromRecord(val interface{}) (string, error) {
	return "", errNotImplemented
}
