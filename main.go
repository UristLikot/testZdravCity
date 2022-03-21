package main

func main() {
	arr := [][]string{

		{"1014065", "1", "4708", "4709"},

		{"1014071", "1", "4708", "4697"},

		{"1014130", "2", "9416", "8004"},

		{"1014130", "2", "9416", "1412"},

		{"1014238", "1", "4708", "4694"},

		{"1014240", "3", "14124", "14090"},
	}
	count := 0

	for {
		var slc []string

		for i := 0; i < len(arr); i++ {
			if count >= len(arr[i]) {
				break
			}

			arr1 := arr[i]

			slc = append(slc, arr1[count])
		}

		if count >= len(arr) {
			break
		}
		count++
		// fmt.Println(findDupes(slc)) TODO
	}
}
