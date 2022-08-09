package typeconv

import "strconv"

func ConvertStringIdToUintId(userId string) (uint, error) {
	convUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(convUserId), nil
}
