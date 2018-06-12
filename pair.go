package cfg

import (
	"errors"
)

type Pair struct {
	Key interface{}
	Val *PairVal
}

var pError error

func NewPair(key interface{}, val interface{}) *Pair {
	return &Pair{key, newPairVal(val)}
}

func (self *Pair) ClearError() {
	pError = nil
}

func (self *Pair) Error() error {
	return pError
}

func (self *Pair) Type() string {
	self.ClearError()
	return "pair"
}
func (self *Pair) GetKey() (interface{}, error) {
	self.ClearError()
	if nil != self {
		pError = nil
		return self.Key, nil
	} else {
		pError = errors.New("Pair Self Null")
		return "", pError
	}
}

func (self *Pair) Set(key interface{}, tmp Config) bool {
	self.ClearError()
	return true
}

func (self *Pair) Del(tmp interface{}) bool {
	self.ClearError()
	return true
}

func (self *Pair) Get(tmp ...interface{}) Config {
	self.ClearError()
	return self.Val
}

func (self *Pair) ToString() (string, error) {
	return self.Val.ToString()
}

func (self *Pair) ToInt() (int, error) {
	return self.Val.ToInt()
}
func (self *Pair) ToInt64() (int64, error) {
	return self.Val.ToInt64()
}
func (self *Pair) ToFloat32() (float32, error) {
	return self.Val.ToFloat32()
}
func (self *Pair) ToFloat64() (float64, error) {
	return self.Val.ToFloat64()
}
