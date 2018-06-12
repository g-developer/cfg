package main

import (
	cfg "cfg"
	"fmt"
)

func main() {
	conf := cfg.Load("../data/2.cfg")
	fmt.Println(cfg.ToString(conf))
	f := conf.Get("root.global")
	fmt.Println(cfg.ToString(f))
	if nil != f {
		ff, err := f.ToString()
		fmt.Println("zyc--row--", ff, err)
	}
}
