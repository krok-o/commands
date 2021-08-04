package pkg

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifySlack(t *testing.T) {
	content, err := ioutil.ReadFile(filepath.Join("testdata", "create.json"))
	assert.NoError(t, err)

	n := Notifier{Payload: string(content)}
	err = n.Notify()
	assert.NoError(t, err)
}
