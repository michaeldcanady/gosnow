package gosnow

import "errors"

type OauthClient struct {
	clientId     string
	clientSecret string
	tokenUpdater string
	tokenUrl     string
}

func NewOauth(clientId, clientSecret, tokenUpdater string) (O OauthClient, err error) {

	if clientId == "" {
		err = errors.New("Testing")
		return O, err
	}

	O.clientId = clientId
	O.clientSecret = clientSecret
	O.tokenUpdater = tokenUpdater
	return
}
