package data

import (
	"testing"
)

func TestMarshalFromFile(t *testing.T) {
	ad := AuditData{}
	ad.MarshalFromFile("./testdata/valid-polaris-report.json")

	got := ad.Score
	want := uint(26)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	err := ad.MarshalFromFile("doesnotexist")
	if err == nil {
		t.Errorf("No error was thrown for missing polaris report")
	}
}
