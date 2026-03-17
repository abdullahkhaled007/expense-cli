package expense

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddExpense(t *testing.T) {
	store := NewStore()
	e, err := NewExpense("Test Expense", 100.0, "This is a test expense")
	require.NoError(t, err)

	result, err := store.Add(e)
	require.NoError(t, err)

	assert.Equal(t, 1, result.ID)
}

func TestAddNilExpense(t *testing.T) {
	store := NewStore()
	_, err := store.Add(nil)
	assert.Error(t, err)
}

func TestAddTwoExpenses(t *testing.T) {
	store := NewStore()
	e1, err := NewExpense("Test Expense 1", 100.0, "This is a test expense")
	require.NoError(t, err)

	e2, err := NewExpense("Test Expense 2", 200.0, "This is the second expense")
	require.NoError(t, err)

	result1, err := store.Add(e1)
	require.NoError(t, err)
	result2, err := store.Add(e2)
	require.NoError(t, err)

	assert.EqualValues(t, 1, result1.ID)
	assert.EqualValues(t, 2, result2.ID)
}

func TestList(t *testing.T) {
	store := NewStore()
	currentList, err := store.List()
	require.NoError(t, err)
	assert.Equal(t, 0, len(currentList))
}

func TestListAfterAdd(t *testing.T) {
	store := NewStore()
	e, err := NewExpense("Test Expense", 100.0, "This is a test expense")
	require.NoError(t, err)
	_, err = store.Add(e)
	require.NoError(t, err)

	currentList, err := store.List()
	require.NoError(t, err)

	assert.Len(t, currentList, 1)
	assert.EqualValues(t, "Test Expense", currentList[0].Name)
}

func TestListAfterDelete(t *testing.T) {
	store := NewStore()
	e, err := NewExpense("Test Expense", 100.0, "This is a test expense")
	require.NoError(t, err)
	_, err = store.Add(e)
	require.NoError(t, err)

	_, err = store.Delete(1)
	require.NoError(t, err)
	currentList, err := store.List()
	require.NoError(t, err)
	assert.Len(t, currentList, 0)
}

func TestDeleteMissingIndex(t *testing.T) {
	store := NewStore()
	_, err := store.Delete(999)
	assert.Error(t, err)
}
