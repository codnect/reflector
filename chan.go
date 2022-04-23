package reflector

type ChanDirection int

const (
	SEND ChanDirection = 1 << iota
	RECEIVE
)

type Chan interface {
	Direction() ChanDirection
}
