package logger

import (
	"context"
	"io"
	"os"
)

// Option is a function that configures the logger options
// Option 是一个配置日志记录器选项的函数
type Option func(*Options)

// Options are logger options
// Options 是日志记录器选项
type Options struct {
	// Level is the logging level the logger should log at. default is `InfoLevel`
	// Level 是日志记录器的日志级别，默认是 `InfoLevel`
	Level Level
	// Fields are fields to always be logged
	// Fields 是始终要记录的字段
	Fields map[string]interface{}
	// Out is the output writer for the logger. default is `os.Stderr`
	// Out 是日志记录器的输出写入器，默认是 `os.Stderr`
	Out io.Writer
	// CallerSkipCount is the frame count to skip for file:line info
	// CallerSkipCount 是跳过的帧数，用于文件:行信息
	CallerSkipCount int
	// Context is alternative options
	// Context 是替代选项
	Context context.Context
	// Name is the logger name
	// Name 是日志记录器的名称
	Name string
}

// WithFields sets default fields for the logger
// WithFields 设置日志记录器的默认字段
func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

// WithLevel sets default level for the logger
// WithLevel 设置日志记录器的默认级别
func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithOutput sets default output writer for the logger
// WithOutput 设置日志记录器的默认输出写入器
func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithCallerSkipCount sets frame count to skip
// WithCallerSkipCount 设置要跳过的帧数
func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

// WithName sets name for logger
// WithName 设置日志记录器的名称
func WithName(name string) Option {
	return func(args *Options) {
		args.Name = name
	}
}

// SetOption sets a custom option
// SetOption 设置自定义选项
func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// DefaultOptions returns default options
// DefaultOptions 返回默认选项
func DefaultOptions() Options {
	return Options{
		Level:           InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 3,
		Context:         context.Background(),
		Name:            "",
	}
}
