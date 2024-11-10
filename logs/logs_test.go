package logs

import (
	"context"
	"testing"
)

func TestLogs(t *testing.T) {
	Init(WithDevelopment(), WithWriterConsole())
	Infot(context.Background(), "hello")
}
