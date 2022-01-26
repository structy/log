package file

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/structy/log"
)

func TestFileWrite(t *testing.T) {
	now = func() time.Time { return time.Unix(1498405744, 0) }

	fileWrite(
		log.ErrorLog,
		log.LineOut,
		map[string]interface{}{"fileName": "logfile.txt"},
		"test log")
	fileWrite(
		log.DebugLog,
		log.LineOut,
		map[string]interface{}{"fileName": "logfile.txt"},
		"test log")
	fileWrite(
		log.WarningLog,
		log.LineOut,
		map[string]interface{}{"fileName": "logfile.txt"},
		"test log")

	b, err := ioutil.ReadFile("logfile.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = os.Remove("logfile.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	expectd := "2017/06/25 15:49:04 [error] test log\n2017/06/25 15:49:04 [warning] test log\n"
	if string(b) != expectd {
		t.Fatalf("Error expectd %q, got %q\n", expectd, string(b))
	}
}

func TestDebugFileWrite(t *testing.T) {
	now = func() time.Time { return time.Unix(1498405744, 0) }

	log.DebugMode = true
	fileWrite(
		log.DebugLog,
		log.LineOut,
		map[string]interface{}{"fileName": "logfile-debug.txt"},
		"test debug log")

	b, err := ioutil.ReadFile("logfile-debug.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = os.Remove("logfile-debug.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	expectd := "2017/06/25 15:49:04 [debug] .:0 test debug log\n"
	if string(b) != expectd {
		t.Fatalf("Error expectd %q, got %q\n", expectd, string(b))
	}
}
