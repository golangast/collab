package user

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) CreateUser(e string, p string) {
	u.Email = e
	u.Password = p

}
