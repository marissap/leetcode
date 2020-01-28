import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    
    res := []int{}
    m := make(map[int]int)
    
    for i, n := range nums {  
        m[n] = i
    }
    
    for i,v := range nums {
        k := target - v
        if val, ok := m[k]; ok && val != i {
            res = append(res, i)
            res = append(res, val) 
            return res
        }
        
    }
    
    return res
}
