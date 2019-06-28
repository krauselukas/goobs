package goobs

import (
  "net/http"
  "fmt"
  "encoding/xml"
)

type About struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Revision string `xml:"revision"`
	Commit string `xml:"commit"`
}

func (s *Client) GetAbout() (*About, error) {
	url := fmt.Sprintf(s.BaseUrl + "/about")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	  return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
	  return nil, err
	}
	var data About
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
	  return nil, err
	}
	return &data, nil
}