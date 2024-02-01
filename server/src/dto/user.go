package dto

type SingUp struct {
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type Login struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
