package helm_test

import (
	"testing"

	"github.com/redpanda-data/helm-charts/pkg/helm"
	"github.com/stretchr/testify/require"
)

type Flags struct {
	NoWait        bool `flag:"wait"`
	NoWaitForJobs bool `flag:"no-wait-for-jobs"`
	NotAFlag      string
	StringFlag    string `flag:"string-flag"`
}

func TestToFlags(t *testing.T) {
	testCases := []struct {
		in  Flags
		out []string
	}{
		{
			in: Flags{},
			out: []string{
				"--wait=true",
				"--no-wait-for-jobs=false",
			},
		},
		{
			in: Flags{},
			out: []string{
				"--wait=true",
				"--no-wait-for-jobs=false",
			},
		},
		{
			in: Flags{
				StringFlag: "something",
			},
			out: []string{
				"--wait=true",
				"--no-wait-for-jobs=false",
				"--string-flag=something",
			},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.out, helm.ToFlags(tc.in))
	}
}
