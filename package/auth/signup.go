package auth

import "encoding/json"

type SignUpResponce struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SignUp(username string, password string, email string) string {
	res := SignUpResponce{
		Message: "SignUp Successful",
		Status:  200,
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		return "Error in signup"
	}

	return string(jsonRes)
}
