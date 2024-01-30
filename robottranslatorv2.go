package main

import (
	"fmt"
	"strings"
)

type node struct {
	item string
	next *node
}

type stack struct {
	top  *node
	size int
}

func (S *stack) push(newItem string) {
	newNode := &node{item: newItem, next: nil}
	if S == nil {
		S.top = newNode
		S.size++
	} else {
		newNode.next = S.top
		S.top = newNode
		S.size++
	}
}

func (S *stack) pop() string {
	if S.top == nil {
		return ""
	}
	popValue := S.top.item
	S.top = S.top.next
	S.size--
	return popValue
}

func removeX(S stack, cmdSlice []string) stack {
	for _, cmd := range cmdSlice {
		if cmd == "X" {
			S.pop()
		} else {
			S.push(cmd)
		}
	}
	return S
}

func moveStackValueToSlice(S stack) []string {
	var sliceOfS []string
	for S.size > 0 {
		sliceOfS = append([]string{S.pop()}, sliceOfS...)
	}
	return sliceOfS
}

func checkStrTime(counter int) (strTime string) {
	if counter > 1 {
		strTime = "times"
	} else {
		strTime = "time"
	}
	return strTime
}

// Task 3.a
func RobotTranslatorV2(cmd string) string {
	var S stack
	var direction string
	var result string
	var strTime string
	counter := 1
	cmdSlice := strings.Split(cmd, "")

	S = removeX(S, cmdSlice)
	sliceOfS := moveStackValueToSlice(S)

	for i := 0; i < len(sliceOfS); i++ {
		tempCmd := sliceOfS[i]
		switch tempCmd {
		case "R":
			direction = "right"
		case "A":
			direction = "advance"
		case "L":
			direction = "left"
		default:
			return "Invalid command"
		}

		strTime = checkStrTime(counter)

		if i == len(sliceOfS)-1 {
			result += fmt.Sprintf("Move %s %d %s", direction, counter, strTime)
			break
		}

		if tempCmd == sliceOfS[i+1] {
			counter++
		} else {
			result += fmt.Sprintf("Move %s %d %s\n", direction, counter, strTime)
			counter = 1
		}
	}
	return result
}
