package task01

/**
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
输入：strs = ["dog","racecar","car"]
*/
//strs = []string{"flower", "flow", "flight"}:
func LongestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	for i := 0; i < len(str[0]); i++ {
		for j := 1; j < len(str); j++ {
			if i >= len(str[j]) || str[j][i] != str[0][i] {
				//第一个字符串，从0开始截取，到 i 位置 字符串
				return str[0][:i]
			}
		}
	}
	return str[0]
}
