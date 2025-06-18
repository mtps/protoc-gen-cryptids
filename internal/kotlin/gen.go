package kotlin

import (
	"fmt"
)

func EProvider(param string) string {
	return fmt.Sprintf("encryptionProvider(%s)", param)
}

func DProvider(param string) string {
	return fmt.Sprintf("decryptionProvider(%s)", param)
}
