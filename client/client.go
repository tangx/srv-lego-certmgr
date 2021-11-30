package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Client struct {
	Endpoint string `yaml:"endpoint,omitempty" json:"endpoint,omitempty" env:"endpoint,omitempty"`
}

func (lego *Client) urlJoin(uri string) string {

	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(lego.Endpoint, "/"),
		strings.TrimLeft(uri, "/"),
	)
}

func (lego *Client) GetCertByName(provider, domain string) (*legox.Certificate, error) {
	uri := fmt.Sprintf("/lego-certmgr/gen/%s/%s", provider, domain)
	resp := &GetCertResponse{}

	err := lego.doHandle(http.MethodGet, uri, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("%s", resp.Error)
	}

	return &resp.Data, nil
}

func (lego *Client) CreateCert(provider, domain string) (bool, error) {
	uri := fmt.Sprintf("/lego-certmgr/gen/%s/%s", provider, domain)

	resp := &CreateCertReponse{}

	err := lego.doHandle(http.MethodPost, uri, resp)
	if err != nil {
		return false, err
	}

	if resp.Code == 201 {
		return true, nil
	}

	return false, nil
}

func (lego *Client) doHandle(method string, uri string, obj interface{}) error {
	url := lego.urlJoin(uri)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}
