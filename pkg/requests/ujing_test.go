package requests

import "testing"

func TestUJingSession_Captcha(t *testing.T) {
	session := NewUJingSession()
	_ = session.Captcha()
}
