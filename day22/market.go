package day22

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Input struct {
	Secrets []int64
}

func ParseInput(in io.Reader) (*Input, error) {
	secrets := []int64{}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		secretText := scanner.Text()
		secret, err := strconv.ParseInt(secretText, 10, 64)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, secret)
	}

	return &Input{Secrets: secrets}, nil
}

func Part1(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	secrets := input.Secrets

	for _ = range 2000 {
		newSecrets := make([]int64, 0, len(secrets))
		for _, secret := range secrets {
			newSecrets = append(newSecrets, NextSecret(secret))
		}
		secrets = newSecrets
	}

	sum := int64(0)
	for _, secret := range secrets {
		sum += secret
	}
	return fmt.Sprintf("%d", sum), nil
}

func NextSecret(secret int64) int64 {
	mul := secret * 64
	secret = mul ^ secret
	secret = secret % 16777216

	divided := secret / 32
	secret = secret ^ divided
	secret = secret % 16777216

	multiplied := secret * 2048
	secret = secret ^ multiplied
	secret = secret % 16777216

	return secret
}
