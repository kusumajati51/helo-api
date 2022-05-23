package authdto


type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
