package debug

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		debug bool
		t     string
	}{
		{
			debug: false,
			t:     "debug.NopLog",
		},
		{
			debug: true,
			t:     "debug.Log",
		},
	}

	for _, tc := range cases {
		resType := reflect.TypeOf(New(tc.debug)).Elem().String()

		if resType != tc.t {
			t.Errorf("Type must be %s but got %s", tc.t, resType)
		}
	}
}

func TestNopLogDebug(t *testing.T) {
	cases := []struct {
		layout string
		args   []interface{}
		res    string
	}{
		{
			layout: "example",
			res:    "",
		},
		{
			layout: "example %s",
			args:   []interface{}{"foo"},
			res:    "",
		},
	}

	oldStdout := os.Stdout
	oldStderr := os.Stderr

	for _, tc := range cases {
		r, w, err := os.Pipe()
		if err != nil {
			t.Error(err)
			continue
		}

		os.Stdout = w
		os.Stderr = w

		log.SetOutput(w)

		n := &NopLog{}
		n.Debug(tc.layout, tc.args...)

		w.Close()

		by, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
			continue
		}

		if tc.res != string(by) {
			t.Errorf("Res must be empty but got %s", string(by))
		}
	}

	os.Stdout = oldStdout
	os.Stderr = oldStderr
}

func TestLogDebug(t *testing.T) {
	cases := []struct {
		layout string
		args   []interface{}
		res    string
	}{
		{
			layout: "example",
			res:    "example",
		},
		{
			layout: "example %s",
			args:   []interface{}{"foo"},
			res:    "example foo",
		},
	}

	oldStdout := os.Stdout
	oldStderr := os.Stderr

	for _, tc := range cases {
		r, w, err := os.Pipe()
		if err != nil {
			t.Error(err)
			continue
		}

		os.Stdout = w
		os.Stderr = w

		log.SetOutput(w)

		l := &Log{}
		l.Debug(tc.layout, tc.args...)

		w.Close()

		by, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
			continue
		}

		if strings.HasSuffix(string(by), tc.res) {
			t.Errorf("Res (%s) must has suffix %s", string(by), tc.res)
		}
	}

	os.Stdout = oldStdout
	os.Stderr = oldStderr
}
