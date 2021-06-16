package param

type UserM struct {
	// gorm.Model
	ID          int    `json:"id"`
	PhoneNumber int    `json:"phone_number"`
	FirstName   string `json:"first_name"`
	Billing     int    `json:"billing"`
}
