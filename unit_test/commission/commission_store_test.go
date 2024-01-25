package commissionimpl

import (
	"context"
	"main/pkg/affiliate/commission"
	"main/pkg/affiliate/commissionenum"
	dbtesting "main/pkg/infra/storage/db/testing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		searchQuery    *commission.SearchCommissionQuery
		commission     []*commission.Commission
		expectedResult *commission.SearchCommissionResult
		expectedError  error
	}{
		{
			name: "Success",
			searchQuery: &commission.SearchCommissionQuery{
				AffiliateID:             1,
				Uuid:                    "uuid",
				LoginName:               "loginName",
				Currency:                "currency",
				CommissionStatus:        commissionenum.CommissionStatus(1),
				PaymentStatus:           commissionenum.CommissionPaymentStatus(1),
				PayoutFrequency:         "payoutFrequency",
				PerPage:                 10,
				Page:                    1,
				HasMinActivePlayerCheck: true,
			},
			commission: []*commission.Commission{},
			expectedResult: &commission.SearchCommissionResult{
				Commissions: []*commission.CommissionDTO{},
			},
			expectedError: nil,
		},
		{
			name:        "Success",
			searchQuery: &commission.SearchCommissionQuery{},
			commission:  []*commission.Commission{},
			expectedResult: &commission.SearchCommissionResult{
				Commissions: []*commission.CommissionDTO{},
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDB := &dbtesting.FakeSqlxdb{
				ExpectedResult: tc.commission,
				ExpectedError:  tc.expectedError,
			}

			store := NewStore(fakeDB)

			result, err := store.Search(context.Background(), tc.searchQuery)
			if err == nil && result != nil && tc.expectedError != nil {
				t.Fatalf("expected error %q, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Fatalf("unexpected error: %v", err)
			}

			require.Equal(t, tc.expectedResult, result)
		})
	}
}
