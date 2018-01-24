package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type User struct{
    Name string `json:"username"`
    Password string `json:"password"`
    Rolecode string `json:"rolecode"`
    PublicKey string `json:"publicKey"`
    PrivateKey string `json:"privateKey"`
}


// UnMarshalUserFromReader unmarshal
func UnMarshalUserFromReader(r io.Reader) (*User, error) {
	v := &User{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(v)

	if nil != err {
		return nil, err
	}
	return v, nil
}

// ToString return a desc string
func (b *User) ToString() string {
	return fmt.Sprintf("%s-%s", b.Name, b.Password)
}

// Marshal marshal
func (b *User) Marshal() []byte {
	v, _ := json.Marshal(b)
	return v
}

// UnMarshalUser unmarshal
func UnMarshalUser(data []byte) *User {
	v := &User{}
	err := json.Unmarshal(data, v)

	if err != nil {
		return v
	}

	return v
}
