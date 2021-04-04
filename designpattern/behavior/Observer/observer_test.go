package Observer

import (
	"container/list"
	"testing"
)


func TestFactory(t *testing.T) {
	newsCenter := newsPublisher{
		subscribers: new(list.List),
	}
	subscriber1 := subscriber{name: "sub1"}
	subscriber2 := subscriber{name: "sub2"}
	subscriber3 := subscriber{name: "sub3"}

	newsCenter.add(subscriber1)
	newsCenter.add(subscriber2)
	newsCenter.add(subscriber3)

	newsCenter.update("new~~~s")

	newsCenter.delete(subscriber2)

	newsCenter.update("1,3 new~~~s")
}


