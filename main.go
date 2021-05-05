package nvosrestgo

import (
	"bytes"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// System defines and tracks the system connection information
type System struct {
	host       string
	User       string `json:"username"`
	Password   string `json:"password"`
	port       int
	token      string
	HTTPClient *http.Client `json:"-"`
}

//new initializes the system structure that is used to interact with this lib
//   it also attemps to connect and returns the system structure for client use
func New(ip string, username string, password string, port int) *System {
	//create the new system
	sys := &System{
		host:       ip,
		User:       username,
		Password:   password,
		port:       port,
		HTTPClient: http.DefaultClient,
	}

	//need to ignore invalid certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	sys.connect()

	return sys
}

//connect to the NVOS based system, authenticate and store the token
func (s *System) connect() error {

	req, err := http.NewRequest("GET", "https://"+s.host+":"+strconv.Itoa(s.port)+"/api/auth", nil)
	if err != nil {
		return err
	}

	//add the headers
	auth_b64 := b64.StdEncoding.EncodeToString([]byte(s.User + ":" + s.Password))
	req.Header.Set("Authorization", "Basic "+string(auth_b64))
	req.Header.Set("Content-Type", "application/json")

	//run the request
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("error on auth")
		fmt.Println(string(auth_b64))
		fmt.Println(req.URL)
		return err
	}
	defer res.Body.Close()

	//if all worked out store the response
	s.token = res.Header.Get("x-auth-token")

	return nil
}

func (s *System) sendRequest(httpMethod string, url string, args ...string) (string, error) {

	req, err := http.NewRequest(httpMethod, "https://"+s.host+":"+strconv.Itoa(s.port)+url, nil)
	if err != nil {
		return "", err
	}

	//add the headers
	req.Header.Set("Authentication", s.token)
	req.Header.Set("Content-Type", "application/json")

	//run the request
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("error on " + httpMethod)
		fmt.Println(req.URL)
		return "", err
	}

	resp, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(resp), nil

}

//GetSystemProperty requests the provided named property and returns the api response
func (s *System) GetSystemProperty(property string) (string, error) {

	return s.sendRequest("GET", "/api/system?properties="+property)
}

//ExportConfig returns the settings file from an NTO system to a local file
func (s *System) ExportConfig(description string, export_type string, file_name string) error {
	exportReqSpec, _ := json.Marshal(map[string]string{"description": description, "export_type": export_type, "file_name": file_name})

	req, err := http.NewRequest("POST", "https://"+s.host+":"+strconv.Itoa(s.port)+"/api/actions/export", bytes.NewBuffer(exportReqSpec))
	if err != nil {
		return err
	}

	//add the headers
	req.Header.Set("Authentication", s.token)
	req.Header.Set("Content-Type", "application/json")

	//run the request
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("error on export")
		fmt.Println(req.URL)
		return err
	}

	defer res.Body.Close()
	out, err := os.Create(file_name)
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(out, res.Body)

	return nil
}
