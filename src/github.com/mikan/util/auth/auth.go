// Copyright 2016 mikan. All rights reserved.

// Utilities for http authentication operations.
package auth

import "encoding/base64"

type Credential struct {
	Username string
	Password string
}

func BasicAuth(cred *Credential) (string, string) {
	auth := cred.Username + ":" + cred.Password
	return "Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
