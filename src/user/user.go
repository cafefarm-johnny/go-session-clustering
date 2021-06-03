package user

type User struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Password string `json:"password,omitempty"`
}

var (
	UserJohnny = &User{
		Name: "Johnny",
		ID:   "jamong",
	}
)
