package handlers

type registerRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
}
