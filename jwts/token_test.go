package jwts

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	token := Token{Key: "go-flysystem", ExpiresIn: 7200 * time.Second}
	resp, err := token.BuildToken("go-flysystem", "test", "test-bucket")
	if err != nil {
		t.Fatal(err)
	}
	parseToken, err := token.ParseToken(resp.Token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parseToken.Disk)
	t.Log(parseToken.Bucket)
}

func TestTokenExpTimeAdd(t *testing.T) {
	token := Token{Key: "go-flysystem", ExpiresIn: 1 * time.Second}
	resp, err := token.BuildToken("go-flysystem", "test", "test-bucket")
	if err != nil {
		t.Fatal(err)
	}
	tokenStr := resp.Token
	time.Sleep(2 * time.Second)
	_, err = token.ParseToken(tokenStr)
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok && e.Errors != jwt.ValidationErrorExpired {
			t.Fatal(err)
		} else {
			t.Log(err)
		}
	}
}
