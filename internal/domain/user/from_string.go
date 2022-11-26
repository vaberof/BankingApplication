package user

import "strconv"

type UserId uint

func FromString(userId string) (UserId, error) {
	convUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return 0, err
	}
	return UserId(uint(convUserId)), nil
}
