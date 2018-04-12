package ulb

import (
	"net/http"

	"github.com/chanyipiaomiao/ucloud-sdk-go/ucloud"
	"github.com/chanyipiaomiao/ucloud-sdk-go/ucloud/service"
)

type ULB struct {
	*service.Service
}

func New(config *ucloud.Config) *ULB {

	service := &service.Service{
		Config:      ucloud.DefaultConfig.Merge(config),
		ServiceName: "ULB",
		APIVersion:  ucloud.APIVersion,

		BaseUrl:    ucloud.APIBaseURL,
		HttpClient: &http.Client{},
	}

	return &ULB{service}

}
