package dao

type Author struct {
	BaseDao
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Biography string `gorm:"type:text;not null" json:"biography"`

	Book []Book `gorm:"many2many:author_books;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"books"`
}
