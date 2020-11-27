package client

import (
	"context"
	"github.com/dgoodwin/syncsets/api"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/types"
	"net/http"
	"net/url"
	"path"
)

type APIClient interface {
	Get(ctx context.Context, key types.NamespacedName, obj api.APIResource) error
}

var _ APIClient = &APIClientImpl{}

func NewAPIClient(endpoint string) APIClient {
	return &APIClientImpl{
		endpoint: endpoint,
	}
}

type APIClientImpl struct {
	endpoint string
}

func (c *APIClientImpl) Get(ctx context.Context, key types.NamespacedName, obj api.APIResource) error {
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, obj.APIVersion())
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	obj.Scan(bodyBytes)
	return nil
}
