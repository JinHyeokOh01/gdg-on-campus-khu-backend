package auth

import(
	- "embed"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct{
	PrivateKey, PublicKey	jwk.Key
	Store					Store
	Clocker					clock.Clocker
}

type Store interface{
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store, c clock.Clocker) (*JWTer, error){
	j := &JWTer{Store: s}
	privKey, err := parse(rawPrivKey)
	if err != nil{
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil{
		return nil, fmt.Errof("failed in NewJWter: publick key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c
	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error){
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil{
		return nil, err
	}
	return key, nil
}