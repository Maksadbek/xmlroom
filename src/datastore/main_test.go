package datastore

import (
	"log"
	"os/exec"
	"strings"
	"testing"
)

// database for testing purposes
const TESTDB string = "./testdb"

func TestMain(m *testing.M) {
	err := Init(TESTDB)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
	cmd := exec.Command("rm", TESTDB)
	cmd.Stdin = strings.NewReader(TESTDB)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
