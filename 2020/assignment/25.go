package assignment

import (
	"fmt"
	"strconv"
	"strings"
)

type Day25 struct{}

func (d *Day25) getKeys(input string) (int64, int64) {
	split := strings.Split(input, "\n")

	cardPK, err := strconv.ParseInt(split[0], 10, 32)
	CheckErr(err)

	doorPK, err := strconv.ParseInt(split[1], 10, 32)
	CheckErr(err)
	return cardPK, doorPK
}

func (d *Day25) transformOnce(val, subject int64) int64 {
	return (val * subject) % 20201227
}

func (d *Day25) transform(subject int64, loopSz int) int64 {
	val := int64(1)
	for i := 0; i <= loopSz; i++ {
		val = d.transformOnce(val, subject)
	}
	return val
}

func (d *Day25) findLoopSize(subject, publicKey int64) int {
	const trialSize = 100_000_000
	val := int64(1)
	for i := 0; i < trialSize; i++ {
		val = d.transformOnce(val, subject)
		if val == publicKey {
			return i
		}
	}
	panic("no loop size found")
}

func (d *Day25) SolveI(input string) int64 {
	cardPK, doorPK := d.getKeys(input)

	cardLoopSz := d.findLoopSize(7, cardPK)
	doorLoopSz := d.findLoopSize(7, doorPK)

	fmt.Println(cardLoopSz, doorLoopSz)

	encKey1 := d.transform(doorPK, cardLoopSz)
	encKey2 := d.transform(cardPK, doorLoopSz)

	if encKey1 != encKey2 {
		panic("encryption keys are not equal")
	}

	return encKey1
}

func (d *Day25) SolveII(input string) int64 {
	return 0
}
