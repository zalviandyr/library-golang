package dao

type User struct {
	BaseDao

	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
