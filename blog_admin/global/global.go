package global

import (
	"admin/config"
	"admin/utils/timer"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.ProjectConfig
	DB     *gorm.DB
	LOG    *zap.Logger
	ENV    string
	Timer  timer.Timer = timer.NewTimerTask()
)
