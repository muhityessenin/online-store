package user

type Request struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Roles   string `json:"roles"`
}
