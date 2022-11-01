package dbService

import (
	"github.com/fatih/color"
	"testing"
)

func Test_InitDB(t *testing.T) {
	err := initDB()
	if err != nil {
		color.Red("Failed to initialize connection\n", err)
	}
}
