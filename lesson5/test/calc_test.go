package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func TestSum(t *testing.T) {
	actual := sum(3, 5)
	expected := 7

	// if actual != expected {
	// 	t.Errorf("got %v\nwant %v", actual, expected)
	// }

	assert.Equal(t, expected, actual)
}

func setUp() {
	fmt.Println("Do BEFORE the tests!")
}

func tearDown() {
	fmt.Println("Do After the tests!")
}
