package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Messages string `json:"messages"`
}

type OutputDataUser struct {
	Username string `json:"username"`
	Name     string `jon:"name"`
	Email    string `json:"email"`
}

type OutputDataUserLogin struct {
	Username string `json:"username"`
	Name     string `jon:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type OutputProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// reformat API Response
func APIResponse(status string, code int, messages string, data interface{}) Response {
	meta := Meta{
		Status:   status,
		Code:     code,
		Messages: messages,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
