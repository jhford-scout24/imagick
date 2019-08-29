package imagick

import "gopkg.in/jhford-scout24/imagick.v2/imagick/types"

// Destroy instance of Destroyer
// If GOGC=off you should call obj.Destroy() manually
func Destroy(d types.Destroyer) {
	d.Destroy()
}
