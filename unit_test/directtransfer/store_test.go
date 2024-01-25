package directtransferimpl

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	dt "main/pkg/affiliate/directtransfer"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		query          *dt.SearchTransfersQuery
		expectedError  error
		expectedResult *dt.SearchTransfersResult
	}{
		{
			name:           "search transfers Success",
			query:          &dt.SearchTransfersQuery{},
			expectedError:  nil,
			expectedResult: &dt.SearchTransfersResult{},
		},
		{
			name:           "search transfers Error",
			query:          &dt.SearchTransfersQuery{},
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
		{
			name:           "search transfers Error- Not Found",
			query:          &dt.SearchTransfersQuery{},
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.Search(context.Background(), tc.query)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetCount(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		sql            bytes.Buffer
		whereParams    []any
		expectedResult int64
	}{
		{
			name:           "get count Success",
			expectedError:  nil,
			sql:            bytes.Buffer{},
			whereParams:    []any{},
			expectedResult: 1,
		},
		{
			name:           "get count Error",
			expectedError:  fmt.Errorf("error"),
			sql:            bytes.Buffer{},
			whereParams:    []any{},
			expectedResult: 0,
		},
		{
			name:           "get count Error- Not Found",
			expectedError:  sql.ErrNoRows,
			sql:            bytes.Buffer{},
			whereParams:    []any{},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getCount(context.Background(), tc.sql, tc.whereParams)
			if err != nil && result == 0 && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetTransferGrandTotal(t *testing.T) {
	testCases := []struct {
		name           string
		whereCondtions []string
		whereParams    []any
		expectedError  error
		expectedResult *dt.TransferGrandTotalResult
	}{
		{
			name:           "get transfer grand total Success",
			whereCondtions: []string{},
			whereParams:    []any{},
			expectedError:  nil,
			expectedResult: &dt.TransferGrandTotalResult{},
		},
		{
			name:           "get transfer grand total Error",
			whereCondtions: []string{},
			whereParams:    []any{},
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getTransferGrandTotal(context.Background(), tc.whereCondtions, tc.whereParams)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetLogByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		transferID     int64
		expectedResult []*dt.TransferLog
	}{
		{
			name:           "get log by id Success",
			transferID:     1,
			expectedError:  nil,
			expectedResult: []*dt.TransferLog{},
		},
		{
			name:           "get log by id Error",
			transferID:     1,
			expectedError:  fmt.Errorf("error"),
			expectedResult: []*dt.TransferLog{},
		},
		{
			name:           "get log by id Error- Not Found",
			transferID:     1,
			expectedError:  sql.ErrNoRows,
			expectedResult: []*dt.TransferLog{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getLogByID(context.Background(), tc.transferID)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		name           string
		id             int64
		expectedError  error
		expectedResult *dt.TransferDTO
	}{
		{
			name:           "get by id Success",
			id:             1,
			expectedError:  nil,
			expectedResult: &dt.TransferDTO{},
		},
		{
			name:           "get by id Error",
			id:             1,
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getByID(context.Background(), tc.id)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetByTransactionID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		transactionID  string
		expectedResult *dt.TransferDTO
	}{
		{
			name:           "get by transaction id Success",
			transactionID:  "1",
			expectedError:  nil,
			expectedResult: &dt.TransferDTO{},
		},
		{
			name:           "get by transaction id Error",
			transactionID:  "1",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getByTransactionID(context.Background(), tc.transactionID)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetStatusStats(t *testing.T) {
	testCases := []struct {
		name           string
		query          *dt.GetStatustatsQuery
		expectedError  error
		expectedResult *dt.GetStatusStatsResult
	}{
		{
			name:           "get status stats Success",
			query:          &dt.GetStatustatsQuery{},
			expectedError:  nil,
			expectedResult: &dt.GetStatusStatsResult{},
		},
		{
			name:           "get status stats Error",
			query:          &dt.GetStatustatsQuery{},
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
		{
			name:           "get status stats Error- Not Found",
			query:          &dt.GetStatustatsQuery{},
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.GetStatusStats(context.Background(), tc.query)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		entity         *dt.Transfer
		expectedResult int64
	}{
		{
			name:           "create success",
			expectedError:  nil,
			entity:         &dt.Transfer{},
			expectedResult: 1,
		},
		{
			name:           "create error",
			expectedError:  fmt.Errorf("error"),
			entity:         &dt.Transfer{},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.create(context.Background(), tc.entity)
			if err != nil && result == 0 && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *dt.Transfer
	}{
		{
			name:          "update status success",
			expectedError: nil,
			entity:        &dt.Transfer{},
		},
		{
			name:          "update status error",
			expectedError: fmt.Errorf("error"),
			entity:        &dt.Transfer{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.updateStatus(context.Background(), tc.entity)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestCreateLog(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *dt.TransferLog
	}{
		{
			name:          "create log success",
			expectedError: nil,
			entity:        &dt.TransferLog{},
		},
		{
			name:          "create log error",
			expectedError: fmt.Errorf("error"),
			entity:        &dt.TransferLog{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.createLog(context.Background(), tc.entity)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}
