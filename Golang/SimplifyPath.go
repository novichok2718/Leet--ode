package main

import (
	"container/list"
	"strings"
)

func simplifyPath(path string) string {
	lst := list.New()
	token := strings.Split(path, "/")
	for _, name := range token {
		if name == ".." && lst.Len() != 0{
			lst.Remove(lst.Back())
		} else if name != "" && name != "." && name != ".." {
			lst.PushBack(name)
		}
	}
	ans := ""
	for iter := lst.Front(); iter != nil; iter = iter.Next() {
		ans += "/" + iter.Value.(string)
	}
	
	if ans == "" {	return "/"}
	return ans
}
