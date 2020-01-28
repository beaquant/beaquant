package strategy

type Strategy interface {
	Init()
	OnData()
	Exit()
}
