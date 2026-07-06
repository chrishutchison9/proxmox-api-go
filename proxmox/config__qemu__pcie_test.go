package proxmox

import (
	"errors"
	"testing"

	"github.com/Telmate/proxmox-api-go/internal/util"
	"github.com/stretchr/testify/require"
)

func Test_ConfigQemu_PciDevices_MapToApi(t *testing.T) {
	t.Parallel()
	tests := qemuTestsApiFunc(func() qemuTestsAPI {
		return qemuTestsAPI{
			create: []qemuTestCaseAPI{
				{name: `Delete`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Delete: true}}}}},
			createUpdate: []qemuTestCaseAPI{
				{name: `Mapping.DeviceID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID15: QemuPci{Mapping: &QemuPciMapping{
							DeviceID: new(PciDeviceID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID15: QemuPci{Mapping: &QemuPciMapping{
							DeviceID: new(PciDeviceID("0x8000"))}}}},
					body: map[string]string{"hostpci15": "mapping%3D%2Crombar%3D0%2Cdevice-id%3D0x8086"}}, // "mapping=,rombar=0,device-id=0x8086"
				{name: `Mapping.ID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							ID: new(ResourceMappingPciID("aaaaa"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							ID: new(ResourceMappingPciID("bbbbb"))}}}},
					body: map[string]string{"hostpci14": "mapping%3Daaaaa%2Crombar%3D0"}}, // "mapping=aaaaa,rombar=0"
				{name: `MApping.MDev`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							MDev: new(PciMediatedDevice("vendor-665"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							MDev: new(PciMediatedDevice(PciMediatedDevice("vendor-000")))}}}},
					body: map[string]string{"hostpci14": "mapping%3D%2Crombar%3D0%2Cmdev%3Dvendor-665"}}, // "mapping=,rombar=0,mdev=vendor-665"
				{name: `Mapping.Pci`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID13: QemuPci{Mapping: &QemuPciMapping{
							PCIe: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID13: QemuPci{Mapping: &QemuPciMapping{
							PCIe: new(false)}}}},
					body: map[string]string{"hostpci13": "mapping%3D%2Cpcie%3D1%2Crombar%3D0"}}, // "mapping=,pcie=1,rombar=0"
				{name: `Mapping.PrimaryGPU`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID12: QemuPci{Mapping: &QemuPciMapping{
							PrimaryGPU: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID12: QemuPci{Mapping: &QemuPciMapping{
							PrimaryGPU: new(false)}}}},
					body: map[string]string{"hostpci12": "mapping%3D%2Cx-vga%3D1%2Crombar%3D0"}}, // "mapping=,x-vga=1,rombar=0"
				{name: `Mapping.ROMbar`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID11: QemuPci{Mapping: &QemuPciMapping{
							ROMbar: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID11: QemuPci{Mapping: &QemuPciMapping{
							ROMbar: new(false)}}}},
					body: map[string]string{"hostpci11": "mapping%3D"}}, // "mapping="
				{name: `Mapping.SubDeviceID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID10: QemuPci{Mapping: &QemuPciMapping{
							SubDeviceID: new(PciSubDeviceID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID10: QemuPci{Mapping: &QemuPciMapping{
							SubDeviceID: new(PciSubDeviceID("0x8000"))}}}},
					body: map[string]string{"hostpci10": "mapping%3D%2Crombar%3D0%2Csub-device-id%3D0x8086"}}, // "mapping=,rombar=0,sub-device-id=0x8086"
				{name: `Mapping.SubVendorID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID9: QemuPci{Mapping: &QemuPciMapping{
							SubVendorID: new(PciSubVendorID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID9: QemuPci{Mapping: &QemuPciMapping{
							SubVendorID: new(PciSubVendorID("0x8000"))}}}},
					body: map[string]string{"hostpci9": "mapping%3D%2Crombar%3D0%2Csub-vendor-id%3D0x8086"}}, // "mapping=,rombar=0,sub-vendor-id=0x8086"
				{name: `Mapping.VendorID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID8: QemuPci{Mapping: &QemuPciMapping{
							VendorID: new(PciVendorID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID8: QemuPci{Mapping: &QemuPciMapping{
							VendorID: new(PciVendorID("0x8000"))}}}},
					body: map[string]string{"hostpci8": "mapping%3D%2Crombar%3D0%2Cvendor-id%3D0x8086"}}, // "mapping=,rombar=0,vendor-id=0x8086"
				{name: `Raw.DeviceID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID0: QemuPci{Raw: &QemuPciRaw{
							DeviceID: new(PciDeviceID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID0: QemuPci{Raw: &QemuPciRaw{
							DeviceID: new(PciDeviceID("0x8000"))}}}},
					body: map[string]string{"hostpci0": "%2Crombar%3D0%2Cdevice-id%3D0x8086"}}, // ",rombar=0,device-id=0x8086"
				{name: `Raw.ID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID1: QemuPci{Raw: &QemuPciRaw{
							ID: new(PciID("0000:00:00.0"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID1: QemuPci{Raw: &QemuPciRaw{
							ID: new(PciID("0000:00:00.1"))}}}},
					body: map[string]string{"hostpci1": "0000%3A00%3A00.0%2Crombar%3D0"}}, // "0000:00:00.0,rombar=0"
				{name: `Raw.MDev`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							MDev: new(PciMediatedDevice("vendor-665"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							MDev: new(PciMediatedDevice("vendor-000"))}}}},
					body: map[string]string{"hostpci2": "%2Crombar%3D0%2Cmdev%3Dvendor-665"}}, // ",rombar=0,mdev=vendor-665"
				{name: `Raw.Pci`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							PCIe: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							PCIe: new(false)}}}},
					body: map[string]string{"hostpci2": "%2Cpcie%3D1%2Crombar%3D0"}}, // ",pcie=1,rombar=0"
				{name: `Raw.PrimaryGPU`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID3: QemuPci{Raw: &QemuPciRaw{
							PrimaryGPU: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID3: QemuPci{Raw: &QemuPciRaw{
							PrimaryGPU: new(false)}}}},
					body: map[string]string{"hostpci3": "%2Cx-vga%3D1%2Crombar%3D0"}}, // ",x-vga=1,rombar=0"
				{name: `Raw.ROMbar`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID4: QemuPci{Raw: &QemuPciRaw{
							ROMbar: new(true)}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID4: QemuPci{Raw: &QemuPciRaw{
							ROMbar: new(false)}}}},
					body: map[string]string{"hostpci4": ""}},
				{name: `Raw.SubDeviceID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Raw: &QemuPciRaw{
							SubDeviceID: new(PciSubDeviceID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Raw: &QemuPciRaw{
							SubDeviceID: new(PciSubDeviceID("0x8000"))}}}},
					body: map[string]string{"hostpci5": "%2Crombar%3D0%2Csub-device-id%3D0x8086"}}, // ",rombar=0,sub-device-id=0x8086"
				{name: `Raw.SubVendorID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID6: QemuPci{Raw: &QemuPciRaw{
							SubVendorID: new(PciSubVendorID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID6: QemuPci{Raw: &QemuPciRaw{
							SubVendorID: new(PciSubVendorID("0x8000"))}}}},
					body: map[string]string{"hostpci6": "%2Crombar%3D0%2Csub-vendor-id%3D0x8086"}}, // ",rombar=0,sub-vendor-id=0x8086"
				{name: `Raw.VendorID`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID7: QemuPci{Raw: &QemuPciRaw{
							VendorID: new(PciVendorID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID7: QemuPci{Raw: &QemuPciRaw{
							VendorID: new(PciVendorID("0x8000"))}}}},
					body: map[string]string{"hostpci7": "%2Crombar%3D0%2Cvendor-id%3D0x8086"}}}, // ",rombar=0,vendor-id=0x8086"
			update: []qemuTestCaseAPI{
				{name: `Delete`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Delete: true}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{}}},
					body: map[string]string{"delete": "hostpci5"}},
				{name: `Delete create no effect`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Delete: true}}}},
				{name: `Delete no effect`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Delete: true}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID7: QemuPci{}}}},
				{name: `Mapping.DeviceID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID15: QemuPci{Mapping: &QemuPciMapping{
							DeviceID: new(PciDeviceID("8086"))}}}},
					body: map[string]string{"hostpci15": "mapping%3D%2Crombar%3D0%2Cdevice-id%3D0x8086"}}, // "mapping=,rombar=0,device-id=0x8086"
				{name: `Mapping.ID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							ID: new(ResourceMappingPciID("aaaaa"))}}}},
					body: map[string]string{"hostpci14": "mapping%3Daaaaa%2Crombar%3D0"}}, // "mapping=aaaaa,rombar=0"
				{name: `MApping.MDev create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID14: QemuPci{Mapping: &QemuPciMapping{
							MDev: new(PciMediatedDevice("vendor-665"))}}}},
					body: map[string]string{"hostpci14": "mapping%3D%2Crombar%3D0%2Cmdev%3Dvendor-665"}}, // "mapping=,rombar=0,mdev=vendor-665"
				{name: `Mapping.Pci create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID13: QemuPci{Mapping: &QemuPciMapping{
							PCIe: new(true)}}}},
					body: map[string]string{"hostpci13": "mapping%3D%2Cpcie%3D1%2Crombar%3D0"}}, // "mapping=,pcie=1,rombar=0"
				{name: `Mapping.PrimaryGPU create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID12: QemuPci{Mapping: &QemuPciMapping{
							PrimaryGPU: new(true)}}}},
					body: map[string]string{"hostpci12": "mapping%3D%2Cx-vga%3D1%2Crombar%3D0"}}, // "mapping=,x-vga=1,rombar=0"
				{name: `Mapping.ROMbar create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID11: QemuPci{Mapping: &QemuPciMapping{
							ROMbar: new(true)}}}},
					body: map[string]string{"hostpci11": "mapping%3D"}}, // "mapping="
				{name: `Mapping.SubDeviceID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID10: QemuPci{Mapping: &QemuPciMapping{
							SubDeviceID: new(PciSubDeviceID("8086"))}}}},
					body: map[string]string{"hostpci10": "mapping%3D%2Crombar%3D0%2Csub-device-id%3D0x8086"}}, // "mapping=,rombar=0,sub-device-id=0x8086"
				{name: `Mapping.SubVendorID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID9: QemuPci{Mapping: &QemuPciMapping{
							SubVendorID: new(PciSubVendorID("8086"))}}}},
					body: map[string]string{"hostpci9": "mapping%3D%2Crombar%3D0%2Csub-vendor-id%3D0x8086"}}, // "mapping=,rombar=0,sub-vendor-id=0x8086"
				{name: `Mapping.VendorID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID8: QemuPci{Mapping: &QemuPciMapping{
							VendorID: new(PciVendorID("8086"))}}}},
					body: map[string]string{"hostpci8": "mapping%3D%2Crombar%3D0%2Cvendor-id%3D0x8086"}}, // "mapping=,rombar=0,vendor-id=0x8086"
				{name: `Raw.DeviceID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID0: QemuPci{Raw: &QemuPciRaw{
							DeviceID: new(PciDeviceID("8086"))}}}},
					body: map[string]string{"hostpci0": "%2Crombar%3D0%2Cdevice-id%3D0x8086"}}, // ",rombar=0,device-id=0x8086"
				{name: `Raw.ID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID1: QemuPci{Raw: &QemuPciRaw{
							ID: new(PciID("0000:00:00.0"))}}}},
					body: map[string]string{"hostpci1": "0000%3A00%3A00.0%2Crombar%3D0"}}, // "0000:00:00.0,rombar=0"
				{name: `Raw.MDev create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							MDev: new(PciMediatedDevice("vendor-665"))}}}},
					body: map[string]string{"hostpci2": "%2Crombar%3D0%2Cmdev%3Dvendor-665"}}, // ",rombar=0,mdev=vendor-665"
				{name: `Raw.Pci create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID2: QemuPci{Raw: &QemuPciRaw{
							PCIe: new(true)}}}},
					body: map[string]string{"hostpci2": "%2Cpcie%3D1%2Crombar%3D0"}}, // ",pcie=1,rombar=0"
				{name: `Raw.PrimaryGPU create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID3: QemuPci{Raw: &QemuPciRaw{
							PrimaryGPU: new(true)}}}},
					body: map[string]string{"hostpci3": "%2Cx-vga%3D1%2Crombar%3D0"}}, // ",x-vga=1,rombar=0"
				{name: `Raw.ROMbar create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID4: QemuPci{Raw: &QemuPciRaw{
							ROMbar: new(true)}}}},
					body: map[string]string{"hostpci4": ""}},
				{name: `Raw.SubDeviceID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID5: QemuPci{Raw: &QemuPciRaw{
							SubDeviceID: new(PciSubDeviceID("8086"))}}}},
					body: map[string]string{"hostpci5": "%2Crombar%3D0%2Csub-device-id%3D0x8086"}}, // ",rombar=0,sub-device-id=0x8086"
				{name: `Raw.SubVendorID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID6: QemuPci{Raw: &QemuPciRaw{
							SubVendorID: new(PciSubVendorID("8086"))}}}},
					body: map[string]string{"hostpci6": "%2Crombar%3D0%2Csub-vendor-id%3D0x8086"}}, // ",rombar=0,sub-vendor-id=0x8086"
				{name: `Raw.VendorID create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID7: QemuPci{Raw: &QemuPciRaw{
							VendorID: new(PciVendorID("8086"))}}}},
					body: map[string]string{"hostpci7": "%2Crombar%3D0%2Cvendor-id%3D0x8086"}},
				{name: `Raw.VendorID update create`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID7: QemuPci{Raw: &QemuPciRaw{
							VendorID: new(PciVendorID("8086"))}}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{QemuPciID10: QemuPci{}}},
					body:          map[string]string{"hostpci7": "%2Crombar%3D0%2Cvendor-id%3D0x8086"}},
				{name: `same`,
					config: &ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID4: QemuPci{
							Mapping: &QemuPciMapping{
								DeviceID:    new(PciDeviceID("8086")),
								ID:          new(ResourceMappingPciID("aaaaa")),
								MDev:        new(PciMediatedDevice("vendor-665")),
								PCIe:        new(false),
								PrimaryGPU:  new(true),
								ROMbar:      new(false),
								SubDeviceID: new(PciSubDeviceID("4522")),
								SubVendorID: new(PciSubVendorID("74526")),
								VendorID:    new(PciVendorID("2321"))},
						}}},
					currentLegacy: ConfigQemu{PciDevices: QemuPciDevices{
						QemuPciID4: QemuPci{
							Mapping: &QemuPciMapping{
								DeviceID:    new(PciDeviceID("8086")),
								ID:          new(ResourceMappingPciID("aaaaa")),
								MDev:        new(PciMediatedDevice("vendor-665")),
								PCIe:        new(false),
								PrimaryGPU:  new(true),
								ROMbar:      new(false),
								SubDeviceID: new(PciSubDeviceID("4522")),
								SubVendorID: new(PciSubVendorID("74526")),
								VendorID:    new(PciVendorID("2321"))},
						}}},
				},
			}}
	})
	tests.Test(t)
}

func Test_QemuPciDevices_Validate(t *testing.T) {
	t.Parallel()
	type testInput struct {
		config  QemuPciDevices
		current QemuPciDevices
	}
	tests := []struct {
		name   string
		input  testInput
		output error
	}{
		{name: `Valid Delete`,
			input: testInput{config: QemuPciDevices{
				QemuPciID0: QemuPci{Delete: true}}}},
		{name: `Valid Mapping update`,
			input: testInput{
				config: QemuPciDevices{
					QemuPciID1: QemuPci{
						Mapping: &QemuPciMapping{}}},
				current: QemuPciDevices{
					QemuPciID1: QemuPci{
						Mapping: &QemuPciMapping{
							ID: util.Pointer(ResourceMappingPciID("aaa"))}}}}},
		{name: `Valid Raw update`,
			input: testInput{
				config: QemuPciDevices{
					QemuPciID2: QemuPci{
						Raw: &QemuPciRaw{}}},
				current: QemuPciDevices{
					QemuPciID2: QemuPci{
						Raw: &QemuPciRaw{
							ID: util.Pointer(PciID("0000:00:00"))}}}}},
		{name: `Invalid update errors.New(QemuPci_Error_MutualExclusive)`,
			input: testInput{
				config: QemuPciDevices{
					QemuPciID3: QemuPci{
						Mapping: &QemuPciMapping{ID: util.Pointer(ResourceMappingPciID("aaa"))},
						Raw:     &QemuPciRaw{ID: util.Pointer(PciID("0000:00:00"))}}},
				current: QemuPciDevices{
					QemuPciID3: QemuPci{}}},
			output: errors.New(QemuPci_Error_MutualExclusive)},
		{name: `Invalid errors.New(QemuPciID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				16: QemuPci{}}},
			output: errors.New(QemuPciID_Error_Invalid)},
		{name: `Invalid errors.New(QemuPci_Error_MutualExclusive)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID4: QemuPci{
					Mapping: &QemuPciMapping{
						ID: util.Pointer(ResourceMappingPciID("aaa"))},
					Raw: &QemuPciRaw{
						ID: util.Pointer(PciID("0000:00:00"))}}}},
			output: errors.New(QemuPci_Error_MutualExclusive)},
		{name: `Invalid errors.New(QemuPciMapping_Error_RequiredID)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID5: QemuPci{
					Mapping: &QemuPciMapping{}}}},
			output: errors.New(QemuPciMapping_Error_RequiredID)},
		{name: `Invalid errors.New(QemuPciRaw_Error_RequiredID)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID6: QemuPci{
					Raw: &QemuPciRaw{}}}},
			output: errors.New(QemuPciRaw_Error_RequiredID)},
		{name: `Invalid errors.New(ResourceMappingPciID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID7: QemuPci{
					Mapping: &QemuPciMapping{
						ID: util.Pointer(ResourceMappingPciID("a0%^#"))}}}},
			output: errors.New(ResourceMappingPciID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID8: QemuPci{
					Mapping: &QemuPciMapping{
						ID:       util.Pointer(ResourceMappingPciID("aaa")),
						DeviceID: util.Pointer(PciDeviceID("a0%^#"))}}}},
			output: errors.New(PciDeviceID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciMediatedDevice_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID8: QemuPci{
					Mapping: &QemuPciMapping{
						ID:   util.Pointer(ResourceMappingPciID("aaa")),
						MDev: util.Pointer(PciMediatedDevice("vendor,-643"))}}}},
			output: errors.New(PciMediatedDevice_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciSubDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID9: QemuPci{
					Mapping: &QemuPciMapping{
						ID:          util.Pointer(ResourceMappingPciID("aaa")),
						SubDeviceID: util.Pointer(PciSubDeviceID("a0%^#"))}}}},
			output: errors.New(PciSubDeviceID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciSubVendorID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID10: QemuPci{
					Mapping: &QemuPciMapping{
						ID:          util.Pointer(ResourceMappingPciID("aaa")),
						SubVendorID: util.Pointer(PciSubVendorID("a0%^#"))}}}},
			output: errors.New(PciSubVendorID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciVendorID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID11: QemuPci{
					Mapping: &QemuPciMapping{
						ID:       util.Pointer(ResourceMappingPciID("aaa")),
						VendorID: util.Pointer(PciVendorID("a0%^#"))}}}},
			output: errors.New(PciVendorID_Error_Invalid)},
		{name: `Invalid errors.New(PciID_Error_MaximumFunction)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID12: QemuPci{
					Raw: &QemuPciRaw{ID: util.Pointer(PciID("0000:00:00.8"))}}}},
			output: errors.New(PciID_Error_MaximumFunction)},
		{name: `Invalid Raw errors.New(PciDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID13: QemuPci{
					Raw: &QemuPciRaw{
						ID:       util.Pointer(PciID("0000:00:00")),
						DeviceID: util.Pointer(PciDeviceID("a0%^#"))}}}},
			output: errors.New(PciDeviceID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciMediatedDevice_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID13: QemuPci{
					Raw: &QemuPciRaw{
						ID:   util.Pointer(PciID("0000:00:00")),
						MDev: util.Pointer(PciMediatedDevice("vendor,-643"))}}}},
			output: errors.New(PciMediatedDevice_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciSubDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID14: QemuPci{
					Raw: &QemuPciRaw{
						ID:          util.Pointer(PciID("0000:00:00")),
						SubDeviceID: util.Pointer(PciSubDeviceID("a0%^#"))}}}},
			output: errors.New(PciSubDeviceID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciSubVendorID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID15: QemuPci{
					Raw: &QemuPciRaw{
						ID:          util.Pointer(PciID("0000:00:00")),
						SubVendorID: util.Pointer(PciSubVendorID("a0%^#"))}}}},
			output: errors.New(PciSubVendorID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciVendorID_Error_Invalid)`,
			input: testInput{config: QemuPciDevices{
				QemuPciID0: QemuPci{
					Raw: &QemuPciRaw{
						ID:       util.Pointer(PciID("0000:00:00")),
						VendorID: util.Pointer(PciVendorID("a0%^#"))}}}},
			output: errors.New(PciVendorID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.config.Validate(test.input.current))
		})
	}
}

func Test_QemuPciID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  QemuPciID
		output error
	}{
		{name: `Valid`,
			input: QemuPciIDMaximum},
		{name: `Invalid`,
			input:  QemuPciIDMaximum + 1,
			output: errors.New(QemuPciID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_QemuPci_Validate(t *testing.T) {
	t.Parallel()
	type testInput struct {
		config  QemuPci
		current QemuPci
	}
	tests := []struct {
		name   string
		input  testInput
		output error
	}{
		{name: `Valid Delete`,
			input: testInput{config: QemuPci{
				Delete: true}}},
		{name: `Valid Mapping update`,
			input: testInput{
				config: QemuPci{
					Mapping: &QemuPciMapping{}},
				current: QemuPci{
					Mapping: &QemuPciMapping{
						ID: util.Pointer(ResourceMappingPciID("aaa"))}}}},
		{name: `Valid Raw update`,
			input: testInput{
				config: QemuPci{
					Raw: &QemuPciRaw{}},
				current: QemuPci{
					Raw: &QemuPciRaw{
						ID: util.Pointer(PciID("0000:00:00"))}}}},
		{name: `Invalid errors.New(QemuPci_Error_MutualExclusive)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID: util.Pointer(ResourceMappingPciID("aaa"))},
				Raw: &QemuPciRaw{
					ID: util.Pointer(PciID("0000:00:00"))}}},
			output: errors.New(QemuPci_Error_MutualExclusive)},
		{name: `Invalid errors.New(QemuPci_Error_MappedID)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{}}},
			output: errors.New(QemuPciMapping_Error_RequiredID)},
		{name: `Invalid errors.New(QemuPci_Error_RawID)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{}}},
			output: errors.New(QemuPciRaw_Error_RequiredID)},
		{name: `Invalid errors.New(ResourceMappingPciID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID: util.Pointer(ResourceMappingPciID("a0%^#"))}}},
			output: errors.New(ResourceMappingPciID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID:       util.Pointer(ResourceMappingPciID("aaa")),
					DeviceID: util.Pointer(PciDeviceID("a0%^#"))}}},
			output: errors.New(PciDeviceID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciSubDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID:          util.Pointer(ResourceMappingPciID("aaa")),
					SubDeviceID: util.Pointer(PciSubDeviceID("a0%^#"))}}},
			output: errors.New(PciSubDeviceID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciSubVendorID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID:          util.Pointer(ResourceMappingPciID("aaa")),
					SubVendorID: util.Pointer(PciSubVendorID("a0%^#"))}}},
			output: errors.New(PciSubVendorID_Error_Invalid)},
		{name: `Invalid Mapping errors.New(PciVendorID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Mapping: &QemuPciMapping{
					ID:       util.Pointer(ResourceMappingPciID("aaa")),
					VendorID: util.Pointer(PciVendorID("a0%^#"))}}},
			output: errors.New(PciVendorID_Error_Invalid)},
		{name: `Invalid errors.New(PciID_Error_MaximumFunction)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{ID: util.Pointer(PciID("0000:00:00.8"))}}},
			output: errors.New(PciID_Error_MaximumFunction)},
		{name: `Invalid Raw errors.New(PciDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{
					ID:       util.Pointer(PciID("0000:00:00")),
					DeviceID: util.Pointer(PciDeviceID("a0%^#"))}}},
			output: errors.New(PciDeviceID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciSubDeviceID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{
					ID:          util.Pointer(PciID("0000:00:00")),
					SubDeviceID: util.Pointer(PciSubDeviceID("a0%^#"))}}},
			output: errors.New(PciSubDeviceID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciSubVendorID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{
					ID:          util.Pointer(PciID("0000:00:00")),
					SubVendorID: util.Pointer(PciSubVendorID("a0%^#"))}}},
			output: errors.New(PciSubVendorID_Error_Invalid)},
		{name: `Invalid Raw errors.New(PciVendorID_Error_Invalid)`,
			input: testInput{config: QemuPci{
				Raw: &QemuPciRaw{
					ID:       util.Pointer(PciID("0000:00:00")),
					VendorID: util.Pointer(PciVendorID("a0%^#"))}}},
			output: errors.New(PciVendorID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.config.Validate(test.input.current))
		})
	}
}

