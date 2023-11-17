//go:build with_quic

package inbound

import (
	"github.com/sagernet/sing-box/option"
)

func (h *Hysteria) AddUsers(users []option.HysteriaUser) error {
	pws := make([]string, len(users)+len(h.userNameList))
	for i := range users {
		pws[i] = users[i].AuthString
	}
	if cap(h.userNameList)-len(h.userNameList) >= len(users) {
		h.userNameList = append(h.userNameList, pws...)
	} else {
		tmp := make([]string, 0, len(h.userNameList)+len(users)+10)
		tmp = append(tmp, h.userNameList...)
		tmp = append(tmp, pws...)
		h.userNameList = tmp
	}
	indexs := make([]int, len(h.userNameList))
	for i := range h.userNameList {
		indexs[i] = i
	}
	h.service.UpdateUsers(indexs, h.userNameList)
	return nil
}

func (h *Hysteria) DelUsers(name []string) error {
	if len(name) == 0 {
		return nil
	}
	is := make([]int, 0, len(name))
	ulen := len(name)
	for i := range h.userNameList {
		for _, u := range name {
			if h.userNameList[i] == u {
				is = append(is, i)
				ulen--
			}
			if ulen == 0 {
				break
			}
		}
	}
	ulen = len(h.userNameList)
	for _, i := range is {
		h.userNameList[i] = h.userNameList[ulen-1]
		h.userNameList[ulen-1] = ""
		h.userNameList = h.userNameList[:ulen-1]
		ulen--
	}
	indexs := make([]int, len(h.userNameList))
	for i := range h.userNameList {
		indexs[i] = i
	}
	h.service.UpdateUsers(indexs, h.userNameList)
	return nil
}

func (h *Hysteria2) AddUsers(users []option.Hysteria2User) error {
	indexs := make([]int, len(users)+len(h.userNameList))
	pws := make([]string, len(users)+len(h.userNameList))
	for i := range users {
		indexs[i] = i
		pws[i] = users[i].Password
	}
	if cap(h.userNameList)-len(h.userNameList) >= len(users) {
		h.userNameList = append(h.userNameList, pws...)
	} else {
		tmp := make([]string, 0, len(h.userNameList)+len(users)+10)
		tmp = append(tmp, h.userNameList...)
		tmp = append(tmp, pws...)
		h.userNameList = tmp
	}
	h.service.UpdateUsers(indexs, h.userNameList)
	return nil
}

func (h *Hysteria2) DelUsers(name []string) error {
	if len(name) == 0 {
		return nil
	}
	is := make([]int, 0, len(name))
	ulen := len(name)
	for i := range h.userNameList {
		for _, u := range name {
			if h.userNameList[i] == u {
				is = append(is, i)
				ulen--
			}
			if ulen == 0 {
				break
			}
		}
	}
	ulen = len(h.userNameList)
	for _, i := range is {
		h.userNameList[i] = h.userNameList[ulen-1]
		h.userNameList[ulen-1] = ""
		h.userNameList = h.userNameList[:ulen-1]
		ulen--
	}
	indexs := make([]int, len(h.userNameList))
	for i := range h.userNameList {
		indexs[i] = i
	}
	h.service.UpdateUsers(indexs, h.userNameList)
	return nil
}
