package proxmox

type GuestRef interface {
	GetID() GuestID
	GetNode() NodeName
	GetType() GuestType
}

// This will be the future replacement of VmRef
type guestRef struct {
	ID   GuestID
	Node NodeName
	Type GuestType
}
