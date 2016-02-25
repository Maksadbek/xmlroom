package models

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	members := []Member{}
	err := xml.Unmarshal([]byte(testXML), &members)
	if err != nil {
		t.Error(err)
	}
	testTime, err := time.Parse(FORM, FORM)
	if err != nil {
		t.Error(err)
	}
	testMembers[0].Items[0].Available = Date(testTime)
	testMembers[0].Items[0].InsertDate = Date(testTime)

	if reflect.DeepEqual(members[0].Items[0], testMembers[0].Items[0]) {
		t.Error("members are not equal")
	}
}
