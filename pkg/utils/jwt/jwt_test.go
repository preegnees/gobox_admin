package jwt

import (
	"os"
	"testing"
)

const SECREST = "test"
const USERNAME = "my_name"

var jwt_ IJWT = nil

func TestMain(m *testing.M) {
	os.Setenv("SECRET", SECREST)
	jwt_, _ = New()
	m.Run()
}

func TestGenJWT(t *testing.T) {
	aToken, rToken, err := jwt_.GenerateJWT(USERNAME, User)
	if err != nil || aToken == "" || rToken == "" {
		t.Fatal(err)
	}
}

func TestCheckJWT(t *testing.T) {
	b := jwt_.CheckJwt("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2Njg3OTI3NjgsImV4cCI6NDg1NjAwMjM2OCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.5yE8QDPshNnVCR5GX3HT6_TtpNd2KsU6ywpf_MQUDQg")
	if !b {
		t.Fatal("!b")
	}
}

func TestGetValues(t *testing.T) {
	aToken, _, _ := jwt_.GenerateJWT(USERNAME, User)
	myCl, err := jwt_.GetValuesFromJWT(aToken)
	if err != nil {
		t.Fatal(err)
	}

	if myCl.Username != USERNAME {
		t.Fatal("myCl.Username != USERNAME")
	}
}
