package annotationimpl

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"main/pkg/affiliate/annotation"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult int64
		entity         *annotation.Annotation
		expectedError  error
	}{
		{
			name:           "create annotation Success",
			expectedResult: 1,
			entity:         &annotation.Annotation{},
			expectedError:  nil,
		},
		{
			name:           "create annotation Error",
			expectedResult: 0,
			entity:         &annotation.Annotation{},
			expectedError:  fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			_, err := store.create(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestCreateAttachement(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult interface{}
		entity         *annotation.AnnotationAttachment
	}{
		{
			name:           "create attachment Success",
			expectedError:  nil,
			expectedResult: 1,
			entity:         &annotation.AnnotationAttachment{},
		},
		{
			name:           "create attackment Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: 0,
			entity:         &annotation.AnnotationAttachment{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			err := store.createAttachment(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestAnnotationTaken(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult interface{}
		entity         *annotation.AnnotationDTO
	}{
		{
			name:           "annotation taken Success",
			expectedError:  nil,
			expectedResult: 1,
			entity:         &annotation.AnnotationDTO{},
		},
		{
			name:           "annotation taken Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: 0,
			entity:         &annotation.AnnotationDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.annotationTaken(context.Background(), 1, "test", 1)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *annotation.Annotation
	}{
		{
			name:          "Success",
			expectedError: nil,
			entity:        &annotation.Annotation{},
		},
		{
			name:          "Error",
			expectedError: fmt.Errorf("error"),
			entity:        &annotation.Annotation{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)

			err := store.update(context.Background(), tc.entity)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		entity        *annotation.Annotation
	}{
		{
			name:          "Success",
			expectedError: nil,
			entity:        &annotation.Annotation{},
		},
		{
			name:          "Error",
			expectedError: fmt.Errorf("error"),
			entity:        &annotation.Annotation{},
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
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestDeleteAttachmentByID(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
		annotationID  int64
		fileName      string
	}{
		{
			name:          "delete attachment by id Success",
			expectedError: nil,
			annotationID:  1,
			fileName:      "test",
		},
		{
			name:          "delete attachment by id Error",
			expectedError: fmt.Errorf("error"),
			annotationID:  1,
			fileName:      "test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError: tc.expectedError,
			}

			store := NewStore(fakeDB)
			err := store.deleteAttachmentByID(context.Background(), tc.annotationID, tc.fileName)
			if err == nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult *annotation.AnnotationDTO
		annotationID   int64
	}{
		{
			name:           "get annotation by attachment Success",
			expectedError:  nil,
			expectedResult: &annotation.AnnotationDTO{},
			annotationID:   1,
		},
		{
			name:           "get annotation by id Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
			annotationID:   1,
		},
		{
			name:           "get annotation by id Error- Not Found",
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
			annotationID:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getByID(context.Background(), tc.annotationID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}

			require.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestGetAttachments(t *testing.T) {
	testCases := []struct {
		name           string
		annotationID   int64
		expectedError  error
		expectedResult []*annotation.AnnotationAttachment
	}{
		{
			name:           "get annotation attachments Success",
			annotationID:   1,
			expectedError:  nil,
			expectedResult: []*annotation.AnnotationAttachment{},
		},
		{
			name:           "get annotation attachments Error",
			annotationID:   2,
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
			result, err := store.getAttachments(context.Background(), tc.annotationID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		query          *annotation.SearchAnnotationQuery
		expectedError  error
		expectedResult *annotation.SearchAnnotationResult
	}{
		{
			name:           "search annotations Success",
			query:          &annotation.SearchAnnotationQuery{},
			expectedError:  nil,
			expectedResult: &annotation.SearchAnnotationResult{},
		},
		{
			name:           "search annotations Error",
			query:          &annotation.SearchAnnotationQuery{},
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
		},
		{
			name:           "search annotations Error- Not Found",
			query:          &annotation.SearchAnnotationQuery{},
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
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
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
			name:          "get count",
			expectedError: nil,
		},
		{
			name:          "get count error",
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
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetAttachmentByAnnotationID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedResult []string
		annotationID   int64
	}{
		{
			name:           "get attachment by annotation id Success",
			expectedError:  nil,
			expectedResult: []string{},
			annotationID:   1,
		},
		{
			name:           "get attachment by annotation id Error",
			expectedError:  fmt.Errorf("error"),
			expectedResult: nil,
			annotationID:   2,
		},
		{
			name:           "get attachment by annotation id Error- Not Found",
			expectedError:  sql.ErrNoRows,
			expectedResult: nil,
			annotationID:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedError:  tc.expectedError,
				ExpectedResult: tc.expectedResult,
			}

			store := NewStore(fakeDB)

			result, err := store.getAttachmentsByAnnotationID(context.Background(), tc.annotationID)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}

			require.Equal(t, tc.expectedResult, result)
		})
	}
}
