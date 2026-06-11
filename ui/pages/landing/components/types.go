package components

import "github.com/a-h/templ"

type Stat struct {
	value string
	label string
	icon  templ.Component
}

type BusPin struct {
	top   string
	left  string
	route string
	delay string
}

type BusRoute struct {
	route    string
	timeline string
	seats    string
}

type Notif struct {
	message   string
	timeline  string
	notifType string
}

type Fare struct {
	from     string
	to       string
	duration string
	fare     string
}
