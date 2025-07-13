package logging

import (
    "strings"

    "github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
    logger := logrus.New()

    lvl, err := logrus.ParseLevel(strings.ToLower(level))
    if err != nil {
        lvl = logrus.InfoLevel
    }
    logger.SetLevel(lvl)
    logger.SetFormatter(&logrus.JSONFormatter{})

    return logger
}

