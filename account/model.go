package account

type User struct {
	UserId     uint   `gorm:"primaryKey" json:"user_id"`
	Email      string `gorm:"unique;not null" json:"email"`
	Nickname   string `gorm:"unique;not null" json:"nickname"`
	IsVerified bool   `json:"is_verified"`
	IsGuru     bool   `json:"is_guru"`
	IsBlocked  bool   `json:"is_blocked"`
	Starred    uint   `json:"starred"`
	Starring   uint   `json:"starring"`
	JoinedAt   int64  `gorm:"autoCreateTime" json:"joined_at"`
	LastLogin  int64  `gorm:"autoUpdateTime" json:"last_login"`
}

type UserAuth struct {
	User     User `gorm:"primaryKey;foreignkey:UserId;constraint:OnDelete:CASCADE;refereces:UserId"`
	UserId   uint
	Password string `gorm:"not null" json:"password"`
}

// 로그인 폼을 따로 받는 이우는....
// 1. DB에는 암호화된 PW를 집어 넣기 위함
// 2. 이메일 인증 절차를 거치기 위해서 (추후 구현)
// 싱글톤패턴을 쓴 이유는 별건 없고 디폴트 값 설정을 위해서...
// 추후 이메일 검증 로직을 추가할 예정
type userSignUpModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Repeat   string `json:"repeat"`
	// IsEmailVerified bool   `json:"is_email_verified"` // tmp
}

func UserSignUpModel() *userSignUpModel {
	u := &userSignUpModel{}
	// u.IsEmailVerified = true
	return u
}

type userLogInModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
