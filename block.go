package cfg

import (
	util "cfg/util"
	"errors"
	"strconv"
	"strings"
)

const (
	PAIR_TYPE  = 0
	KV_TYPE    = 1
	ARRAY_TYPE = 2
)

type Block struct {
	Name     string
	Level    int
	Start    int
	End      int
	Type     int
	Index    int
	Prefix   string
	Pprefix  string
	PType    int
	Rows     []string
	Children []*Block
	parent   *Block
}

var stack *util.Stack = util.NewStack()

func IsNewBlock(line string) (bool, int, string) {
	line = strings.TrimSpace(line)
	isBlock := false
	level := 1
	key := []byte{}
	last := len(line) - 1
	if 0 == strings.Index(line, CONFIG_LEX_NODE_START) && last == strings.LastIndex(line, CONFIG_LEX_NODE_END) {
		isBlock = true
		pureLine := util.TrimMultiBlank(line)
		lens := len(pureLine) - 1
		levelEnd := false
		for i := 1; i < lens; i++ {
			if CONFIG_LEX_NODE_LEVEL == pureLine[i] && !levelEnd {
				level++
			} else {
				levelEnd = true
				key = append(key, pureLine[i])
			}
		}
	}
	return isBlock, level, string(key)
}

func NewBlock(name string, level int, Type int) *Block {
	return &Block{name, level, -1, -1, Type, -1, "", "", 0, []string{}, []*Block{}, nil}
}

func (self *Block) SetLineStart(lineNo int) {
	self.Start = lineNo
}

func (self *Block) SetLineEnd(lineNo int) {
	self.End = lineNo
}

func (self *Block) AddToRow(line string) {
	self.Rows = append(self.Rows, line)
}

func (self *Block) AddChild(child *Block) {
	if nil != child {
		self.Children = append(self.Children, child)
	}
}

func (self *Block) SetParent(p *Block) {
	self.parent = p
}

func (self *Block) GetParent() *Block {
	return self.parent
}

func Push2Stack(b *Block) {
	if nil != b {
		stack.Push(b)
	}
}

func PopFromStack() *Block {
	ret := stack.Pop()
	if b, ok := ret.(*Block); ok {
		return b
	} else {
		return nil
	}
}

func GetStackTop() *Block {
	ret := stack.Top()
	if b, ok := ret.(*Block); ok {
		return b
	} else {
		return nil
	}
}

func ReconizeKeyType(str string, isBlock bool) (int, string, string) {
	strLen := len(str)
	if 0 >= strLen {
		return -1, "", ""
	}
	if isBlock {
		if CONFIG_LEX_ARRAY_PREFIX == str[0] {
			return ARRAY_TYPE, str[1:strLen], ""
		} else {
			return KV_TYPE, str, ""
		}
	} else {
		pos := strings.Index(str, ":")
		var key string
		var Type int
		if CONFIG_LEX_ARRAY_PREFIX == str[0] {
			key = str[1:pos]
			Type = ARRAY_TYPE
		} else {
			key = str[0:pos]
			Type = PAIR_TYPE
		}
		val := str[(pos + 1):strLen]
		return Type, strings.TrimSpace(key), strings.TrimSpace(val)
	}
}

func Scan(lineNo int, str string) error {
	str = strings.TrimSpace(str)
	if 0 >= len(str) {
		return errors.New("Len(str) Must > 0")
	}
	isBlock, level, name := IsNewBlock(str)
	if 0 >= level {
		return errors.New("Block.IsNewBlock Level Error!")
	}
	top := GetStackTop()
	if nil == top {
		return errors.New("Block Stack Empty! Need New Root!")
	}
	ty, key, _ := ReconizeKeyType(name, isBlock)
	var p *Block
	if isBlock {
		tmp := NewBlock(key, level, ty)
		if ARRAY_TYPE == ty {
			tmp.Index = 0
		}
		tmp.SetLineStart(lineNo)
		top.SetLineEnd(lineNo - 1)
		for {
			top = GetStackTop()
			if nil == top {
				return errors.New("Block Stack Empty! Need New Root!")
			} else {
				if top.Level > level {
					PopFromStack()
				} else {
					break
				}
			}
		}
		p = top.GetParent()
		if level == top.Level {
			if nil == p {
				top.AddChild(tmp)
				if ARRAY_TYPE == ty {
					tmp.Prefix = top.Prefix + "." + key + ".0"
				} else {
					tmp.Prefix = top.Prefix + "." + key
				}
				tmp.Pprefix = top.Prefix
				tmp.PType = top.Type
				tmp.SetParent(top)
			} else {
				pChildSize := len(p.Children)
				if ARRAY_TYPE == ty {
					index := 0
					for i := 0; i < pChildSize; i++ {
						if key == p.Children[i].Name {
							index++
						}
					}
					tmp.Index = index
					tmp.Prefix = p.Prefix + "." + key + "." + strconv.Itoa(index)
					tmp.Pprefix = p.Prefix + "." + key
					tmp.PType = p.Type
				} else {
					tmp.Prefix = p.Prefix + "." + key
					tmp.Pprefix = p.Prefix
					tmp.PType = p.Type
				}
				p.AddChild(tmp)
				tmp.SetParent(p)
			}
		} else if level == (top.Level + 1) {
			top.AddChild(tmp)
			if ARRAY_TYPE == ty {
				tmp.Prefix = top.Prefix + "." + key + ".0"
				tmp.Pprefix = top.Prefix + "." + key
			} else {
				tmp.Prefix = top.Prefix + "." + key
				tmp.Pprefix = top.Prefix
			}
			tmp.PType = top.Type
			tmp.SetParent(top)
		} else {
			return errors.New("Config File Level Error!")
		}
		Push2Stack(tmp)
	} else {
		top.AddToRow(str)
	}
	return nil
}
