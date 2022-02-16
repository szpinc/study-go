package channel

import "testing"

func TestMirrorQuery(t *testing.T) {
	resp := mirrorQuery()
	println("Resp:", resp)
}