func Test_PciDeviceID_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciDeviceID
		output string
	}{
		{name: `No prefix`,
			input:  "ffff",
			output: "0xffff"},
		{name: `With prefix`,
			input:  "0x0000",
			output: "0x0000"},
		{name: `Empty`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.String())
		})
	}
}

func Test_PciDeviceID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciDeviceID
		output error
	}{
		{name: `Valid Maximum`,
			input: "0xffff"},
		{name: `Valid Minimum`,
			input: "0x0000"},
		{name: `Valid no prefix`,
			input: "8086"},
		{name: `Valid empty`,
			input: ""},
		{name: `Invalid not hex`,
			input:  "0xg000",
			output: errors.New(PciDeviceID_Error_Invalid)},
		{name: `Invalid Maximum`,
			input:  "0x1ffff",
			output: errors.New(PciDeviceID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_PciMediatedDevice_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciMediatedDevice
		output error
	}{
		{name: `Valid`,
			input: "vendor-643"},
		{name: `Invalid`,
			input:  "vendor,-643",
			output: errors.New(PciMediatedDevice_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_PciSubDeviceID_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciSubDeviceID
		output string
	}{
		{name: `No prefix`,
			input:  "ffff",
			output: "0xffff"},
		{name: `With prefix`,
			input:  "0x0000",
			output: "0x0000"},
		{name: `Empty`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.String())
		})
	}
}

func Test_PciSubDeviceID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciSubDeviceID
		output error
	}{
		{name: `Valid Maximum`,
			input: "0xffff"},
		{name: `Valid Minimum`,
			input: "0x0000"},
		{name: `Valid no prefix`,
			input: "8086"},
		{name: `Valid empty`,
			input: ""},
		{name: `Invalid not hex`,
			input:  "0xg000",
			output: errors.New(PciSubDeviceID_Error_Invalid)},
		{name: `Invalid Maximum`,
			input:  "0x1ffff",
			output: errors.New(PciSubDeviceID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_PciSubVendorID_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciSubVendorID
		output string
	}{
		{name: `No prefix`,
			input:  "ffff",
			output: "0xffff"},
		{name: `With prefix`,
			input:  "0x0000",
			output: "0x0000"},
		{name: `Empty`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.String())
		})
	}
}

