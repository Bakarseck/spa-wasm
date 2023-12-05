package pages

import "stack/frameGo"

type Pages interface {
	View() []frameGo.View
}
