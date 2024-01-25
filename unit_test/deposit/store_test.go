package depositimpl

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"main/pkg/affiliate/deposit"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"
)

func TestGetDepositLog(t *testing.T) {
	testCases := []struct {
		name           string
		depositID      int64
		expectedError  error
		expectedResult []*deposit.DepositLog
	}{
		{
			name:           "get deposit log Success",
			depositID:      1,
			expectedError:  nil,
			expectedResult: []*deposit.DepositLog{},
		},
		{
			name:           "get deposit log Error",
			depositID:      2,
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

			result, err := store.getDepositLog(context.Background(), tc.depositID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetDepositByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult *deposit.DepositDTO
		depositID      int64
	}{
		{
			name:           "get deposit by id Success",
			expectedError:  nil,
			expectedResult: &deposit.DepositDTO{},
			depositID:      1,
		},
		{
			name:           "get deposit by id Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
			depositID:      1,
		},
		{
			name:           "get deposit by id Error- Not Found",
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
			depositID:      1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getDepositByID(context.Background(), tc.depositID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetDepositByTransactionID(t *testing.T) {
	testCases := []struct {
		name           string
		transactionID  string
		expectedError  error
		expectedResult *deposit.DepositDTO
	}{
		{
			name:           "get deposit by transaction id Success",
			transactionID:  "1",
			expectedError:  nil,
			expectedResult: &deposit.DepositDTO{},
		},
		{
			name:           "get deposit by transaction id Error",
			transactionID:  "2",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
		{
			name:           "get deposit by transaction id Error- Not Found",
			transactionID:  "3",
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

			result, err := store.getDepositByTransactionID(context.Background(), tc.transactionID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestSearchDeposit(t *testing.T) {
	testCases := []struct {
		name           string
		query          *deposit.SearchDepositQuery
		expectedError  error
		expectedResult *deposit.SearchDepositQueryResult
	}{
		{
			name:           "search deposit Success",
			query:          &deposit.SearchDepositQuery{},
			expectedError:  nil,
			expectedResult: &deposit.SearchDepositQueryResult{},
		},
		{
			name:           "search deposit Error",
			query:          &deposit.SearchDepositQuery{},
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
		{
			name:           "search deposit Error- Not Found",
			query:          &deposit.SearchDepositQuery{},
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

			result, err := store.SearchDeposit(context.Background(), tc.query)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetSearchDepositCount(t *testing.T) {
	testCases := []struct {
		name          string
		sql           bytes.Buffer
		expectedError error
	}{
		{
			name:          "get search deposit count Success",
			sql:           bytes.Buffer{},
			expectedError: nil,
		},
		{
			name:          "get search deposit count Error",
			sql:           bytes.Buffer{},
			expectedError: fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)
			_, err := store.getSearchDepositCount(context.Background(), tc.sql, []interface{}{})
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestCreateDeposit(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *deposit.Deposit
	}{
		{
			name:          "create deposit Success",
			expectedError: nil,
			entity:        &deposit.Deposit{},
		},
		{
			name:          "create deposit Error",
			expectedError: fmt.Errorf("error"),
			entity:        &deposit.Deposit{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			_, err := store.CreateDeposit(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestSaveDepositLog(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *deposit.DepositLog
	}{
		{
			name:          "save deposit log Success",
			expectedError: nil,
			entity:        &deposit.DepositLog{},
		},
		{
			name:          "save deposit log Error",
			expectedError: fmt.Errorf("error"),
			entity:        &deposit.DepositLog{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.SaveDepositLog(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetDepositByStatus(t *testing.T) {
	testCases := []struct {
		name           string
		id             int64
		expectedError  error
		status         int
		expectedResult *deposit.DepositDTO
	}{
		{
			name:           "get deposit by status Success",
			id:             1,
			expectedError:  nil,
			status:         1,
			expectedResult: &deposit.DepositDTO{},
		},
		{
			name:           "get deposit by status Error",
			id:             2,
			status:         1,
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

			result, err := store.GetDepositByStatus(context.Background(), tc.id, 1)
			if err == nil && result == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestUpdateDepositAmount(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *deposit.Deposit
	}{
		{
			name:          "update deposit amount Success",
			expectedError: nil,
			entity:        &deposit.Deposit{},
		},
		{
			name:          "update deposit amount Error",
			expectedError: fmt.Errorf("error"),
			entity:        &deposit.Deposit{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.UpdateDepositAmount(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestUpdateDepositAttachments(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		cmd           *deposit.UpdateDepositAttachmentsCommand
	}{
		{
			name:          "update deposit attachments Success",
			expectedError: nil,
			cmd:           &deposit.UpdateDepositAttachmentsCommand{},
		},
		{
			name:          "update deposit attachments Error",
			expectedError: fmt.Errorf("error"),
			cmd:           &deposit.UpdateDepositAttachmentsCommand{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.UpdateDepositAttachments(context.Background(), tc.cmd)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestUpdateDepositPartnerTransactionID(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		cmd           *deposit.UpdateDepositPartnerTransactionIDCommand
	}{
		{
			name:          "update deposit partner transaction id Success",
			expectedError: nil,
			cmd:           &deposit.UpdateDepositPartnerTransactionIDCommand{},
		},
		{
			name:          "update deposit partner transaction id Error",
			expectedError: fmt.Errorf("error"),
			cmd:           &deposit.UpdateDepositPartnerTransactionIDCommand{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.UpdateDepositPartnerTransactionID(context.Background(), tc.cmd)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetTotalByStatus(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		to             string
		status         int
		expectedResult []*deposit.Stats
	}{
		{
			name:           "get total by status Success",
			expectedError:  nil,
			to:             "2020-01-01 00:00:00",
			status:         1,
			expectedResult: []*deposit.Stats{},
		},
		{
			name:           "get total by status Error",
			expectedError:  fmt.Errorf("error"),
			to:             "2020-01-01 00:00:00",
			status:         1,
			expectedResult: []*deposit.Stats{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getTotalByStatus(context.Background(), "2020-01-01 00:00:00", tc.to, deposit.Status(tc.status))
			if err == nil && result == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestStatusStats(t *testing.T) {
	testCases := []struct {
		name           string
		query          *deposit.GetStatusStatsQuery
		expectedError  error
		expectedResult *deposit.GetStatsChartResult
	}{
		{
			name:           "get status stats Success",
			query:          &deposit.GetStatusStatsQuery{},
			expectedError:  nil,
			expectedResult: &deposit.GetStatsChartResult{},
		},
		{
			name:           "get status stats Error",
			query:          &deposit.GetStatusStatsQuery{},
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

			result, err := store.GetStatusStats(context.Background(), tc.query)
			if err == nil && result == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}


