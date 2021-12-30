package article

// article 모델
// belongs to: 종목
// 추후 추가 기능: weather
type Article struct {
	PostNo      uint   `gorm:"primaryKey"`
	CodeNo      string //연동 게시판 추후 foreign key 적용 예정
	Title       string
	Content     string
	ReportedCnt uint
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
}
