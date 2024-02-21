package memory

import (
	"testing"

	"github.com/rarimo/go-merkletree/db"
	"github.com/rarimo/go-merkletree/db/test"
	"github.com/stretchr/testify/require"
)

func TestMemoryStorageInterface(t *testing.T) {
	var db db.Storage //nolint:gosimple

	db = NewMemoryStorage()
	require.NotNil(t, db)
}

func TestMemory(t *testing.T) {
	test.TestReturnKnownErrIfNotExists(t, NewMemoryStorage())
	test.TestStorageInsertGet(t, NewMemoryStorage())
	test.TestStorageWithPrefix(t, NewMemoryStorage())
	test.TestConcatTx(t, NewMemoryStorage())
	test.TestList(t, NewMemoryStorage())
	test.TestIterate(t, NewMemoryStorage())
}
