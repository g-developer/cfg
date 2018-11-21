package cfg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	i := 0
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil {
			if io.EOF == err {
				if 0 < len(line) {
					err = Scan(i, line)
					if nil != err {
						return err
					}
					i++
				} else {
					break
				}
			} else {
				panic(err)
			}
		} else {
			err = Scan(i, line)
			if nil != err {
				return err
			}
			i++
		}
	}
	return nil
}

func initBlock() *Block {
	root := NewBlock("root", 0, KV_TYPE)
	root.Prefix = "root"
	root.PType = KV_TYPE
	global := NewBlock("global", 1, KV_TYPE)
	global.Prefix = "root.global"
	global.Pprefix = "root"
	global.PType = KV_TYPE
	global.SetLineStart(0)
	root.AddChild(global)
	global.SetParent(root)
	Push2Stack(root)
	Push2Stack(global)
	return root
}

func block2Cfg(root Config, index int, parentType int, rb *Block) {
	if nil == rb {
		return
	} else {
		var parent Config
		var tmp Config
		if rb.Level == 0 {
			tmp = NewKV()
			root.Set(rb.Name, tmp)
		}
		target := root.Get(rb.Prefix)
		if nil == target {
			target = NewKV()
			p := root.Get(rb.Pprefix)
			if nil != p {
				if ARRAY_TYPE == rb.Type {
					k := NewKV()
					p.Set(rb.Name, k)
				} else {
					p.Set(rb.Name, target)
				}
			}
		}
		target = root.Get(rb.Prefix)
		if nil == target {
			err := fmt.Errorf("target Nil! key=%v, target=%v, rb=%v, cfg=%v", rb.Prefix, ToString(target), ToString(rb), ToString(root))
			panic(err)
		}

		//scan rows
		rowLen := len(rb.Rows)
		for i := 0; i < rowLen; i++ {
			ty, key, val := ReconizeKeyType(rb.Rows[i], false)
			p := NewPair(key, val)
			if PAIR_TYPE == ty {
				target.Set(key, p)
			} else {
				exist := target.Get(key)
				if nil == exist {
					exist = NewArray()
					target.Set(key, exist)
				}
				exist.Set(key, p)
			}
		}
		//scan children
		childrenLen := len(rb.Children)
		for i := 0; i < childrenLen; i++ {
			key := rb.Children[i].Name
			ty := rb.Children[i].Type
			parent = root.Get(rb.Children[i].Pprefix)
			if nil == parent {
				if ARRAY_TYPE == ty {
					str := rb.Children[i].Pprefix
					pos1 := strings.LastIndex(str, ".")
					key1 := str[0:pos1]
					if ARRAY_TYPE == rb.Children[i].PType {
						if nil == root.Get(rb.Children[i].Prefix) {
							fx := root.Get(key1)
							ar := NewArray()
							fx.Set(key, ar)
						}
					} else {
						root.Get(key1).Set(key, NewArray())
					}
				}
			} else {
				if KV_TYPE == rb.Children[i].PType {
					if ARRAY_TYPE == ty {
						parent.Set(key, NewKV())
					}
				} else {
					if KV_TYPE == ty {
						if nil == root.Get(rb.Children[i].Prefix) {
							t := NewKV()
							t.Set(key, NewKV())
							parent.Set(nil, t)
						}
					} else {
						if nil == root.Get(rb.Children[i].Pprefix) {
							parent.Set(key, NewArray())
						} else {
							parent.Set(key, NewKV())
						}
					}
				}
			}
			block2Cfg(root, i, rb.Type, rb.Children[i])
		}
	}
}

func ToString(rs interface{}) string {
	b, err := json.Marshal(rs)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	} else {
		return string(b)
	}
}

func Load(filename string) Config {
	root := initBlock()
	err := load(filename)
	if nil != err {
		panic(err)
	}
	var cfg Config = NewKV()
	block2Cfg(cfg, 0, KV_TYPE, root)
	return cfg
}
