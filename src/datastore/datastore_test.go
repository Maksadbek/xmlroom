package datastore

import (
	"encoding/json"
	"testing"
)

func TestCreateMemberWithItems(t *testing.T) {
	err := CreateMemberWithItems(testMembers)
	if err != nil {
		t.Error(err)
	}

	items, err := ReadItemsByMember(testMembers[0].EstateAgentID)
	if err != nil {
		t.Error(err)
	}

	j1, err := json.Marshal(items)
	if err != nil {
		t.Error(err)
	}
	j2, err := json.Marshal(testMembers[0].Items)
	if err != nil {
		t.Error(err)
	}
	if string(j1) != string(j2) {
		t.Errorf("members are not equal")
	}
}
