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

type sequence [4]int64

func Part2(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	//maps command sequence to the exchange result for each buyer
	exchangeRates := map[sequence]map[int]int64{}

	changes := make([]sequence, len(input.Secrets))
	secrets := input.Secrets

	for round := range 2000 {
		newSecrets := make([]int64, 0, len(secrets))
		for seller, secret := range secrets {
			newSecret := NextSecret(secret)
			newSecrets = append(newSecrets, newSecret)

			priceChange := Price(newSecret) - Price(secret)

			oldChange := changes[seller]
			newChange := sequence{oldChange[1], oldChange[2], oldChange[3], priceChange}
			changes[seller] = newChange

			// fmt.Printf("round: %d, seller: %d, sequence: %d\n", round, seller, newChange)
			if round >= 4 {
				exchangeRate := exchangeRates[newChange]
				if exchangeRate == nil {
					exchangeRate = map[int]int64{}
				}

				if _, ok := exchangeRate[seller]; !ok {
					exchangeRate[seller] = Price(newSecret)
				}
				exchangeRates[newChange] = exchangeRate
			}
		}

		secrets = newSecrets
	}

	maxBananas := int64(0)
	for _, rate := range exchangeRates {
		bananasCount := int64(0)
		for _, price := range rate {
			bananasCount += price
		}
		if bananasCount > maxBananas {
			maxBananas = bananasCount
		}
	}

	return fmt.Sprintf("%d", maxBananas), nil
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

func Price(secret int64) int64 {
	return secret % 10
}
