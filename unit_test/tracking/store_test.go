package trackingimpl

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"main/pkg/affiliate/tracking"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *tracking.Tracking
	}{
		{
			name:          "create tracking success",
			expectedError: nil,
			entity:        &tracking.Tracking{},
		},
		{
			name:          "create tracking error",
			expectedError: errors.New("test error"),
			entity:        &tracking.Tracking{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.Create(context.Background(), tc.entity)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestTrackingTaken(t *testing.T) {
	testCases := []struct {
		name           string
		affiliateID    int64
		nameTracking   string
		trackingID     int64
		expectedError  error
		expectedResult []*tracking.TrackingDTO
	}{
		{
			name:           "tracking taken success",
			affiliateID:    1,
			nameTracking:   "test",
			trackingID:     1,
			expectedError:  nil,
			expectedResult: []*tracking.TrackingDTO{},
		},
		{
			name:           "tracking taken error",
			affiliateID:    1,
			nameTracking:   "test",
			trackingID:     1,
			expectedError:  errors.New("test error"),
			expectedResult: []*tracking.TrackingDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.TrackingTaken(context.Background(), tc.affiliateID, tc.nameTracking, tc.trackingID)
			if err != nil && result == nil && tc.expectedError == nil {
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
		entity        *tracking.TrackingLog
	}{
		{
			name:          "create log success",
			expectedError: nil,
			entity:        &tracking.TrackingLog{},
		},
		{
			name:          "create log error",
			expectedError: errors.New("test error"),
			entity:        &tracking.TrackingLog{},
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

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		query          *tracking.SearchTrackingQuery
		expectedError  error
		expectedResult *tracking.SearchTrackingResult
	}{
		{
			name:           "search tracking success",
			query:          &tracking.SearchTrackingQuery{},
			expectedError:  nil,
			expectedResult: &tracking.SearchTrackingResult{},
		},
		{
			name:           "search tracking error",
			query:          &tracking.SearchTrackingQuery{},
			expectedError:  errors.New("test error"),
			expectedResult: nil,
		},
		{
			name:           "search tracking error - not found",
			query:          &tracking.SearchTrackingQuery{},
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

			result, err := store.search(context.Background(), tc.query)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetTrackingCount(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		sql            bytes.Buffer
		whereParams    []interface{}
		expectedResult int64
	}{
		{
			name:           "get tracking count success",
			expectedError:  nil,
			sql:            bytes.Buffer{},
			whereParams:    []interface{}{},
			expectedResult: 1,
		},
		{
			name:           "get tracking count error",
			expectedError:  errors.New("test error"),
			sql:            bytes.Buffer{},
			whereParams:    []interface{}{},
			expectedResult: 0,
		},
		{
			name:           "get tracking count error - not found",
			expectedError:  sql.ErrNoRows,
			sql:            bytes.Buffer{},
			whereParams:    []interface{}{},
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

			_, err := store.getTrackingCount(context.Background(), tc.sql, tc.whereParams)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetTrackingByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		trackingID     int64
		expectedResult *tracking.TrackingDTO
	}{
		{
			name:           "get tracking by id success",
			expectedError:  nil,
			trackingID:     1,
			expectedResult: &tracking.TrackingDTO{},
		},
		{
			name:           "get tracking by id error",
			expectedError:  errors.New("test error"),
			trackingID:     1,
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

			result, err := store.getTrackingByID(context.Background(), tc.trackingID)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetSummaryTrackingLog(t *testing.T) {
	testCases := []struct {
		name           string
		query          *tracking.SearchSummaryTrackingLogQuery
		expectedError  error
		expectedResult *tracking.SummaryTrackingLogResult
	}{
		{
			name:           "get summary tracking log success",
			query:          &tracking.SearchSummaryTrackingLogQuery{},
			expectedError:  nil,
			expectedResult: &tracking.SummaryTrackingLogResult{},
		},
		{
			name:           "get summary tracking log error",
			query:          &tracking.SearchSummaryTrackingLogQuery{},
			expectedError:  errors.New("test error"),
			expectedResult: nil,
		},
		{
			name:           "get summary tracking log error - not found",
			query:          &tracking.SearchSummaryTrackingLogQuery{},
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

			result, err := store.getSummaryTrackingLog(context.Background(), tc.query)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestUpdateTracking(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *tracking.Tracking
	}{
		{
			name:          "update tracking success",
			expectedError: nil,
			entity:        &tracking.Tracking{},
		},
		{
			name:          "update tracking error",
			expectedError: errors.New("test error"),
			entity:        &tracking.Tracking{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.updateTracking(context.Background(), tc.entity)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetSummaryTrackingLogForAffiliate(t *testing.T) {
	testCases := []struct {
		name           string
		query          *tracking.GetSummaryTrackingQuery
		expectedError  error
		expectedResult *tracking.GetSummaryTrackingResult
	}{
		{
			name:           "get summary tracking log for affiliate success",
			query:          &tracking.GetSummaryTrackingQuery{},
			expectedError:  nil,
			expectedResult: &tracking.GetSummaryTrackingResult{},
		},
		{
			name:           "get summary tracking log for affiliate error",
			query:          &tracking.GetSummaryTrackingQuery{},
			expectedError:  errors.New("test error"),
			expectedResult: nil,
		},
		{
			name:           "get summary tracking log for affiliate error - not found",
			query:          &tracking.GetSummaryTrackingQuery{},
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
		},
		{
			name:           "get summary tracking log for affiliate error - not found",
			query:          &tracking.GetSummaryTrackingQuery{},
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

			result, err := store.getSummaryTrackingLogForAffiliate(context.Background(), tc.query)
			if err != nil && result == nil && tc.expectedError == nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}
