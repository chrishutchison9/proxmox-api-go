package proxmox

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type lxcTestCaseGet struct {
	name   string
	input  rawConfigLXC
	pool   *PoolName
	state  PowerState
	output *ConfigLXC
}

type lxcTestGetFunc func() []lxcTestCaseGet

func (tests lxcTestGetFunc) Test(t *testing.T) {
	t.Helper()
	for _, test := range tests() {
		t.Run(test.name, func(*testing.T) {
			require.Equal(t, test.output, test.input.Get(test.pool, test.state))
		})
	}
}
