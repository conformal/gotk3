package gdk

type EventMask int

const (
	EXPOSURE_MASK            EventMask = 1 << 1
	POINTER_MOTION_MASK      EventMask = 1 << 2
	POINTER_MOTION_HINT_MASK EventMask = 1 << 3
	BUTTON_MOTION_MASK       EventMask = 1 << 4
	BUTTON1_MOTION_MASK      EventMask = 1 << 5
	BUTTON2_MOTION_MASK      EventMask = 1 << 6
	BUTTON3_MOTION_MASK      EventMask = 1 << 7
	BUTTON_PRESS_MASK        EventMask = 1 << 8
	BUTTON_RELEASE_MASK      EventMask = 1 << 9
	KEY_PRESS_MASK           EventMask = 1 << 10
	KEY_RELEASE_MASK         EventMask = 1 << 11
	ENTER_NOTIFY_MASK        EventMask = 1 << 12
	LEAVE_NOTIFY_MASK        EventMask = 1 << 13
	FOCUS_CHANGE_MASK        EventMask = 1 << 14
	STRUCTURE_MASK           EventMask = 1 << 15
	PROPERTY_CHANGE_MASK     EventMask = 1 << 16
	VISIBILITY_NOTIFY_MASK   EventMask = 1 << 17
	PROXIMITY_IN_MASK        EventMask = 1 << 18
	PROXIMITY_OUT_MASK       EventMask = 1 << 19
	SUBSTRUCTURE_MASK        EventMask = 1 << 20
	SCROLL_MASK              EventMask = 1 << 21
	ALL_EVENTS_MASK          EventMask = 0x3FFFFE
)
