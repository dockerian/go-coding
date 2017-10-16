// Package utils :: aws.go - extended AWS SDK functions
package utils

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var (
	_kmsService KMSDecryptInterface = kms.New(session.New(),
		&aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
)

// KMSDecryptInterface interface
type KMSDecryptInterface interface {
	Decrypt(input *kms.DecryptInput) (*kms.DecryptOutput, error)
}

// DecryptKeyTextByKMS checks possible encrypted KMS key/value and retruns decrypted text
func DecryptKeyTextByKMS(key, text string) string {
	keyLowers := strings.ToLower(key)
	keyOrPass := strings.Contains(keyLowers, "api_key") || strings.Contains(keyLowers, "password")
	encrypted := len(text) > 128 && !strings.Contains(text, " ")
	inputBlob := []byte(text)

	if keyOrPass && encrypted {
		if base64Blob, err := base64.StdEncoding.DecodeString(text); err == nil {
			inputBlob = base64Blob
		}
		deInput := &kms.DecryptInput{CiphertextBlob: inputBlob}
		if decrypt, err := _kmsService.Decrypt(deInput); err == nil {
			return string(decrypt.Plaintext)
		}
	}
	return text
}
