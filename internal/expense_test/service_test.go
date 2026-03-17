package expense

import (
	"golang/internal/expense"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceAdd(t *testing.T) {
	service := expense.NewService()

	addedExpense, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	assert.Equal(t, 1, addedExpense.ID)
	assert.Equal(t, "Test Expense", addedExpense.Name)
}

func TestServiceAdd2Expenses(t *testing.T) {
	service := expense.NewService()
	addedExpense1, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	assert.Equal(t, 1, addedExpense1.ID)
	addedExpense2, err := service.Add("Test Expense", 200.0, "A test")
	require.NoError(t, err)
	assert.Equal(t, 2, addedExpense2.ID)
}

func TestServiceAddEmptyName(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("", 100.0, "A test")
	require.Error(t, err)
}

func TestServiceAddNegValue(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("Test Expense", -100.0, "A test")
	require.Error(t, err)
}

func TestServiceEmptyList(t *testing.T) {
	service := expense.NewService()
	expenses, err := service.List()
	require.NoError(t, err)
	assert.Equal(t, 0, len(expenses))
}

func TestServiceList(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	expenses, err := service.List()
	require.NoError(t, err)
	assert.Equal(t, 1, len(expenses))
}

func TestServiceDelete(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	removedExpense, err := service.Delete(1)
	require.NoError(t, err)
	assert.Equal(t, 1, removedExpense.ID)
}

func TestServiceDeleteThenList(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	_, err = service.Delete(1)
	require.NoError(t, err)
	expenses, err := service.List()
	require.NoError(t, err)
	assert.Equal(t, 0, len(expenses))
}

func TestServiceDeleteNotFound(t *testing.T) {
	service := expense.NewService()
	_, err := service.Add("Test Expense", 100.0, "A test")
	require.NoError(t, err)
	_, err = service.Delete(999)
	require.Error(t, err)
}
