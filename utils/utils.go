package utils

func ValidateBrackets(s string) bool {
	var stack []rune

	for _, bracket := range s {
		n := len(stack) - 1

		if bracket == '}' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '{' {
				return false
			}
		} else if bracket == ']' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '[' {
				return false
			}
		} else if bracket == ')' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '(' {
				return false
			}
		} else {
			stack = append(stack, bracket)
		}
	}

	return len(stack) == 0
}

func Fix(s string) string {
	type st struct {
		value string
		i     int
	}

	stack := []string{}
	stackOfWrong := []st{}

	for i, bracket := range s {
		n := len(stack) - 1

		if i == len(s) {
			if bracket == '{' || bracket == '(' || bracket == '[' {
				s = string(remove([]rune(s), i+1))
				continue
			}
		}

		if n < 0 {
			if s[i] == '}' || s[i] == ')' || s[i] == ']' {
				s = string(remove([]rune(s), i))
				s = Fix(s)
				continue
			}
		}

		if bracket == '}' {
			//if n < 0 {
			//	s = Fix(s)
			//	continue
			//}
			current := stack[n]
			if current != "{" {
				sw := st{
					value: current,
					i:     i,
				}
				stackOfWrong = append(stackOfWrong, sw)
			}
			stack = stack[:n]
		} else if bracket == ']' {
			//if n < 0 {
			//	s = Fix(s)
			//	continue
			//}
			current := stack[n]
			if current != "[" {
				sw := st{
					value: current,
					i:     i,
				}
				stackOfWrong = append(stackOfWrong, sw)
			}
			stack = stack[:n]
		} else if bracket == ')' {
			//if n < 0 {
			//	s = Fix(s)
			//	continue
			//}
			current := stack[n]
			if current != "(" {
				sw := st{
					value: current,
					i:     i,
				}
				stackOfWrong = append(stackOfWrong, sw)
			}
			stack = stack[:n]
		} else {
			stack = append(stack, string(bracket))
		}
	}

	if len(stackOfWrong) != 0 {
		for i := 0; i < len(stackOfWrong); i++ {
			s = string(remove([]rune(s), i))
		}
	}

	if len(stack) != 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] == "{" {
				s = string(insert([]rune(s), '}', len(s)))
				stack = stack[:len(stack)-1]
				continue
			}

			if stack[i] == "(" {
				s = string(insert([]rune(s), ')', len(s)))
				stack = stack[:len(stack)-1]
				continue
			}

			if stack[i] == "[" {
				s = string(insert([]rune(s), ']', len(s)))
				stack = stack[:len(stack)-1]
				continue
			}
		}
	}

	return s
}

func insert(a []rune, c rune, i int) []rune {
	return append(a[:i], append([]rune{c}, a[i:]...)...)
}

func remove(a []rune, i int) []rune {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0
	a = a[:len(a)-1]

	return a
}
