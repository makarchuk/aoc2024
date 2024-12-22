package day22_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day22"
	"github.com/stretchr/testify/require"
)

func TestSecretChanges(t *testing.T) {
	sequence := []int64{
		123,
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}

	secret := sequence[0]
	for i := 1; i < len(sequence); i++ {
		nextSecret := day22.NextSecret(secret)
		require.EqualValues(t, sequence[i], nextSecret)
		secret = nextSecret
	}
}
