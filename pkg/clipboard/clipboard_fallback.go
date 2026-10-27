//go:build !darwin

package clipboard

import "errors"

var errUnsupported = errors.New("clipboard is unsupported on this platform (requires macOS)")

type ClipboardContent struct {
	Type     string
	Data     []byte
	IsText   bool
	IsFile   bool
	FilePath string
}

func CopyFile(_ string) error {
	return errUnsupported
}

func CopyFiles(_ []string) error {
	return errUnsupported
}

func CopyText(_ string) error {
	return errUnsupported
}

func CopyRichText(_ string, _ []byte, _ string) error {
	return errUnsupported
}

func CopyTeamsRichText(_ string, _ []byte) error {
	return errUnsupported
}

func CopySlackMessage(_ string, _ string) error {
	return errUnsupported
}

func CopyTextWithType(_ string, _ string) error {
	return errUnsupported
}

func Clear() error {
	return errUnsupported
}

func GetFiles() []string {
	return nil
}

func GetText() (string, bool) {
	return "", false
}

func GetUTIForFile(_ string) (string, bool) {
	return "", false
}

func GetClipboardTypes() []string {
	return nil
}

func GetClipboardDataForType(_ string) ([]byte, bool) {
	return nil, false
}

func ContainsType(_ string) bool {
	return false
}

func UTIConformsTo(_ string, _ string) bool {
	return false
}

func GetPreferredExtensionForUTI(_ string) string {
	return ""
}

func SaveRTFDToPath(_ []byte, _ string) error {
	return errUnsupported
}

func GetClipboardContent() (*ClipboardContent, error) {
	return nil, errUnsupported
}
