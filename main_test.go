package main

import "testing"

func TestQuickSelect(t *testing.T) {
	input := []int{7, 4, 6, 3, 9, 1}
	Ks := []int{2, 1}
	output := []int{3, 1}

	for i, k := range Ks {
		got := QuickSelect(input, k)

		if got != output[i] {
			t.Errorf("QuickSelect(%v, %d) = %d; want %d",
				input, k, got, output[i])
		}

		t.Logf("QuickSelect(%v, %d) = %d; want %d",
			input, k, got, output[i])
	}
}
