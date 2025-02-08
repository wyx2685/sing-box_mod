package trojan

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common"
)

func (h *Inbound) AddUsers(users []option.TrojanUser) error {
	h.users = append(h.users, users...)
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, user option.TrojanUser) int {
		return index
	}), common.Map(h.users, func(user option.TrojanUser) string {
		return user.Password
	}))
	return err
}
func (h *Inbound) DelUsers(names []string) error {
	nameMap := make(map[string]struct{}, len(names))
	for _, name := range names {
		nameMap[name] = struct{}{}
	}
	filteredUsers := make([]option.TrojanUser, 0, len(h.users))
	for _, user := range h.users {
		if _, found := nameMap[user.Name]; !found {
			filteredUsers = append(filteredUsers, user)
		}
	}
	h.users = filteredUsers
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, user option.TrojanUser) int {
		return index
	}), common.Map(h.users, func(user option.TrojanUser) string {
		return user.Password
	}))
	return err
}
