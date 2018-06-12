package cfg

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

type KV struct {
	size  int
	Data  *map[string]Config
	mutex *sync.Mutex
}

var kvError error

func NewKV() *KV {
	return &KV{0, &map[string]Config{}, &sync.Mutex{}}
}

func (self *KV) ClearError() {
	kvError = nil
}

func (self *KV) Error() error {
	return kvError
}

func (self *KV) Set(tmp interface{}, val Config) bool {
	self.ClearError()
	var ret bool = false
	if key, ok := tmp.(string); ok {
		self.mutex.Lock()
		if _, ok := (*self.Data)[key]; ok {
			self.size += 1
		}
		(*self.Data)[key] = val
		ret = true
		self.mutex.Unlock()
	}
	return ret
}

func (self *KV) Type() string {
	self.ClearError()
	return "kv"
}

func (self *KV) Size() int {
	self.ClearError()
	return self.size
}

func (self *KV) Get(key ...interface{}) Config {
	self.ClearError()
	if nil != key {
		if tmp, ok := key[0].(string); ok {
			arr := strings.Split(tmp, ".")
			arrLen := len(arr)
			if 1 == arrLen {
				if v, ok := (*self.Data)[arr[0]]; ok {
					return v
				} else {
					kvError = errors.New("No Such Key")
					return nil
				}
			} else {
				var p Config = self
				var k interface{}
				for i := 0; i < arrLen; i++ {
					key, err := strconv.Atoi(arr[i])
					if nil != err {
						k = arr[i]
					} else {
						k = key
					}
					p = p.Get(k)
					if nil == p {
						p = nil
						kvError = errors.New("Parent Not Exist!")
						break
					}
				}
				return p
			}
		} else {
			kvError = errors.New("KV Key Must be String")
			return nil
		}
	} else {
		kvError = errors.New("KV Get Must Has Key")
		return nil
	}
}

func (self *KV) Get2(key ...interface{}) Config {
	self.ClearError()
	if nil != key {
		tmp := key[0]
		if key, ok := tmp.(string); !ok {
			kvError = errors.New("KV Key Must Be String")
			return nil
		} else {
			var ret Config
			var err error
			if v, ok := (*self.Data)[key]; ok {
				ret = v
				err = nil
			} else {
				ret = nil
				err = errors.New("No Suck Key[" + key + "]")
			}
			kvError = err
			return ret
		}
	} else {
		kvError = errors.New("KV Get Must Has Key")
		return nil
	}
}

func (self *KV) Del(tmp interface{}) bool {
	self.ClearError()
	if key, ok := tmp.(string); !ok {
		return false
	} else {
		self.mutex.Lock()
		if _, ok := (*self.Data)[key]; ok {
			self.size -= 1
			delete(*(self.Data), key)
		}
		self.mutex.Unlock()
		return true
	}
}

func (self *KV) ToString() (string, error) {
	kvError = errors.New("KV Do Not Support Such Method")
	return "", kvError
}

func (self *KV) ToInt() (int, error) {
	kvError = errors.New("KV Do Not Support Such Method")
	return 0, kvError
}
func (self *KV) ToInt64() (int64, error) {
	kvError = errors.New("KV Do Not Support Such Method")
	return 0, kvError
}
func (self *KV) ToFloat32() (float32, error) {
	kvError = errors.New("KV Do Not Support Such Method")
	return 0.0, kvError
}
func (self *KV) ToFloat64() (float64, error) {
	kvError = errors.New("KV Do Not Support Such Method")
	return 0.0, kvError
}
