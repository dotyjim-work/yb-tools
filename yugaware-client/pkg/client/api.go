package client

import (
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/yugabyte/yb-tools/yugaware-client/entity/yugaware"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/certificate_info"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/universe_management"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

func (c *YugawareClient) RegisterYugaware(request *yugaware.RegisterYugawareRequest) (*yugaware.RegisterYugawareResponse, error) {

	RegisterPath := "/api/v1/register"
	response := &yugaware.RegisterYugawareResponse{}

	_, err := c.newRequest().
		Post().
		Path(RegisterPath).
		RequestBody(request).
		DecodeResponseInto(response).
		Do()

	if err != nil {
		//if yr.Response.StatusCode == 401 {
		//	return nil,
		//}
		return nil, err
	}

	return response, nil
}

func (c *YugawareClient) Login(request *yugaware.LoginRequest) (*yugaware.LoginResponse, error) {

	LoginPath := "/api/v1/login"
	response := &yugaware.LoginResponse{}

	_, err := c.newRequest().
		Post().
		Path(LoginPath).
		RequestBody(request).
		DecodeResponseInto(response).
		Do()

	if err != nil {
		return nil, err
	}

	err = c.persistCookies()
	if err != nil {
		return nil, err
	}

	c.authToken = response.AuthToken
	c.customerUUID = strfmt.UUID(response.CustomerUUID)
	c.userUUID = strfmt.UUID(response.UserUUID)

	return response, nil
}

func (c *YugawareClient) Logout() error {
	c.expireSession()

	return c.cookiejar.Save()
}

func (c *YugawareClient) GetUniverseByIdentifier(identifier string) (*models.UniverseResp, error) {

	params := universe_management.NewListUniversesParams().
		WithCUUID(c.CustomerUUID())

	universes, err := c.PlatformAPIs.UniverseManagement.ListUniverses(params, c.SwaggerAuth)
	if err != nil {
		return nil, err
	}

	for _, universe := range universes.GetPayload() {
		if universe.Name == identifier {
			return universe, nil
		}

		if universe.UniverseUUID == strfmt.UUID(strings.ToLower(identifier)) {
			return universe, nil
		}
	}

	return nil, nil
}

func (c *YugawareClient) GetCertByIdentifier(identifier string) (*models.CertificateInfo, error) {

	params := certificate_info.NewGetListOfCertificateParams().WithCUUID(c.CustomerUUID())

	certs, err := c.PlatformAPIs.CertificateInfo.GetListOfCertificate(params, c.SwaggerAuth)
	if err != nil {
		return nil, err
	}

	for _, cert := range certs.GetPayload() {
		if cert.Label == identifier {
			return cert, nil
		}

		if cert.UUID == strfmt.UUID(strings.ToLower(identifier)) {
			return cert, nil
		}
	}

	return nil, nil
}
