package model

type User struct {
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	County   string `json:"county"`
	Salary   string `json:"salary"`
}
