func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := ""
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}