package entity

type User struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`

	//password always keep hashed password
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at"`
}
