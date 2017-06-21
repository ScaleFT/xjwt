package xjwt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomNonce(t *testing.T) {
	rn := RandomNonce{Size: 10}
	n, err := rn.Nonce()
	require.NoError(t, err)
	require.Len(t, n, 20)

	rn2, err := rn.Nonce()
	require.NoError(t, err)
	require.Len(t, rn2, 20)

	require.NotEqual(t, rn, rn2)
}
