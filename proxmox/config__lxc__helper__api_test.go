package proxmox

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/require"
)

type lxcTestsApiFunc func() lxcTestsAPI

func (tests lxcTestsApiFunc) Test(t *testing.T) {
	t.Helper()
	test := tests.format()
	refrence := tests.format()
	for i := range test.create {
		t.Run("create/"+test.create[i].name, func(*testing.T) {
			config := test.create[i].config
			_, output, pool := config.mapToApiCreate()
			clone := maps.Clone(test.create[i].body)
			if !(test.create[i].omitDefaults == lxcDefaultsAll || test.create[i].omitDefaults == lxcDefaultsCreate) {
				if clone == nil {
					clone = map[string]string{}
				}
				if _, isSet := clone["unprivileged"]; !isSet {
					clone["unprivileged"] = "1" // Default to unprivileged
				}
			}
			testParamsEqualRaw(t, clone, output)
			require.Equal(t, test.create[i].pool, pool)
			require.Equal(t, refrence.create[i].config, config, "mutated input config")
		})
	}
	for i := range test.update {
		t.Run("update/"+test.update[i].name, func(*testing.T) {
			config := test.update[i].config
			_, output := config.mapToApiUpdate(test.update[i].currentConfig)
			clone := maps.Clone(test.update[i].body)
			if !(test.update[i].omitDefaults == lxcDefaultsAll || test.update[i].omitDefaults == lxcDefaultsUpdate) {
				if clone == nil {
					clone = map[string]string{}
				}
				if _, isSet := clone["digest"]; !isSet {
					clone["digest"] = "" // set empty digest
				}
			}
			testParamsEqualRaw(t, test.update[i].body, output)
			require.Equal(t, refrence.update[i].config, config, "mutated input config")
		})
	}
}

func (tests lxcTestsApiFunc) format() struct {
	create []lxcTestCaseAPI
	update []lxcTestCaseAPI
} {
	data := tests()
	return struct {
		create []lxcTestCaseAPI
		update []lxcTestCaseAPI
	}{
		create: append(data.create, data.createUpdate...),
		update: append(data.update, data.createUpdate...),
	}
}
