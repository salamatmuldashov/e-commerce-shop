package middle

import "strconv"

func ConvertToUint(id string) uint {
	newId, _ := strconv.ParseUint(id, 10, 32)
	return uint(newId)
}
