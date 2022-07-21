package gopractice

//create a func that returns hash from string
func Hash(str string) uint32 {
	var hash uint32
	for _, c := range str {
		hash = hash*31 + uint32(c)
	}
	return hash
}
