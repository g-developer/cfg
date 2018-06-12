package cfg

type Config interface {
	Set(key interface{}, val Config) bool
	Get(key ...interface{}) Config
	Del(key interface{}) bool
	Type() string
	Error() error
	ClearError()
	ToString() (string, error)
	ToInt() (int, error)
	ToInt64() (int64, error)
	ToFloat32() (float32, error)
	ToFloat64() (float64, error)
}

const (
	CONFIG_LEX_NODE_START     = "["
	CONFIG_LEX_NODE_END       = "]"
	CONFIG_LEX_NODE_LEVEL     = '.'
	CONFIG_LEX_ARRAY_PREFIX   = '@'
	CONFIG_LEX_COMMENT_PREFIX = '#'
	CONFIG_LEX_LINE_SEP       = "\n"
	CONFIG_LEX_KEY_SEP        = "."
)
