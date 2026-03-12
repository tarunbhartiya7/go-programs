package email

import "testing"

func TestIsValidEmail(t *testing.T) {
	if !IsValidEmail("test@example.com") {
		t.Error("expected valid for test@example.com")
	}

	if IsValidEmail("invalid-email") {
		t.Error("expected invalid for invalid-email")
	}

	if IsValidEmail("") {
		t.Error("expected invalid for empty string")
	}
}

func TestIsValidEmailTable(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"", false},
	}

	for _, data := range tests {
		got := IsValidEmail(data.email)
		if got != data.want {

			t.Errorf("IsValidEmail(%q) = %v, want %v", data.email, got, data.want)
		}
	}
}
