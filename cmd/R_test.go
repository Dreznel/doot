package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestScratch(t *testing.T) {
	fmt.Println("Test begin")
	t.Log("Test begin T")
	t.Log(os.Getwd())
	createRSkeleton(false, false, "")

	answer := 1 + 1
	assert.Equal(t, 2, answer)
	fmt.Println("Test end")
	t.Log("Test end T")
}