// Package pkg содержит утилитарные функции и переменные, используемые в приложении.
// В данном файле определяется и настраивается глобальный логгер с использованием библиотеки logrus.
package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger — глобальный логгер, используемый для регистрации событий приложения.
var Logger = logrus.New()

// InitLogger инициализирует глобальный логгер.
// Функция пытается открыть файл "app.log" для записи логов. Если файл успешно открыт,
// логгирование направляется в файл, в противном случае — в стандартный вывод (stdout).
// Логгер настраивается на использование JSON-формата и уровень логирования Info.
func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.SetOutput(os.Stdout)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}
