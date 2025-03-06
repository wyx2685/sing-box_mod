package vmess

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common"
)

func (h *Inbound) AddUsers(users []option.VMessUser) error {
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
func (h *Inbound) DelUsers(name []string) error {
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
