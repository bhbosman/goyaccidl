package objects

import (
	"context"
	rxgo "github.com/ReactiveX/RxGo"
)

type IDclArray []IDcl

func (t IDclArray) ToItemsObs() rxgo.Observable {
	return rxgo.Defer(
		[]rxgo.Producer{
			func(ctx context.Context, next chan<- rxgo.Item) {
				for _, v := range t {
					next <- rxgo.Of(v)
				}
			},
		})
}

func (t IDclArray) ToObs() rxgo.Observable {
	return rxgo.Defer(
		[]rxgo.Producer{
			func(ctx context.Context, next chan<- rxgo.Item) {
				next <- rxgo.Of(t)
			},
		})
}
