package dtos

type Response struct {
	Data interface{} `json:"data"`
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"full_name"`
	Token    string `json:"token"`
	Money    int64  `json:"money"`
	Phone    string `json:"phone"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type CreateLogRequest struct {
	Money  int64  `json:"money"`
	Tag    string `json:"tag"`
	Detail string `json:"detail"`
}
