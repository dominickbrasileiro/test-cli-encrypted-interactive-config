package encrypter

type Encrypter interface {
	Encrypt(secret, raw string) (string, error)
	Decrypt(secret, encrypted string) (string, error)
}
