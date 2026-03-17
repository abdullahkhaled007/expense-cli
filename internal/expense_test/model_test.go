package expense_test

import (
	"golang/internal/expense"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExpense_success(t *testing.T) {
	e, err := expense.NewExpense("Test Expense", 100.0, "This is a test expense")
	require.NoError(t, err)
	require.NotNil(t, e)
	assert.Equal(t, "Test Expense", e.Name)
	assert.Equal(t, 100.0, e.Price)
	assert.False(t, e.Date.IsZero(), "expected date to be set")
}

func TestNewExpense_missingName(t *testing.T) {
	_, err := expense.NewExpense("", 100.0, "This is a test expense")
	require.Error(t, err)
	require.NotNilf(t, err, "expected error for empty name, got nil")
}

func TestNewExpense_negativePrice(t *testing.T) {
	e, err := expense.NewExpense("Test Expense", -50.0, "This is a test expense")
	require.Error(t, err)
	require.Nil(t, e, "expected nil expense when error occurs, got non-nil")
}
