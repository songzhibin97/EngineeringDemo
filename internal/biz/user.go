package biz

// User: 定义模型
type User struct {
	ID       uint   `json:"id" gorm:"column:id;primaryKey,autoIncrement"`
	Age      uint   `json:"age" gorm:"column:age"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (u *User) TableName() string {
	return "user"
}

// IUser: 抽象接口
type IUser interface {
	Login(username, password string) bool
	Registered(u *User) bool
}

// UserCase
type UserCase struct {
	U IUser
}

// 贫血方法
func (u *UserCase) QLogin(username, password string) bool {
	return u.U.Login(username, password)
}
func (u *UserCase) QReg(us *User) bool {
	return u.U.Registered(us)
}

// 实例化
func NewUserCase(l IUser) *UserCase {
	return &UserCase{U: l}
}
