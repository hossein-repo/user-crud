
// User مدل یوزر
type User struct {
	BaseModel
	Username     string      `gorm:"type:varchar(20);not null;unique"`
	FirstName    string      `gorm:"type:varchar(15);default:null"`
	LastName     string      `gorm:"type:varchar(25);default:null"`
	MobileNumber string      `gorm:"type:varchar(11);unique;default:null"`
	Email        string      `gorm:"type:varchar(64);unique;default:null"`
	Password     string      `gorm:"type:varchar(64);not null"`
	Enabled      bool        `gorm:"default:true"`
	UserRoles    *[]UserRole
}