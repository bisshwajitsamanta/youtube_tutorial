package demo

import "testing"

func Test_UpdateMessage(t *testing.T) {

	wg.Add(1)
	go UpdateMessage("Hello Bisshwajit")
	//go UpdateMessage("Hello Youtube")
	wg.Wait()

	if msg != "Hello Bisshwajit" {
		t.Errorf("Incorrect value in msg")
	}
}
