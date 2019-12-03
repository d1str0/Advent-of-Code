package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	m1 := 12
	m2 := 14
	m3 := 1969
	m4 := 100756

	f1 := 2
	f2 := 2
	f3 := 654
	f4 := 33583

	if calc(m1) != f1 {
		t.Errorf("Calculated %d, got %d, expected %d", m1, calc(m1), f1)
	}
	if calc(m2) != f2 {
		t.Errorf("Calculated %d, got %d, expected %d", m2, calc(m2), f2)
	}
	if calc(m3) != f3 {
		t.Errorf("Calculated %d, got %d, expected %d", m3, calc(m3), f3)
	}
	if calc(m4) != f4 {
		t.Errorf("Calculated %d, got %d, expected %d", m4, calc(m4), f4)
	}

	t1 := 2
	t2 := 2
	t3 := 966
	t4 := 50346

	if total(m1) != t1 {
		t.Errorf("Calculated %d, got %d, expected %d", m1, calc(m1), t1)
	}
	if total(m2) != t2 {
		t.Errorf("Calculated %d, got %d, expected %d", m2, calc(m2), t2)
	}
	if total(m3) != t3 {
		t.Errorf("Calculated %d, got %d, expected %d", m3, calc(m3), t3)
	}
	if total(m4) != t4 {
		t.Errorf("Calculated %d, got %d, expected %d", m4, calc(m4), t4)
	}

}
