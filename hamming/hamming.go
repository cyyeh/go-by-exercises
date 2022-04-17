package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("a, b should be equal size")
    }

    hammingDistance := 0
    for i :=0; i < len(a); i++ {
        if a[i] != b[i] {
            hammingDistance += 1
        }
    }

    return hammingDistance, nil
}
