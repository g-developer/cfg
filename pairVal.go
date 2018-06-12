package cfg

import (
	"errors"
	"strconv"
)

type PairVal struct {
	Data interface{}
}

var pvVal error

func newPairVal(val interface{}) *PairVal {
	return &PairVal{val}
}

func (self *PairVal) ClearError() {
	pvVal = nil
}

func (self *PairVal) Error() error {
	return pvVal
}

func (self *PairVal) Type() string {
	self.ClearError()
	return "pairData"
}

func (self *PairVal) Set(key interface{}, tmp Config) bool {
	self.ClearError()
	return true
}

func (self *PairVal) Del(tmp interface{}) bool {
	self.ClearError()
	return true
}

func (self *PairVal) Get(tmp ...interface{}) Config {
	self.ClearError()
	return self
}

func (self *PairVal) ToString() (string, error) {
	self.ClearError()
	if tmp, ok := self.Data.(string); ok {
		pvVal = nil
		return tmp, nil
	} else {
		pvVal = errors.New("Data Is Not A String")
		return "", errors.New("Data Is Not A String")
	}
}

func (self *PairVal) ToInt() (int, error) {
	self.ClearError()
	if tmp, ok := self.Data.(int); ok {
		return tmp, nil
	} else {
		if str, ok := self.ToString(); nil == ok {
			ret, pvVal := strconv.Atoi(str)
			return ret, pvVal
		} else {
			pvVal = errors.New("Parse Data To Int Error!")
			return 0, errors.New("Parse Data To Int Error!")
		}
	}
}

func (self *PairVal) ToInt64() (int64, error) {
	self.ClearError()
	if tmp, ok := self.Data.(int64); ok {
		return tmp, nil
	} else {
		if str, err := self.ToString(); nil == err {
			ret, pvVal := strconv.ParseInt(str, 10, 64)
			return ret, pvVal
		} else {
			pvVal = errors.New("Parse Data To Int Error!")
			return 0, errors.New("Parse Data To Int Error!")
		}
	}
}

func (self *PairVal) ToFloat32() (float32, error) {
	self.ClearError()
	if tmp, ok := self.Data.(float32); ok {
		pvVal = nil
		return tmp, nil
	} else {
		if str, err := self.ToString(); nil == err {
			ret, pvVal := strconv.ParseFloat(str, 32)
			if nil == err {
				return (float32)(ret), nil
			} else {
				return 0.0, pvVal
			}
		} else {
			pvVal = errors.New("Parse Data To Float32 Error!")
			return 0.0, pvVal
		}
	}
}

func (self *PairVal) ToFloat64() (float64, error) {
	self.ClearError()
	if tmp, ok := self.Data.(float64); ok {
		pvVal = nil
		return tmp, nil
	} else {
		if str, err := self.ToString(); nil == err {
			ret, pvVal := strconv.ParseFloat(str, 64)
			return ret, pvVal
		} else {
			pvVal = errors.New("Parse Data To Float32 Error!")
			return 0.0, pvVal
		}
	}
}