func Test_PciSubVendorID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciSubVendorID
		output error
	}{
		{name: `Valid Maximum`,
			input: "0xffff"},
		{name: `Valid Minimum`,
			input: "0x0000"},
		{name: `Valid no prefix`,
			input: "8086"},
		{name: `Valid empty`,
			input: ""},
		{name: `Invalid not hex`,
			input:  "0xg000",
			output: errors.New(PciSubVendorID_Error_Invalid)},
		{name: `Invalid Maximum`,
			input:  "0x1ffff",
			output: errors.New(PciSubVendorID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_PciVendorID_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciVendorID
		output string
	}{
		{name: `No prefix`,
			input:  "ffff",
			output: "0xffff"},
		{name: `With prefix`,
			input:  "0x0000",
			output: "0x0000"},
		{name: `Empty`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.String())
		})
	}
}

func Test_PciVendorID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  PciVendorID
		output error
	}{
		{name: `Valid Maximum`,
			input: "0xffff"},
		{name: `Valid Minimum`,
			input: "0x0000"},
		{name: `Valid no prefix`,
			input: "8086"},
		{name: `Valid empty`,
			input: ""},
		{name: `Invalid not hex`,
			input:  "0xg000",
			output: errors.New(PciVendorID_Error_Invalid)},
		{name: `Invalid Maximum`,
			input:  "0x1ffff",
			output: errors.New(PciVendorID_Error_Invalid)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.Validate())
		})
	}
}

