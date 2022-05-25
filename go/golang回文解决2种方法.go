package main

// param is string type ,if s int type, shoule change type to string
func isPalindrome1(s string) bool {
	result := []rune(s)
	length := len(result)
	for i, _ := range result {
		if i == length/2 {
			break
		}
		tail := length - i - 1
		if result[i] != result[tail] {
			return false
		}
	}
	return true

}

//param's type is int ,should change int to string
func isPalindrome2(x int) bool {
	//str := strconv.Itoa(x)
	str := fmt.Sprintf("%d", x)
	length := len(str)
	for i := 0 ; i < length/2; i++{
		if str[i] != str[length -1 -i]{
			return false
		}
	}
	return true
}

