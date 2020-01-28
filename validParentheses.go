import "fmt"

func pop(s *[]string) string {
    n := len(*s)-1
    v := (*s)[n]
    *s = (*s)[:n]
    return v
}

func push(s *[]string, e string) {
    *s = append((*s), e)    
}

func isEmpty(s *[]string) bool {
    if len(*s) != 0 {
        return false
    }
    return true
}

func isValid(s string) bool {
    
    stack := []string{}
    p := &stack
    
    sArr := []rune(s)
    
    start := map[string]string {
        "(": ")",
        "[": "]",
        "{": "}",
    }
    end := map[string]string {
        ")": "(",
        "]": "[",
        "}": "{",
    }
    
    if len(sArr) == 1 {
        return false
    }
    
    for _,v := range sArr {
        
        if _,ok := start[string(v)]; ok {
            push(p, string(v))
        }
        
        if c,okay := end[string(v)]; okay {
            if isEmpty(p) {
                return false
            }
            
            if pop(p) != c {
                return false
            }
        
        }
    }
    
    if len(*p) != 0 {
        return false
    }
    
    return true
    
}