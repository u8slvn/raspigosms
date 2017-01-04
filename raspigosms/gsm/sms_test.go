package gsm

import "testing"

// wip
func TestCheckPhoneFormat(t *testing.T) {
	t.Log("Coucou")
	if checkPhoneFormat("abcd") {
		t.Errorf("Wrong")
	}
}
