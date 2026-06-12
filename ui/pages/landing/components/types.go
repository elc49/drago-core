package components

import "github.com/a-h/templ"

type stat struct {
	value string
	label string
	icon  templ.Component
}

type busPin struct {
	top   string
	left  string
	route string
	delay string
}

type busRoute struct {
	route    string
	timeline string
	seats    string
}

type notif struct {
	message   string
	timeline  string
	notifType string
}

type fare struct {
	from     string
	to       string
	duration string
	fare     string
}

type onboardingStep struct {
	number      string
	title       string
	description string
	icon        templ.Component
	color       string
}
