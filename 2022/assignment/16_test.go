package assignment

import "testing"

func Test_d16Opened_addOpened(t *testing.T) {
	open := d16Opened{}
	vars := []uint8{50, 20, 30, 3}

	for _, id := range vars {
		open.add(id)
	}

	t.Log(open.opened[:open.openSize])
	if open.openSize != 4 {
		t.Fatal("opened != 4")
	}
	if open.opened[3] != 3 {
		t.Fatal("opened[0] wrong")
	}
	if open.opened[2] != 20 {
		t.Fatal("opened[1] wrong")
	}
	if open.opened[1] != 30 {
		t.Fatal("opened[2] wrong")
	}
	if open.opened[0] != 50 {
		t.Fatal("opened[3] wrong")
	}

	for _, id := range vars {
		if !open.contains(id) {
			t.Fatalf("%d not found", id)
		}
	}

	if open.contains(0) {
		t.Fatalf("0 found")
	}
}
