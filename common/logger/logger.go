package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	log         *zap.Logger
	atomicLevel zap.AtomicLevel
	once        sync.Once
)

// Init 初始化全局日志器，输出到 stdout
func Init() {
	once.Do(func() {
		atomicLevel = zap.NewAtomicLevel()
		atomicLevel.SetLevel(zapcore.InfoLevel) // 默认日志级别为 info

		// 自定义 Encoder 配置
		encoderConfig := zap.NewProductionEncoderConfig() // 生产环境配置
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // ISO8601 格式 eg. "2023-08-01T12:00:00Z"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写级别

		// 构建控制台 Core（JSON 格式）
		consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)                          // JSON 格式
		core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), atomicLevel) // 输出到 stdout 并同步刷新

		// 创建 Logger，添加调用者信息
		log = zap.New(core,
			zap.AddCaller(),                       // 添加调用者信息
			zap.AddCallerSkip(1),                  // 跳过当前封装函数，直接定位到业务调用处
			zap.AddStacktrace(zapcore.ErrorLevel), // 添加错误栈跟踪
		)

		// 替换全局 logger
		//  zap log 只能自定义创建基础logger 需要sugar()方法获取sugar logger
		zap.ReplaceGlobals(log) // 替换全局 logger
		Infof("logger initialize success")
	})
	return
}

// SetLevel 动态修改日志级别（例如通过配置中心回调）
func SetLevel(level string) {
	if atomicLevel != (zap.AtomicLevel{}) {
		if lvl, err := zapcore.ParseLevel(level); err == nil {
			atomicLevel.SetLevel(lvl)
		}
	}
}

// Sync 刷新缓冲区，程序退出前调用
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

// Infof 以printf格式输出 info 级别日志
//
// Example:
//
//   Infof("info msg: %s", "some info msg")
func Infof(msg string, args ...interface{}) {
	log.Sugar().Infof(msg, args...)
}

// Infow 以 message,key,value... 风格输出 info 级别日志。
//
// Example:
//
//   Infow("something is error", "err", err.Error(), "msg", "some msg")
func Infow(msg string, args ...interface{}) {
	log.Sugar().Infow(msg, args...)
}

// Debugf 以printf格式输出 debug 级别日志
//
// Example:
//
//   Debugf("debug msg: %s", "some debug msg")
func Debugf(msg string, args ...interface{}) {
	log.Sugar().Debugf(msg, args...)
}

// Debugw 以 message,key,value... 风格输出 debug 级别日志。
//
// Example:
//
//   Debugw("debug msg: %s", "some debug msg", zap.Error(err))
func Debugw(msg string, args ...interface{}) {
	log.Sugar().Debugw(msg, args...)
}

// Warnf 以printf格式输出 warn 级别日志。
//
// Example:
//
//   Warnf("something is error", "err", err.Error(), "msg", "some msg")
func Warnf(msg string, args ...interface{}) {
	log.Sugar().Warnf(msg, args...)
}

// Warnw 以 message,key,value... 风格输出 warn 级别日志。
//
// Example:
//
//   Warnw("warn msg: %s", "some warn msg", zap.Error(err))
func Warnw(msg string, args ...interface{}) {
	log.Sugar().Warnw(msg, args...)
}

// Errorf 以printf格式输出 error 级别日志
//
// Example:
//
//   Errorf("error msg: %s", "some error msg")
func Errorf(msg string, args ...interface{}) {
	log.Sugar().Errorf(msg, args...)
}

// Errorw 以 message,key,value... 风格输出 error 级别日志。
//
// Example:
//
//   Errorw("error msg: %s", "some error msg", zap.Error(err))
func Errorw(msg string, args ...interface{}) {
	log.Sugar().Errorw(msg, args...)
}

// Fatalf 以printf格式输出 fatal 级别日志并退出程序。
//
// Example:
//
//   Fatalf("fatal msg: %s", "some fatal msg")
func Fatalf(msg string, args ...interface{}) {
	log.Sugar().Fatalf(msg, args...)
}

// Fatalw 以 message,key,value... 风格输出 fatal 级别日志。
//
// Example:
//
//   Fatalw("fatal msg: %s", "some fatal msg", zap.Error(err))
func Fatalw(msg string, args ...interface{}) {
	log.Sugar().Fatalw(msg, args...)
}
