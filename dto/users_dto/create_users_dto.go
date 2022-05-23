package usersdto

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	No_HP    string `json:"no_hp"`
	Password string `json:"password"`
}
