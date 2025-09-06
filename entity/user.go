package entity

type User struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`

	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at"`
	Role      Role
}
