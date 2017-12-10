// +build functest common pkg cfg

// Package cfg :: aws_decrypt_test.go
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDecryptKeyTextByKMS_functest tests common.DecryptKeyTextByKMS
func TestDecryptKeyTextByKMS_functest(t *testing.T) {
	// cipher is encrypted by KMS key (in AWS atg-infoblox account):
	//   - AliasName: alias/cyberintel
	//   - KeyId: arn:aws:kms:us-east-1:405093580753:key/336e1b0b-39da-44ad-b565-e6fe51a1b810
	//   - AliasArn: arn:aws:kms:us-east-1:405093580753:alias/cyberintel
	cipher := "AQICAHg3HNyIwRj/VgA+LeTSbBvD+KLqwR4I7fTUIoefURORLQF4TeZG1NW7J1fYok4dBM83AAAAYjBgBgkqhkiG9w0BBwagUzBRAgEAMEwGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMW7xGM12GTODwP4JRAgEQgB+Aq191ktDJAFXQrJASRNm8oj4q21rYVlj6OkOOhoaz"
	envKey := "api_key_or_password"
	envVal := "test"
	result := DecryptKeyTextByKMS(envKey, cipher)

	// t.Logf("Environment: %s\n", os.Environ())
	t.Logf("CiphertextBlob (base64): %v\n", cipher)
	assert.Equal(t, envVal, result)
}
