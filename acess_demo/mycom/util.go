package mycom

/**
  public函数-要大写字母开头
*/
func IsBlank(str string) bool {
	return isEmpty(str)
}

/**
  小写字母开头
  同包可以访问
  其他包不能访问
*/
func isEmpty(str string) bool {
	if len(str) > 0 {
		return true
	}

	return false
}
