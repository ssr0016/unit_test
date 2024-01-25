package bannerimpl

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"main/pkg/affiliate/banner"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult int64
		entity         *banner.Banner
		expectedError  error
	}{
		{
			name:           "create banner Success",
			expectedResult: 1,
			entity:         &banner.Banner{},
			expectedError:  nil,
		},
		{
			name:           "create banner Error",
			expectedResult: 0,
			entity:         &banner.Banner{},
			expectedError:  fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.create(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		query          *banner.SearchBannerQuery
		expectedError  error
		expectedResult *banner.SearchBannersResult
	}{
		{
			name:           "search banner Success",
			query:          &banner.SearchBannerQuery{},
			expectedError:  nil,
			expectedResult: &banner.SearchBannersResult{},
		},
		{
			name:           "search banner Error",
			query:          &banner.SearchBannerQuery{},
			expectedError:  nil,
			expectedResult: &banner.SearchBannersResult{},
		},
		{
			name:           "search banner Error- Not Found",
			query:          &banner.SearchBannerQuery{},
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
			_, err := store.search(context.Background(), tc.query)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}

}

func TestGetCount(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "get account",
			expectedError: nil,
		},
		{
			name:          "get account error",
			expectedError: fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)
			_, err := store.getCount(context.Background(), bytes.Buffer{}, []interface{}{})
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult *banner.Banner
		bannerID       int64
	}{
		{
			name:           "get banner by id Success",
			expectedError:  nil,
			expectedResult: &banner.Banner{},
			bannerID:       1,
		},
		{
			name:           "get banner by id Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
			bannerID:       1,
		},
		{
			name:           "get banner by id Error- Not Found",
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
			bannerID:       1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getByID(context.Background(), tc.bannerID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name     string
		expected error
		entity   *banner.Banner
	}{
		{
			name:     "update banner Success",
			expected: nil,
			entity:   &banner.Banner{},
		},
		{
			name:     "update banner Error",
			expected: fmt.Errorf("error"),
			entity:   &banner.Banner{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expected,
			}

			store := NewStore(fakeDB)

			err := store.update(context.Background(), tc.entity)
			if err == nil && tc.expected != nil {
				t.Fatalf("expected error %q, but got none", tc.expected)
			} else if err != nil && tc.expected == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *banner.Banner
	}{
		{
			name:          "update banner status Success",
			expectedError: nil,
			entity:        &banner.Banner{},
		},
		{
			name:          "update banner status Error",
			expectedError: fmt.Errorf("error"),
			entity:        &banner.Banner{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.updateStatus(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("expected no error, but got %q", err)
			}
		})
	}
}
