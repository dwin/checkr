package checkr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// HashCheckReadCloser ...
func HashCheckReadCloser(apiKey, checkrSignature string, b io.ReadCloser) error {
	// Create HMAC
	h := hmac.New(sha256.New, []byte(apiKey))
	// Write
	if _, err := io.Copy(h, b); err != nil {
		return err
	}
	// Digest
	digest := hex.EncodeToString(h.Sum(nil))
	// Check calculated digest match Signature
	if digest != checkrSignature {
		return ErrInvalidSignature
	}
	return nil
}
