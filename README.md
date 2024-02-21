# go-merkletree [![GoDoc](https://godoc.org/github.com/rarimo/go-merkletree?status.svg)](https://godoc.org/github.com/rarimo/go-merkletree) [![Go Report Card](https://goreportcard.com/badge/github.com/rarimo/go-merkletree)](https://goreportcard.com/report/github.com/rarimo/go-merkletree) [![Test](https://github.com/rarimo/go-merkletree/workflows/Test/badge.svg)](https://github.com/rarimo/go-merkletree/actions?query=workflow%3ATest)

MerkleTree compatible with version from [circomlib](https://github.com/iden3/circomlib).

Adaptation of the merkletree from https://github.com/iden3/go-iden3-core/tree/v0.0.8 with several changes and more functionalities.

## Usage
More detailed examples can be found at the [tests](https://github.com/rarimo/go-merkletree/blob/master/merkletree_test.go), and in the [documentation](https://godoc.org/github.com/rarimo/go-merkletree).

```go
import (
	"fmt"
	"math/big"
	"testing"

	"github.com/iden3/go-iden3-core/db"
	"github.com/stretchr/testify/assert"
)

[...]

func TestExampleMerkleTree(t *testing.T) {
	mt, err := NewMerkleTree(db.NewMemoryStorage(), 10)
	assert.Nil(t, err)

	key := big.NewInt(1)
	value := big.NewInt(2)
	err = mt.Add(key, value)
	assert.Nil(t, err)
	fmt.Println(mt.Root().String())

	v, err := mt.Get(key)
	asseert.Equal(t, value, v)

	value = big.NewInt(3)
	err = mt.Update(key, value)

	proof, err := mt.GenerateProof(key, nil)
	assert.Nil(t, err)

	assert.True(t, VerifyProof(mt.Root(), proof, key, value))

	err := mt.Delete(big.NewInt(1)) // delete the leaf of key=1
	assert.Nil(t, err)
}
```
