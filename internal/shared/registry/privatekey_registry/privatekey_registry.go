package privatekey_registry

import (
	"crypto/rsa"
	"encoding/base64"

	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
)

type RegistryI interface {
	Get() *rsa.PrivateKey
}

type registry struct {
	pk *rsa.PrivateKey
}

func New(conf config.Config) RegistryI {
	r := registry{}
	r.process(conf)
	return &r
}

func (r *registry) Get() *rsa.PrivateKey {
	return r.pk
}

func (r *registry) process(conf config.Config) {
	pkString, err := base64.StdEncoding.DecodeString(conf.JWT.PrivateKey)
	if err != nil {
		panic(err)
	}

	pk, err := mapPEMToRSAPrivateKey(pkString)
	if err != nil {
		panic(err)
	}

	r.pk = pk
}
