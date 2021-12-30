package article

// article 모델
type Article struct {
	PostNo      uint   `gorm:"primaryKey"`
	CodeNo      string //연동 게시판
	Title       string
	Content     string
	ReportedCnt uint
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
}
