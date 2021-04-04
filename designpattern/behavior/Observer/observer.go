package Observer

import (
	"container/list"
	"fmt"
)

type subscriber struct {
	name string
}

func (s *subscriber) onUpdate(msg string) {
	fmt.Println("구독자:", s.name, "메시지:", msg)
}

type iPublisher interface {
	update(text string)
	add(subscriber)
	delete(subscriber)
}

type newsPublisher struct {
	subscribers *list.List
}

func (n *newsPublisher) update(text string) {
	for e := n.subscribers.Front(); e != nil; e = e.Next() {
		s := e.Value.(subscriber)
		s.onUpdate(text)
	}
}

func (n *newsPublisher) add(s subscriber) {
	n.subscribers.PushBack(s)
}

func (n *newsPublisher) delete(s subscriber) {
	var removeItem subscriber
	for e := n.subscribers.Front(); e != nil; e = e.Next() {
		removeItem =  e.Value.(subscriber)
		if s == removeItem {
			n.subscribers.Remove(e)
			break;
		}
	}
}