func Test_PciID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []PciID
		output error
	}{
		{name: `Valid`,
			input: []PciID{"1234:56:78", "0000:00:00.0"}},
		{name: `Invalid errors.New(PciID_Error_MissingBus)`,
			input:  []PciID{"0000"},
			output: errors.New(PciID_Error_MissingBus)},
		{name: `Invalid errors.New(PciID_Error_MissingDevice)`,
			input:  []PciID{"0000:00"},
			output: errors.New(PciID_Error_MissingDevice)},
		{name: `Invalid errors.New(PciID_Error_LengthDomain)`,
			input:  []PciID{"0:00:00", "0:00:00.0", "00:00:00", "00:00:00.0", "000:00:00", "000:00:00.0", "00000:00:00", "00000:00:00.0"},
			output: errors.New(PciID_Error_LengthDomain)},
		{name: `Invalid errors.New(PciID_Error_InvalidDomain)`,
			input:  []PciID{"gggg:00:00", "gggg:00:00.0"},
			output: errors.New(PciID_Error_InvalidDomain)},
		{name: `Invalid errors.New(PciID_Error_LengthBus)`,
			input:  []PciID{"0000:0:00", "0000:0:00.0", "0000:000:00", "0000:000:00.0"},
			output: errors.New(PciID_Error_LengthBus)},
		{name: `Invalid errors.New(PciID_Error_InvalidBus)`,
			input:  []PciID{"0000:gg:00", "0000:gg:00.0"},
			output: errors.New(PciID_Error_InvalidBus)},
		{name: `Invalid errors.New(PciID_Error_LengthDevice)`,
			input:  []PciID{"0000:00:0", "0000:00:0.0", "0000:00:000", "0000:00:000.0"},
			output: errors.New(PciID_Error_LengthDevice)},
		{name: `Invalid errors.New(PciID_Error_InvalidDevice)`,
			input:  []PciID{"0000:00:gg", "0000:00:gg.0"},
			output: errors.New(PciID_Error_InvalidDevice)},
		{name: `Invalid errors.New(PciID_Error_InvalidFunction)`,
			input:  []PciID{"0000:00:00.", "0000:00:00.a"},
			output: errors.New(PciID_Error_InvalidFunction)},
		{name: `Invalid errors.New(PciID_Error_MaximumFunction)`,
			input:  []PciID{"0000:00:00.8", "0000:00:00.76"},
			output: errors.New(PciID_Error_MaximumFunction)},
	}
	for _, test := range tests {
		for _, item := range test.input {
			t.Run(test.name+" :"+item.String(), func(t *testing.T) {
				require.Equal(t, test.output, item.Validate())
			})
		}
	}
}
