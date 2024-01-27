package inbound

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common"
)

func (h *ShadowsocksMulti) AddUsers(users []option.ShadowsocksUser) error {
	h.users = append(h.users, users...)

	err := h.service.UpdateUsersWithPasswords(common.MapIndexed(h.users, func(index int, user option.ShadowsocksUser) int {
		return index
	}), common.Map(h.users, func(user option.ShadowsocksUser) string {
		return user.Password
	}))

	return err
}

func (h *ShadowsocksMulti) DelUsers(names []string) error {
	toDelete := make(map[string]struct{})
	for _, name := range names {
		toDelete[name] = struct{}{}
	}

	remaining := make([]option.ShadowsocksUser, 0, len(h.users))
	for _, user := range h.users {
		if _, found := toDelete[user.Name]; !found {
			remaining = append(remaining, user)
		}
	}

	h.users = remaining

	err := h.service.UpdateUsersWithPasswords(common.MapIndexed(h.users, func(index int, user option.ShadowsocksUser) int {
		return index
	}), common.Map(h.users, func(user option.ShadowsocksUser) string {
		return user.Password
	}))

	return err
}

func (h *Trojan) AddUsers(users []option.TrojanUser) error {
	if cap(h.users)-len(h.users) >= len(users) {
		h.users = append(h.users, users...)
	} else {
		tmp := make([]option.TrojanUser, 0, len(h.users)+len(users)+10)
		tmp = append(tmp, h.users...)
		tmp = append(tmp, users...)
		h.users = tmp
	}
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, user option.TrojanUser) int {
		return index
	}), common.Map(h.users, func(user option.TrojanUser) string {
		return user.Password
	}))
	if err != nil {
		return err
	}
	return nil
}

func (h *Trojan) DelUsers(names []string) error {
	is := make([]int, 0, len(names))
	ulen := len(names)
	for i := range h.users {
		for _, n := range names {
			if h.users[i].Name == n {
				is = append(is, i)
				ulen--
			}
			if ulen == 0 {
				break
			}
		}
	}
	ulen = len(h.users)
	for _, i := range is {
		h.users[i] = h.users[ulen-1]
		h.users[ulen-1] = option.TrojanUser{}
		h.users = h.users[:ulen-1]
		ulen--
	}
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, user option.TrojanUser) int {
		return index
	}), common.Map(h.users, func(user option.TrojanUser) string {
		return user.Password
	}))
	return err
}

func (h *VLESS) AddUsers(users []option.VLESSUser) error {
	h.users = append(h.users, users...)

	h.service.UpdateUsers(
		common.MapIndexed(h.users, func(index int, it option.VLESSUser) int {
			return index
		}),
		common.Map(h.users, func(it option.VLESSUser) string {
			return it.UUID
		}),
		common.Map(h.users, func(it option.VLESSUser) string {
			return it.Flow
		}),
	)
	return nil
}

func (h *VLESS) DelUsers(names []string) error {
	toDelete := make(map[string]struct{})
	for _, name := range names {
		toDelete[name] = struct{}{}
	}

	remaining := make([]option.VLESSUser, 0)
	for _, user := range h.users {
		if _, found := toDelete[user.Name]; !found {
			remaining = append(remaining, user)
		}
	}

	h.users = remaining

	h.service.UpdateUsers(
		common.MapIndexed(h.users, func(index int, it option.VLESSUser) int {
			return index
		}),
		common.Map(h.users, func(it option.VLESSUser) string {
			return it.UUID
		}),
		common.Map(h.users, func(it option.VLESSUser) string {
			return it.Flow
		}),
	)
	return nil
}

func (h *VMess) AddUsers(users []option.VMessUser) error {
	if cap(h.users)-len(h.users) >= len(users) {
		h.users = append(h.users, users...)
	} else {
		tmp := make([]option.VMessUser, 0, len(h.users)+len(users)+10)
		tmp = append(tmp, h.users...)
		tmp = append(tmp, users...)
		h.users = tmp
	}
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, it option.VMessUser) int {
		return index
	}), common.Map(h.users, func(it option.VMessUser) string {
		return it.UUID
	}), common.Map(h.users, func(it option.VMessUser) int {
		return it.AlterId
	}))
	if err != nil {
		return err
	}
	return nil
}

func (h *VMess) DelUsers(name []string) error {
	is := make([]int, 0, len(name))
	ulen := len(name)
	for i := range h.users {
		for _, u := range name {
			if h.users[i].Name == u {
				is = append(is, i)
				ulen--
			}
			if ulen == 0 {
				break
			}
		}
	}
	ulen = len(h.users)
	for _, i := range is {
		h.users[i] = h.users[ulen-1]
		h.users[ulen-1] = option.VMessUser{}
		h.users = h.users[:ulen-1]
		ulen--
	}
	err := h.service.UpdateUsers(common.MapIndexed(h.users, func(index int, it option.VMessUser) int {
		return index
	}), common.Map(h.users, func(it option.VMessUser) string {
		return it.UUID
	}), common.Map(h.users, func(it option.VMessUser) int {
		return it.AlterId
	}))
	if err != nil {
		return err
	}
	return nil
}
