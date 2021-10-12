package echohelper

var (
	actions = make([]Action, 0)
)

func AddAction(action Action) {
	actions = append(actions, action)
}

func Actions() []Action {
	return actions
}
