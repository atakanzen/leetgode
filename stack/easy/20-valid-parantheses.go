type Stack struct {
  Values []interface{}
  Top int
}

func NewStack() *Stack {
  return &Stack{Values: make([]interface{}, 0), Top: -1}
}

func (s *Stack) Push(v interface{}) {
  s.Values = append(s.Values, v)
  s.Top++
}

func (s *Stack) Pop() interface{} {
  if len(s.Values) != 0 {
    val := s.Values[s.Top]
    s.Values = s.Values[:s.Top]
    s.Top--

    return val
  }

  return nil
}

func isValid(s string) bool {
  runes := []rune(s)
  myStack := NewStack()

  bracketMap := map[rune]rune{
    ']': '[',
    ')': '(',
    '}': '{',
  }

  for _, bracket := range runes {
    switch bracket {
      case '(', '[', '{':
        myStack.Push(bracket)
      default: 
        openingBracket := myStack.Pop()
        mappedBracket, _ := bracketMap[bracket]
        if openingBracket != mappedBracket {
          return false
        }
    }
  }

  return len(myStack.Values) == 0
}
