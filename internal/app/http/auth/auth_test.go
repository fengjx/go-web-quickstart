package auth

import "testing"

func TestJwt(t *testing.T) {
	tokenString, err := Signed(1000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokenString)
	uid, err := Parse(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uid)
}
