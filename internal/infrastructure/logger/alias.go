package logger

import (
	"go.uber.org/zap"
)

// Sync это функция, которая синхронизирует регистратор.
// Он возвращает ошибку, если синхронизация не удается.
func Sync() error {
	_ = zap.L().Sync()
	return nil
}

// Info это функция, которая регистрирует сообщение об информационном уровне.
// Он берет сообщение и вариативный параметр полей.
func Info(message string, fields ...zap.Field) {
	zap.L().Info(message, fields...)
}

// Debug это функция, которая регистрирует сообщение об уровне отладки.
// Он берет сообщение и вариативный параметр полей.
func Debug(message string, fields ...zap.Field) {
	zap.L().Debug(message, fields...)
}

// Error это функция, которая регистрирует сообщение уровня ошибки.
// Он берет сообщение и вариативный параметр полей.
func Error(message string, fields ...zap.Field) {
	zap.L().Error(message, fields...)
}

// Fatal это функция, которая регистрирует сообщение о роковом уровне.
// Он берет сообщение и вариативный параметр полей.
// После регистрации сообщения он вызывает OS.Exit (1), чтобы завершить программу.
func Fatal(message string, fields ...zap.Field) {
	zap.L().Fatal(message, fields...)
}

// Warn это функция, которая регистрирует сообщение уровня предупреждения.
// Он берет сообщение и вариативный параметр полей.
func Warn(message string, fields ...zap.Field) {
	zap.L().Warn(message, fields...)
}
