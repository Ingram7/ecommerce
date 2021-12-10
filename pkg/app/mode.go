package app

import "log"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

type Mode struct {
	val string
}

func NewMode(val string) *Mode {
	mode := new(Mode)
	mode.val = val

	if !mode.valid() {
		log.Fatalf("mode val error: %s", val)
	}

	return mode
}

func (m *Mode) IsDebug() bool {
	return m.val == DebugMode
}

func (m *Mode) IsRelease() bool {
	return m.val == ReleaseMode
}

func (m *Mode) IsTest() bool {
	return m.val == TestMode
}

func (m *Mode) String() string {
	return m.val
}

func (m *Mode) valid() bool {
	values := map[string]bool{
		DebugMode:   true,
		ReleaseMode: true,
		TestMode:    true,
	}

	if _, ok := values[m.val]; ok {
		return true
	}

	return false
}
