package proxmox

type QemuRef struct {
	ID   GuestID
	Node NodeName
}

var _ GuestRef = (*QemuRef)(nil)

func (ref QemuRef) GetID() GuestID     { return ref.ID }
func (ref QemuRef) GetNode() NodeName  { return ref.Node }
func (ref QemuRef) GetType() GuestType { return GuestLxc }

func (ref QemuRef) generalize() guestRef {
	return guestRef{
		ID:   ref.GetID(),
		Node: ref.GetNode(),
		Type: ref.GetType(),
	}
}

func (ref QemuRef) legacy() *VmRef {
	return &VmRef{
		node:   ref.GetNode(),
		vmId:   ref.GetID(),
		vmType: ref.GetType(),
	}
}
