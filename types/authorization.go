package types

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	//Token holds the returned token code
	Token string `json:"token"`
}
