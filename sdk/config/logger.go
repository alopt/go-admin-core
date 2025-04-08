package config

import "github.com/go-admin-team/go-admin-core/sdk/pkg/logger"

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
	Format    string // 新增日志格式配置
	Output    string // 新增日志输出目标配置
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
		logger.WithFormat(e.Format), // 应用新的配置
		logger.WithOutput(e.Output), // 应用新的配置
	)
}

var LoggerConfig = new(Logger)
