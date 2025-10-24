package auth

import "encoding/json"

type SignInResponce struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SignIn(username string, password string) string {
	res := SignInResponce{
		Message: "SignIn Successful",
		Status:  200,
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		return "Error in signin"
	}

	return string(jsonRes)
}
