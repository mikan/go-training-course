package github

import "encoding/base64"

// Credential contains username and password.
type Credential struct {
	Username string
	Password string
}

// BasicAuth generates basic authentication attribute.
func BasicAuth(cred *Credential) (string, string) {
	auth := cred.Username + ":" + cred.Password
	return "Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
