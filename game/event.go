package game

type event interface {
	Key() byte
	Name() string
}

type keyPressed struct {
	key byte
}

func (k *keyPressed) Key() byte {
	return k.key
}
func (k *keyPressed) Name() string {
	return "keyPressed"
}

type eventNone struct {
	key byte
}

func (k *eventNone) Key() byte {
	return 0
}
func (k *eventNone) Name() string {
	return "none"
}
