package patterns_test

import (
	"testing"

	"github.com/nulln0ne/tinyrebuilder/patterns"
)

func TestPredefinedEmail(t *testing.T) {
	re := patterns.Email().MustCompile()
	good := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"x@y.z",
	}
	bad := []string{
		"plainaddress",
		"@missingusername.com",
		"username@.com",
	}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match email %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid email %q", s)
		}
	}
}

func TestIPv4(t *testing.T) {
	re := patterns.IPv4().MustCompile()
	good := []string{
		"192.168.0.1",
		"0.0.0.0",
		"255.255.255.255",
	}
	bad := []string{
		"256.0.0.1",
		"192.168.1",
		"abc.def.ghi.jkl",
	}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match IPv4 %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid IPv4 %q", s)
		}
	}
}

func TestUUID(t *testing.T) {
	re := patterns.UUID().MustCompile()
	good := "550e8400-e29b-41d4-a716-446655440000"
	bad := "550e8400e29b41d4a716446655440000"
	if !re.IsMatch(good) {
		t.Errorf("Expected to match UUID %q", good)
	}
	if re.IsMatch(bad) {
		t.Errorf("Expected NOT to match invalid UUID %q", bad)
	}
}

func TestHexColor(t *testing.T) {
	re := patterns.HexColor().MustCompile()
	good := []string{"#fff", "#123ABC", "#abcdef"}
	bad := []string{"fff", "123456", "#1234", "#12345G"}

	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match HexColor %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid HexColor %q", s)
		}
	}
}

func TestURL(t *testing.T) {
	re := patterns.URL().MustCompile()
	good := []string{
		"http://example.com",
		"https://example.com/path?query=1",
		"http://www.example.com:8080",
	}
	bad := []string{
		"ftp://example.com",
		"example.com",
	}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match URL %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid URL %q", s)
		}
	}
}

func TestDate_YYYYMMDD(t *testing.T) {
	re := patterns.Date_YYYYMMDD().MustCompile()
	good := []string{"2025-07-21", "1999-12-31"}
	bad := []string{"21-07-2025", "2025/07/21", "2025-13-01"}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match Date %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid Date %q", s)
		}
	}
}

func TestTime_HHMMSS(t *testing.T) {
	re := patterns.Time_HHMMSS().MustCompile()
	good := []string{"12:30:00", "23:59:59", "00:00:00"}
	bad := []string{"24:00:00", "12:60:00", "12:30"}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match Time %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid Time %q", s)
		}
	}
}

func TestUsername(t *testing.T) {
	re := patterns.Username().MustCompile()
	good := []string{"user_123", "gemini", "a_b"}
	bad := []string{"us", "user-name", "user!", "averylongusername_12345"}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match Username %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid Username %q", s)
		}
	}
}

func TestSlug(t *testing.T) {
	re := patterns.Slug().MustCompile()
	good := []string{"hello-world", "slug", "a-b-c"}
	bad := []string{"-start-with-dash", "end-with-dash-", "double--dash"}
	for _, s := range good {
		if !re.IsMatch(s) {
			t.Errorf("Expected to match Slug %q", s)
		}
	}
	for _, s := range bad {
		if re.IsMatch(s) {
			t.Errorf("Expected NOT to match invalid Slug %q", s)
		}
	}
}
