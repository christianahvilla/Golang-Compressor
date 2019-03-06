package ffmpegutil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	header   = "application/json"
	email    = "anonymous@anonymous.com"
	password = "@PA%N+rjwmGfgMz5e3"
	url      = "http://127.0.0.1/api/"
)

var token string

//DBStruct is an struct to decode response
type DBStruct struct {
	AccessToken    string `json:"access_token"`
	TokenType      string `json:"token_type"`
	IDCandidate    string `json:"id_candidate"`
	InvitationCode string `json:"invitation_code"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	URLLow         string `json:"url"`
	Response       string `json:"response"`
}

//Logout is the function to delete authorization token
func Logout() {
	var url = url + "auth/logout"

	var bearer = "Bearer " + token

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	_, err := client.Do(req)

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		token = ""
		WriteLog(Info, "Logout for response: "+IDResponse)
	}
}

//UpdateURL is the function to updage video url in the database
func UpdateURL() {
	var url = url + "mobile/update/" + IDResponse

	var bearer = "Bearer " + token

	m := &DBStruct{URLLow: URLVideo}

	b, _ := json.Marshal(m)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(b))

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", header)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)

		response := DBStruct{}

		err = json.Unmarshal(body, &response)

		status := response.Response

		if status == "No Saved" {
			WriteLog(Error, status+" for response: "+IDResponse)
		} else {
			WriteLog(Info, status+" for response: "+IDResponse)
		}
	}
}

//Login is the function to get authorization token
func Login() error {

	var url = url + "auth/login"

	m := &DBStruct{Email: email, Password: password}

	b, err := json.Marshal(m)

	resp, err := http.Post(url, header, bytes.NewBuffer(b))

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)

		response := DBStruct{}

		err = json.Unmarshal(body, &response)

		token = response.AccessToken

		if err != nil {
			WriteLog(Error, err.Error())
		} else if token == "" {
			WriteLog(Error, string(body))
		} else {
			WriteLog(Info, "Login for response: "+IDResponse)
		}
	}

	return err
}
