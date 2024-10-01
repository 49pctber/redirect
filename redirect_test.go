package redirect

import (
	"testing"
)

func TestRedirect(t *testing.T) {
	_, err := NewRedirect("test", "https://google.com")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewRedirect("test", "https://g oogle.com")
	if err == nil {
		t.Fatal("this should have thrown an error")
	}

	_, err = NewRedirect("test:", "https://google.com")
	if err == nil {
		t.Fatal("this should have thrown an error")
	}
}
