package components

import (
	"github.com/templui/templui/components/icon"
)

var stats = []stat{
	{
		value: "340+",
		label: "Routes Covered",
		icon:  icon.Route(icon.Props{Class: "size-6 text-gray-950"}),
	},
	{
		value: "180K",
		label: "Daily Riders",
		icon:  icon.Users(icon.Props{Class: "size-6 text-gray-950"}),
	},
	{
		value: "98.4%",
		label: "System Uptime",
		icon:  icon.ShieldCheck(icon.Props{Class: "size-6 text-gray-950"}),
	},
	{
		value: "< 5ms",
		label: "Route Refresh",
		icon:  icon.Zap(icon.Props{Class: "size-6 text-gray-950"}),
	},
}

var mockPins = []busPin{
	{
		top:   "25%",
		left:  "30%",
		route: "46",
		delay: "0s",
	},
	{
		top:   "55%",
		left:  "60%",
		route: "111",
		delay: "0.5s",
	},
	{
		top:   "40%",
		left:  "75%",
		route: "45",
		delay: "1s",
	},
	{
		top:   "70%",
		left:  "20%",
		route: "24",
		delay: "0.3s",
	},
}

var mockBusRoutes = []busRoute{
	{
		route:    "Route 46",
		timeline: "3 min",
		seats:    "8 seats",
	},
	{
		route:    "Route 111",
		timeline: "12 min",
		seats:    "14 seats",
	},
	{
		route:    "Route 24",
		timeline: "7 min",
		seats:    "Full",
	},
}

var mockNotifs = []notif{
	{
		message:   "Route 46 arriving in 4 min at Westlands",
		timeline:  "Just now",
		notifType: "urgent",
	},
	{
		message:   "Your usual Route 34 is running 8 min late",
		timeline:  "2 min ago",
		notifType: "warn",
	},
	{
		message:   "Route 12 now has seats — boarding at Archives",
		timeline:  "5 min ago",
		notifType: "info",
	},
}

var mockRoutePrices = []fare{
	{
		from:     "Westlands",
		to:       "CBD",
		fare:     "KES 50",
		duration: "18 min",
	},
	{
		from:     "Ngong Rd",
		to:       "Kencom",
		fare:     "KES 80",
		duration: "25 min",
	},
	{
		from:     "Thika Rd",
		to:       "GPO",
		fare:     "KES 100",
		duration: "45min",
	},
}

var onboardingSteps = []onboardingStep{
	{
		number:      "01",
		title:       "Pick your stage",
		description: "Type your nearest bus stage or tap it on the map. Drago can find bus stops in your city",
		icon:        icon.MapPin(icon.Props{Class: "size-6"}),
		color:       "bg-blue-50 text-blue-700",
	},
	{
		number:      "02",
		title:       "See your options",
		description: "Instantly see every bus on your route — ETAs, current occupancy, and the cheapest fare.",
		icon:        icon.Eye(icon.Props{Class: "size-6"}),
		color:       "accent-current/10 accent-black",
	},
	{
		number:      "03",
		title:       "Board with confidence",
		description: "Pay what you already know. Get a ping when your bus is nearby. Walk to the stage, board.",
		icon:        icon.CircleCheck(icon.Props{Class: "size-6"}),
		color:       "bg-green-50 text-green-700",
	},
}
