package strategy

import goex "github.com/nntaoli-project/GoEx"

type Strategy interface {
	Init()
	OnTicker()
	OnDepth(depths ...*goex.Depth)
	Exit()
}
