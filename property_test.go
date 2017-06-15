package cor

import "testing"

func TestLoad(t *testing.T) {
	t.Log(Load("server"))
	t.Log(GetAddr())
	t.Log(GetPort())
}
