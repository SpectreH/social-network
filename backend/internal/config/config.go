package config

import "time"

const (
	PORT                    = ":4000"
	DATABASE_PATH           = "./db/network.db"
	MIGRATIONS_SOURCE       = "file://internal/database/migrations/sqlite/"
	AVATAR_MAX_SIZE         = 5156000 // 5 MB
	SESSION_NAME            = "sn_token"
	SESSION_EXPIRATION_TIME = 1200 * time.Second
	AVATAR_SAVE_PATH        = "./images/"
	DEFAULT_AVATAR          = "default_avatar.png"
)
