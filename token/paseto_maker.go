package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symmericKey []byte
}

func NewPasetoMaker(symmericKey string) (Maker, error) {
	if len(symmericKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symmericKey: []byte(symmericKey),
	}
	return maker, nil
}

// CreateToken create a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayLoad(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmericKey, payload, nil)
}

// VerifyToken checks if the token is vbalid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmericKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
