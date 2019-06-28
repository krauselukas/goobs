package goobs

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

type Client struct {
  Username string
  Password string
  BaseUrl string
}

func NewObsAuthClient(username, password, baseurl string) *Client {
  if baseurl == "" {
    baseurl = "https://api.opensuse.org"
  }
  return &Client{ Username: username, Password: password, BaseUrl: baseurl }
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
  req.SetBasicAuth(s.Username, s.Password)
  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  if res.StatusCode != 200 {
    return nil, fmt.Errorf("%s", body)
  }
  return body, nil
}

