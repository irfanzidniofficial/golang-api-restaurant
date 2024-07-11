package user

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

func (ur *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, ur.time, ur.memory, ur.threads, ur.keyLen)

	b64Hash := ur.encypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, ur.memory, ur.time, ur.threads, b64Salt, b64Hash)
	return encodedHash, nil
}

func (ur *userRepo) encypt(text []byte) string {
	nonce := make([]byte, ur.gcm.NonceSize())
	ciphertext := ur.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func (ur *userRepo) decrypt(ciphertext string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	if len(decoded) < ur.gcm.NonceSize() {
		return nil, errors.New("invalid nonce sizes")
	}

	return ur.gcm.Open(nil,
		decoded[:ur.gcm.NonceSize()],
		decoded[ur.gcm.NonceSize():], nil)
}

func (ur *userRepo) comparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	var memory, time uint32
	var parallerism uint32

	switch parts[1] {
	case "argon2id":
		_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &parallerism)
		if err != nil {
			return false, err
		}
		salt, err := base64.RawStdEncoding.DecodeString(parts[4])
		if err != nil {
			return false, err
		}
		hash := parts[5]
		decryptedHash, err := ur.decrypt(hash)
		if err != nil {
			return false, err
		}

		var keyLen = uint32(len(decryptedHash))
		comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, uint8(parallerism), keyLen)
		return subtle.ConstantTimeCompare(comparisonHash, decryptedHash) == 1, nil

	}

	return false, nil

}
