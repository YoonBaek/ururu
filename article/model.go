package article

// article 모델
// belongs to: 종목
// 추후 추가 기능: weather
type Article struct {
	PostNo      uint   `gorm:"primaryKey" json:"post_no"`
	CodeNo      string `json:"code_no"` //연동 게시판 추후 foreign key 적용 예정
	Title       string `json:"title"`
	Content     string `json:"content"`
	ReportedCnt uint   `json:"reported_cnt"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
