package metrics

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	customValidatorFile = "pubkeys.json"
)

func TestReadCustomValidator(t *testing.T) {
	keys, err := ReadCustomValidatorsFile(customValidatorFile)
	require.Equal(t, nil, err)
	require.Equal(t, 10, len(keys))

	fmt.Print(keys)
}
