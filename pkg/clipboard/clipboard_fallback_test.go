//go:build !darwin

package clipboard

import "testing"

func TestUnsupportedPlatform(t *testing.T) {
	operations := map[string]func() error{
		"CopyFile":          func() error { return CopyFile("file") },
		"CopyFiles":         func() error { return CopyFiles([]string{"file"}) },
		"CopyText":          func() error { return CopyText("text") },
		"CopyRichText":      func() error { return CopyRichText("html", []byte("rtf"), "text") },
		"CopyTeamsRichText": func() error { return CopyTeamsRichText("html", []byte("rtf")) },
		"CopySlackMessage":  func() error { return CopySlackMessage("delta", "text") },
		"CopyTextWithType":  func() error { return CopyTextWithType("text", "public.plain-text") },
		"Clear":             Clear,
		"SaveRTFDToPath":    func() error { return SaveRTFDToPath([]byte("rtfd"), "file.rtfd") },
	}
	const want = "clipboard is unsupported on this platform (requires macOS)"
	for name, operation := range operations {
		t.Run(name, func(t *testing.T) {
			if err := operation(); err == nil || err.Error() != want {
				t.Errorf("%s() error = %v; want %q", name, err, want)
			}
		})
	}

	if content, err := GetClipboardContent(); content != nil || err == nil || err.Error() != want {
		t.Errorf("GetClipboardContent() = %v, %v; want nil, %q", content, err, want)
	}
}

func TestUnsupportedPlatformQueries(t *testing.T) {
	if files := GetFiles(); files != nil {
		t.Errorf("GetFiles() = %v; want nil", files)
	}
	if text, ok := GetText(); text != "" || ok {
		t.Errorf("GetText() = %q, %t; want empty, false", text, ok)
	}
	if uti, ok := GetUTIForFile("file.txt"); uti != "" || ok {
		t.Errorf("GetUTIForFile() = %q, %t; want empty, false", uti, ok)
	}
	if types := GetClipboardTypes(); types != nil {
		t.Errorf("GetClipboardTypes() = %v; want nil", types)
	}
	if data, ok := GetClipboardDataForType("public.plain-text"); data != nil || ok {
		t.Errorf("GetClipboardDataForType() = %v, %t; want nil, false", data, ok)
	}
	if ContainsType("public.plain-text") {
		t.Error("ContainsType() = true; want false")
	}
	if UTIConformsTo("public.plain-text", "public.text") {
		t.Error("UTIConformsTo() = true; want false")
	}
	if ext := GetPreferredExtensionForUTI("public.png"); ext != "" {
		t.Errorf("GetPreferredExtensionForUTI() = %q; want empty", ext)
	}
}
