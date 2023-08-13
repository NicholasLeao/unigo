package unigo

import (
	"testing"
)

/*
*	Unigo.MakeSet
 */
func TestMakeSet(t *testing.T) {
	iuf := &Unigo[string]{}
	iuf.MakeSet("set1")

	val, exists := iuf.id["set1"]
	if !exists {
		t.Errorf("expected set to exist in id map")
	}
	if val != "set1" {
		t.Errorf("expected set identifier of 'set1' to be 'set1'")
	}
}
func TestMakeSetSizeIndexInitialization(t *testing.T) {
	iuf := &Unigo[string]{}
	iuf.MakeSet("set1")
	iuf.MakeSet("set2")
	iuf.MakeSet("set3")

	for key, val := range iuf.sz {
		if val != 1 {
			t.Errorf("expected key %s to have size index of 1 but got %d", key, val)
		}
	}
}

/*
*	Unigo.Find
 */
func TestFindOnExistingSet(t *testing.T) {
	iuf := &Unigo[string]{}
	iuf.MakeSet("set1")

	setId, err := iuf.Find("set1")

	if err != nil {
		t.Error("find returned error for existing set")
	}

	if setId != "set1" {
		t.Errorf("expected set id 'set1' but got %v", setId)
	}

}
func TestFindPathCompression(t *testing.T) {
	iuf := &Unigo[string]{}
	iuf.MakeSet("set1")
	iuf.MakeSet("set2")
	iuf.MakeSet("set3")
	iuf.MakeSet("set4")
	iuf.MakeSet("set5")
	iuf.Union("set1", "set2")
	iuf.Union("set2", "set3")
	iuf.Union("set3", "set4")
	iuf.Union("set4", "set5")

	iuf.Find("set5")

	for key, value := range iuf.id {
		if value != "set1" {
			t.Errorf("expected every key to point directly to set identifier but got %s for key %s", value, key)
		}
	}
}

func TestFindOnNonExistingSet(t *testing.T) {
	var zeroValue string
	iuf := &Unigo[string]{}

	setId, err := iuf.Find("set1")

	if err == nil {
		t.Error("expected error for find call on non existing set")
	}
	if setId != zeroValue {
		t.Errorf("expected set id 'set1' but got %v", setId)
	}
}

/*
*	Unigo.Union
 */
func TestUnionOnExistingSet(t *testing.T) {
	iuf := &Unigo[int]{}
	iuf.MakeSet(1)
	iuf.MakeSet(2)

	setId, err := iuf.Union(1, 2)
	foundSetId, _ := iuf.Find(1)
	foundSetId2, _ := iuf.Find(2)

	if err != nil {
		t.Errorf("expected no errors but received %v", err)
	}
	if foundSetId != foundSetId2 {
		t.Errorf("expected both keys to belong to the same set")
	}
	if setId != foundSetId || setId != foundSetId2 {
		t.Error("expected found set identifier to be the same as returned by Union")
	}
}
func TestUnionOnNonExistingSet(t *testing.T) {
	var zeroValue int
	iuf := &Unigo[int]{}

	setId, err := iuf.Union(1, 2)

	if err == nil {
		t.Error("expected Union to return error but got nil")
	}
	if setId != zeroValue {
		t.Errorf("expected new set id to be zero value but got %d", setId)
	}
}
func TestUnionTreeBalancing(t *testing.T) {}

/*
*	Unigo.Connected
 */
func TestConnectedOnExistingSet(t *testing.T) {
	iuf := &Unigo[int]{}
	iuf.MakeSet(1)
	iuf.MakeSet(2)
	iuf.MakeSet(3)
	iuf.Union(1, 2)
	iuf.Union(2, 3)

	connected, err := iuf.Connected(1, 3)

	if err != nil {
		t.Errorf("expected no error but go %v", err)
	}
	if !connected {
		t.Error("expected connected to be true but got false")
	}
}

func TestConnectedOnNonExistingSet(t *testing.T) {
	iuf := &Unigo[int]{}
	iuf.MakeSet(1)

	connected, err := iuf.Connected(1, 2)

	if err == nil {
		t.Error("expected an error but got none")
	}
	if connected {
		t.Errorf("expected connected to be false but got true")
	}
}
