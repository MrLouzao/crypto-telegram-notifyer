package controllers

import "net/http"

// Generate a 403 error from Response
func _send_403_error_response(writer http.ResponseWriter, errMessage string) {
	errJson := `{"code": 403 ,"message": "` + errMessage + `"}`
	writer.WriteHeader(403)
	writer.Write([]byte(errJson))
}

func _send_500_error_response(writer http.ResponseWriter, errMessage string) {
	errJson := `{"code": 500 ,"message": "` + errMessage + `"}`
	writer.WriteHeader(500)
	writer.Write([]byte(errJson))
}
