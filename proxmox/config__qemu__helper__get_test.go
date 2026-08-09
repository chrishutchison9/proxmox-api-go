package proxmox

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type qemuTestCaseGet struct {
	name   string
	input  map[string]any
	vmr    VmRef
	output *ConfigQemu
	err    error
}

type qemuTestGetFunc func() []qemuTestCaseGet

func (tests qemuTestGetFunc) Test(t *testing.T) {
	t.Helper()
	for _, test := range tests() {
		t.Run(test.name, func(*testing.T) {
			output, err := (&rawConfigQemu{a: test.input}).get(test.vmr)
			if err != nil {
				require.Equal(t, test.err, err)
			} else {
				require.Equal(t, test.output, output)
			}
		})
	}
}
