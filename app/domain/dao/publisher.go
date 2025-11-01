package dao

type Publisher struct {
	BaseDao

	Name    string `gorm:"not null" json:"name"`
	Address string `gorm:"type:text;not null" json:"address"`
	Phone   string `gorm:"not null" json:"phone"`
	Website string `gorm:"not null" json:"website"`

	Book []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"books"`
}
