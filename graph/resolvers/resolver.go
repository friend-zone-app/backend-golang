package graph

import "encoding/base64"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *Resolver) MarshalBytes(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (r *Resolver) UnmarshalBytes(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}
