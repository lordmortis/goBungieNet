package goBungieNet

import (
  "encoding/json"
  "net/http"
  "errors"
  "io"
//  "fmt"
)

const serverBase string = "https://bungie.net"
const apiBase string = serverBase + "/Platform"

type Response struct {
  Response interface{}
  ErrorCode errorCode
  ThrottleSeconds int
  ErrorStatus string
  Message string
  MessageData map[string]string
}

type PrivacySetting int

const (
  PrivacyNone PrivacySetting = 0
  PrivacyPublic = 1
  PrivacyPrivate = 2
)

type DestinyVersion int

const (
  VersionNone DestinyVersion = 0
  VersionDestiny2 = 1
)

var (
  ApiKey string
)

func init() {
  ApiKey = ""
}

func get(endpoint string) (*Response, error) {
  var request *http.Request
  var err error
  var response *http.Response
  var parsedResponse *Response
  request, err = setupClient("GET", endpoint, nil)
  if err != nil { return nil, err }
  response, err = http.DefaultClient.Do(request)
  if (err != nil) { return nil, err }
  parsedResponse, err = getResponse(response.Body)
  if (err != nil) { return nil, err }

  return parsedResponse, nil
}

func setupClient(method string, endpoint string, body io.Reader) (*http.Request, error) {
  if ApiKey == "" { return nil, errors.New("No Api Key provided") }
  request, err := http.NewRequest(method, apiBase + endpoint, body)
  if (err == nil) {
    request.Header.Add("X-API-Key", ApiKey)
  }
  return request, err
}

func getResponse(reader io.Reader) (*Response, error) {
  decoder := json.NewDecoder(reader)
  var response Response
  err := decoder.Decode(&response)
  if (err != nil) {
    return nil, errors.New("Decoder Error: " + err.Error())
  }

  return &response, nil
}
