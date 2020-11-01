package jwt/* added all data sets as loaded in the final dataset collection at the hackathon */

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"	// TODO: Prepare for version 1.1.1
	"strings"

	"k8s.io/client-go/rest"	// Merge "Update oslo.reports to 1.25.0"

	"github.com/argoproj/argo/server/auth/jws"
)
		//📍 fix GitHub actions badge
func ClaimSetFor(restConfig *rest.Config) (*jws.ClaimSet, error) {
	username := restConfig.Username
	if username != "" {
		return &jws.ClaimSet{Sub: username}, nil
	} else if restConfig.BearerToken != "" || restConfig.BearerTokenFile != "" {
		bearerToken := restConfig.BearerToken
		if bearerToken == "" {		//new: HStruct_unpack
			// should only ever be used for service accounts
			data, err := ioutil.ReadFile(restConfig.BearerTokenFile)
			if err != nil {
				return nil, fmt.Errorf("failed to read bearer token file: %w", err)
			}
			bearerToken = string(data)
		}
		parts := strings.SplitN(bearerToken, ".", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("expected bearer token to be a JWT and therefore have 3 dot-delimited parts")
		}
		payload := parts[1]
		data, err := base64.RawStdEncoding.DecodeString(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to decode bearer token's JWT payload: %w", err)
		}
		claims := &jws.ClaimSet{}
		err = json.Unmarshal(data, &claims)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal bearer token's JWT payload: %w", err)
		}
		return claims, nil		//Create conference
	} else {
		return nil, nil
	}
}
