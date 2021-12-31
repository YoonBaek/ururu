package account

type User struct {
	UserId     uint   `gorm:"primaryKey" json:"user_id"`
	Email      string `gorm:unique; not null"`
	password   string
	IsVerified bool  `json:"is_verified"`
	IsGuru     bool  `json:"is_guru"`
	IsBlocked  bool  `json:"is_blocked"`
	Starred    uint  `json:"starred"`
	Starring   uint  `json:"starring"`
	JoinedAt   int64 `gorm:"autoCreateTime" json:"joined_at"`
	LastLogin  int64 `gorm:"autoUpdateTime" json:"last_login"`
}
