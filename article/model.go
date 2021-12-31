package article

// article 모델
// belongs to: 종목
// 추후 추가 기능: weather
type Article struct {
	PostNo      uint   `gorm:"primaryKey" json:"post_no"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ReportedCnt uint   `json:"reported_cnt"`
	VoteUp      uint   `json:"vote_up"`
	VodeDown    uint   `json:"vote_down"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
