package dto

type LoginDto struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type RegisterDto struct {
	Username        string `binding:"required" json:"username"`
	Password        string `binding:"required" json:"password"`
	ConfirmPassword string `binding:"required" json:"confirm_password"`
}
