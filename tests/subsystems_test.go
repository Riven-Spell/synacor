package tests

import (
	"strconv"
	"synacor/interpreter"
	"testing"
)

func TestGetValue(t *testing.T) {
	system := interpreter.NewSystem()
	system.Registers[0] = 500

	if val, err := system.GetValue(32768); val != 500 || err != nil {
		if err != nil {
			t.Error(err.Error())
		}
		t.Errorf("Expected a response of 500, got %d\n", val)
		t.Fail()
	}

	if val, err := system.GetValue(999); val != 999 || err != nil {
		if err != nil {
			t.Error(err.Error())
		}
		t.Errorf("Expected a response of 999, got %d\n", val)
		t.Fail()
	}

	if val, err := system.GetValue(32776); err == nil || val != 0 {
		t.Errorf("Expected a response of (0, error), got (%d, %s)", val, err.Error())
		t.Fail()
	}
}

func TestGetLocation(t *testing.T) {
	system := interpreter.NewSystem()
	system.Registers[0] = 500
	system.Memory[50] = 600

	if val, err := system.GetLocation(32768, false); *val != 500 || val != &system.Registers[0] || err != nil {
		if err != nil {
			t.Error(err.Error())
		}
		t.Errorf("Expected pointer location of %p, got pointer location of %p\n", &system.Registers[0], val)
		t.Errorf("Expected value @ pointer of %d, got value of %d\n", system.Registers[0], val)
		t.Fail()
	}

	if val, err := system.GetLocation(50, true); *val != 600 || val != &system.Memory[50] || err != nil {
		if err != nil {
			t.Error(err.Error())
		}
		t.Errorf("Expected pointer location of %p, got pointer location of %p\n", &system.Memory[50], val)
		t.Errorf("Expected value @ pointer of %d, got value of %d\n", system.Memory[50], val)
		t.Fail()
	}

	if val, err := system.GetLocation(32776, false); val != nil || err == nil {
		t.Errorf("Expected an error, got pointer value %p\n", val)
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	system := interpreter.NewSystem()
	system.LoadProgram([]byte{170, 170})

	if system.Memory[0] != 43690 {
		t.Errorf("Expected 43690 (1010101010101010), got " + strconv.Itoa(int(system.Memory[0])))
		t.Fail()
	}
}
