package result

import (
  "testing"
  "github.com/stvp/assert"
)


type errorString struct {
  value string
}

func (err *errorString) Error() string {
  return err.value
}


func TestFailure(t *testing.T) {
  err := &errorString{"testing error"}
  result := NewFailure(err)

  assert.Nil(t, result.Success)
  assert.Equal(t, result.Failure, err)
}

func TestSuccess(t *testing.T) {
  value := 5
  result := NewSuccess(value)

  assert.Equal(t, result.Success, value)
  assert.Nil(t, result.Failure)
}


