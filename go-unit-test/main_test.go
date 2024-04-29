package main

import (
	"testing"
)

func TestAdd(t *testing.T)  {

	testCases := []struct {
    name     string
    a, b     int
    expected int
  }{
    {"Add positive numbers", 2, 3, 5},
    {"Add negative numbers", -1, -2, -3},
    {"Add zero", 0, 0, 0},
  }


	for _ , tc := range testCases {
		t.Run(tc.name , func(t *testing.T) {
			result := add(tc.a , tc.b)
			expectedresult := tc.expected
			if result != expectedresult {
				t.Errorf("Add(%d,%d) = %d is wrong correct is %d" , tc.a , tc.b , result , expectedresult)
			}
		})
	}
}