package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaseto(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("SymmetricKey too short should be: %v", chacha20poly1305.KeySize)
	}

	paseto := &PasetoMaker{
		Paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return paseto, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", err
	}

	return maker.Paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.Paseto.Decrypt(token, maker.symmetricKey, payload, nil)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
