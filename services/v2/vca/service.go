// service.go - define the client for VCA service

// Package vca defines the VCA services of BCE. The supported APIs are all defined in sub-package
package vca

import (
	"github.com/basebytes/bce-sdk-go/auth"
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/client"
)

const (
	DefaultServiceDomain = "vca." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Service of VCR service is a kind of BceClient, so derived from BceClient
type Service struct {
	*bce.BceClient
	*client.MediaClient
	*client.StreamClient
	*client.FaceClient
	*client.LogoClient
}

// NewService make the VCR service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewService(ak, sk, endpoint string) (*Service, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DefaultServiceDomain
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	bceClient := bce.NewBceClient(defaultConf, v1Signer)
	service := &Service{
		BceClient:    bceClient,
		MediaClient:  client.NewMediaClient(bceClient),
		StreamClient: client.NewStreamClient(bceClient),
		FaceClient:   client.NewFaceClient(bceClient),
		LogoClient:   client.NewLogoClient(bceClient),
	}
	return service, nil
}