package serviceaccount_test

import (
	"fmt"
	"testing"

	acc "github.com/videocoin/cloud-pkg/api/resources/serviceaccount"

	"github.com/stretchr/testify/require"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		projID, accEmail string
		output           acc.Name
	}{
		{
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   acc.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("proj id: %s, acc email: %s\n", test.projID, test.accEmail), func(t *testing.T) {
			require.Equal(t, test.output, acc.NewName(test.projID, test.accEmail))
		})
	}
}

func TestNewNameWildcard(t *testing.T) {
	tests := []struct {
		accEmail string
		output   acc.Name
	}{
		{
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   acc.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("acc email: %s\n", test.accEmail), func(t *testing.T) {
			require.Equal(t, test.output, acc.NewNameWildcard(test.accEmail))
		})
	}
}

func TestNewEmail(t *testing.T) {
	tests := []struct {
		projID, accID string
		output        string
	}{
		{
			projID: "videocoin-123",
			accID:  "account1",
			output: "account1@videocoin-123.vserviceaccount.com",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("proj id: %s, acc id: %s\n", test.projID, test.accID), func(t *testing.T) {
			require.Equal(t, test.output, acc.NewEmail(test.projID, test.accID))
		})
	}
}

func TestNameEmail(t *testing.T) {
	tests := []struct {
		accName acc.Name
		output  string
	}{
		{
			accName: acc.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			output:  "account1@videocoin-123.vserviceaccount.com",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("acc name: %s\n", string(test.accName)), func(t *testing.T) {
			require.Equal(t, test.output, test.accName.Email())
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name   string
		accID  string
		output bool
	}{
		{
			name:   "invalid id: must have at least 6 chars",
			accID:  "accou",
			output: false,
		},
		{
			name:   "invalid id: must have less than 30 chars",
			accID:  "account1sadsfafsdfasdfadsfasdfs",
			output: false,
		},
		{
			name:   "valid id",
			accID:  "account1",
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, acc.IsValidID(test.accID))
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		accEmail string
		output   bool
	}{
		{
			name:     "invalid email: incorrect domain",
			accEmail: "account1@videocoin-123.gserviceaccount.com",
			output:   false,
		},
		{
			name:     "valid email",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, acc.IsValidEmail(test.accEmail))
		})
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		name    string
		accName string
		output  acc.Name
		err     error
	}{
		{
			name:    "invalid name: incorrect collection id",
			accName: "projects/videocoin-123/serviceAccount/account1@videocoin-123.vserviceaccount.com",
			output:  "",
			err:     acc.ErrInvalidName,
		},
		{
			name:    "invalid name: incorrect domain",
			accName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.gserviceaccount.com",
			output:  "",
			err:     acc.ErrInvalidName,
		},
		{
			name:    "invalid name: missing @",
			accName: "projects/videocoin-123/serviceAccounts/account1videocoin-123.vserviceaccount.com",
			output:  "",
			err:     acc.ErrInvalidName,
		},
		{
			name:    "valid name wildcard",
			accName: "projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output:  acc.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			err:     nil,
		},
		{
			name:    "valid name",
			accName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output:  acc.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			err:     nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := acc.ParseName(test.accName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}
