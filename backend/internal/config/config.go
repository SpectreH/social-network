package config

import "time"

const (
	PORT                         = ":4000"
	DATABASE_PATH                = "./db/network.db"
	MIGRATIONS_SOURCE            = "file://internal/database/migrations/sqlite/"
	AVATAR_MAX_SIZE              = 5156000 // 5 MB
	SESSION_NAME                 = "sn_token"
	SESSION_EXPIRATION_TIME      = 1200 * time.Second
	AVATAR_SAVE_PATH             = "./images/"
	DEFAULT_AVATAR               = "default_avatar.png"
	AVATAR_PATH_URL              = "http://localhost:4000/images/"
	FOLLOW_REQUEST_MESSAGE       = "Wants to be your follower"
	GROUP_FOLLOW_REQUEST_MESSAGE = "Wants to be a member of your group: "
	GROUP_FOLLOW_REQUEST_TYPE    = "groupFollowRequest"
	FOLLOW_REQUEST_TYPE          = "followRequest"
	GROUP_INVITE_TYPE            = "inviteRequest"
	GROUP_INVITE_MESSAGE         = "Invites you to join group: "
)
