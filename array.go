package cfg

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

type Array struct {
	size  int
	Data  []Config
	mutex *sync.Mutex
}

var arrError error

func NewArray() *Array {
	return &Array{0, []Config{}, &sync.Mutex{}}
}

func (self *Array) Size() int {
	return self.size
}

func (self *Array) ClearError() {
	arrError = nil
}

func (self *Array) Error() error {
	return arrError
}

func (self *Array) Set(key interface{}, val Config) bool {
	self.ClearError()
	ret := false
	self.mutex.Lock()
	if nil == key {
		self.Data = append(self.Data, val)
		self.size += 1
	} else {
		if index, ok := key.(int); ok {
			var tmp []Config
			if 0 >= self.size {
				tmp = append(tmp, val)
			} else {
				if index >= self.size {
					tmp = self.Data
					tmp = append(tmp, val)
				} else {
					for i := 0; i < self.size; i++ {
						if i < index {
							tmp = append(tmp, self.Data[i])
						} else if i == index {
							tmp = append(tmp, val)
							tmp = append(tmp, self.Data[i])
						} else {
							tmp = append(tmp, self.Data[i])
						}
					}
				}
			}
			self.size += 1
			self.Data = tmp
			ret = true
		} else {
			self.Data = append(self.Data, val)
			self.size += 1
			ret = true
		}
	}
	self.mutex.Unlock()
	return ret
}

func (self *Array) Get(tmp ...interface{}) Config {
	self.ClearError()
	if nil != tmp {
		if index, ok := tmp[0].(int); ok {
			if index >= self.size {
				arrError = errors.New("key Grant Than size")
				return nil
			} else {
				return self.Data[index]
			}
		} else {
			if key, ok := tmp[0].(string); ok {
				arr := strings.Split(key, ".")
				arrLen := len(arr)
				if 1 == arrLen {
					index, err := strconv.Atoi(arr[0])
					if nil != err {
						arrError = errors.New("Array Key(Len=1) Must be Int")
						return nil
					} else {
						if index >= self.size {
							arrError = errors.New("key Grant Than size")
							return nil
						} else {
							return self.Data[index]
						}
					}
				} else {
					var k interface{}
					var p Config = self
					for i := 0; i < arrLen; i++ {
						index, err := strconv.Atoi(arr[i])
						if nil == err {
							k = index
						} else {
							k = arr[i]
						}
						p = p.Get(k)
						if nil == p {
							arrError = errors.New("Parent Error in Array!")
							return nil
						}
					}
					return p
				}
			} else {
				arrError = errors.New("Array key Must be int")
				return nil
			}
		}
	} else {
		if self.size > 0 {
			return self.Data[0]
		} else {
			arrError = errors.New("Array Empty")
			return nil
		}
	}
}

func (self *Array) GetAll() []Config {
	self.ClearError()
	return self.Data
}

func (self *Array) Type() string {
	self.ClearError()
	return "array"
}

func (self *Array) Del(key interface{}) bool {
	self.ClearError()
	if index, ok := key.(int); ok {
		if index >= self.size {
			return false
		} else {
			var ret []Config
			self.mutex.Lock()
			for i := 0; i < self.size; i++ {
				if i != key {
					ret = append(ret, self.Data[i])
				}
			}
			self.Data = ret
			self.mutex.Unlock()
			return true
		}
	} else {
		return false
	}
}

func (self *Array) ToString() (string, error) {
	arrError = errors.New("Array Do Not Support Such Method")
	return "", arrError
}

func (self *Array) ToInt() (int, error) {
	arrError = errors.New("Array Do Not Support Such Method")
	return 0, arrError
}
func (self *Array) ToInt64() (int64, error) {
	arrError = errors.New("Array Do Not Support Such Method")
	return 0, arrError
}
func (self *Array) ToFloat32() (float32, error) {
	arrError = errors.New("Array Do Not Support Such Method")
	return 0.0, arrError
}
func (self *Array) ToFloat64() (float64, error) {
	arrError = errors.New("Array Do Not Support Such Method")
	return 0.0, arrError
}
