package main

import "testing"

func TestIsPrime(t *testing.T) {
	//Arrange
	sut := 11
	expected := true

	//Act
	actual := IsPrime(sut)

	//Assert
	if actual != expected {
		t.Fail()
	}
}

//data driven tests
type PrimeTestCase struct {
	name     string
	no       int
	expected bool
}

func TestAllPrimeNumbers(t *testing.T) {
	//Arrange
	testCases := []PrimeTestCase{
		{name: "Testing Prime for 11", no: 11, expected: true},
		{name: "Testing Prime for 21", no: 21, expected: false},
		{name: "Testing Prime for 33", no: 33, expected: false},
		{name: "Testing Prime for 44", no: 44, expected: false},
		{name: "Testing Prime for 67", no: 67, expected: true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsPrime(testCase.no)
			if actual != testCase.expected {
				t.Fail()
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(11)
	}
}
