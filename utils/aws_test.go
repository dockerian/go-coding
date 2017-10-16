// +build all utils aws

// Package utils :: aws_test.go
package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/stretchr/testify/assert"
)

var (
	_configDecryptKeyTestCases = []ConfigDecryptKeyTestCase{
		{"original text", "original text"},
		{"test", "AQICAHg3HNyIwRj/VgA+LeTSbBv"},
		{"test", "TEST/TEST/TEST/TEST/TEST"},
		{"text", "MORE_MOCKED_CIPHER_TEXT"},
	}
)

// ConfigDecryptTestCase struct
type ConfigDecryptKeyTestCase struct {
	originalText string
	beforeMocked string
}

// MockKmsService mocks KMS struct
type MockKmsService struct {
}

// Decrypt method mocks
func (m *MockKmsService) Decrypt(input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	inputString := string(input.CiphertextBlob)
	decryptTest := &kms.DecryptOutput{Plaintext: input.CiphertextBlob}
	for _, test := range _configDecryptKeyTestCases {
		encrypted := string(mockCiphertext(test.beforeMocked))
		if inputString == encrypted {
			decryptTest.Plaintext = []byte(test.originalText)
			break
		}
	}
	return decryptTest, nil
}

// mockCiphertext mocks cipher text by 10x repeating and encoding a short input
func mockCiphertext(input string) string {
	repeatText := strings.Repeat(input, 10)
	return base64.StdEncoding.EncodeToString([]byte(repeatText))
}

// TestDecryptKeyTextByKMS tests common.DecryptKeyTextByKMS
func TestDecryptKeyTextByKMS(t *testing.T) {
	apiKey := "SOME_API_KEY_TO_ENCRYPT"
	config := Config{settings: map[string]string{}}

	// save original service _kmsService (in aws.go)
	_savedKmsService := _kmsService
	_kmsService = &MockKmsService{}

	for index, test := range _configDecryptKeyTestCases {
		keyVal := apiKey
		if index%2 == 0 {
			keyVal = "SOME_ENCRYPTED_PASSWORD"
		}
		cipherText := mockCiphertext(test.beforeMocked)
		expected := strings.Repeat(test.beforeMocked, 10)
		os.Setenv(keyVal, cipherText)
		result := config.Get(keyVal)
		tstMsg := fmt.Sprintf("config[%s]: %s", keyVal, test.beforeMocked)
		errMsg := fmt.Sprintf("config[%s]: %s --> %s (actual: %s)",
			keyVal, cipherText, test.originalText, result)
		t.Logf("Test %2d: %v\n", index+1, tstMsg)
		assert.Equal(t, expected, result, errMsg)
	}
	// restore from saved service
	_kmsService = _savedKmsService
}
