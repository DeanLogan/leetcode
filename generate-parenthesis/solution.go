package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
    var res []string
    dfs(n, n, "", &res)
    return res
}

func dfs(left, right int, path string, res *[]string) {
    if left == 0 && right == 0 {
        *res = append(*res, path)
        return
    }
    if left > 0 {
        dfs(left-1, right, path+"(", res)
    }
    if right > left {
        dfs(left, right-1, path+")", res)
    }
}