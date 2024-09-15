package privatekey_registry

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
)

func mapPEMToRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	var err error

	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errs.ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errs.ErrNotRSAPrivateKey
	}

	return pkey, nil
}
