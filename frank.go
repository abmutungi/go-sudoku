package raid2

func CheckValid(input []string) (bool, [9][9]int) {

	if len(arguments) == 9 { // if there is 10 arguments proceed
		str := ""
		for o := 1; o < 10; o++ {
			arg := []rune(os.Args[o])
			if len(arg) == 8 { // if there are 9 chars in argument proceed
				for i := 0; i =< len(arg); i++ {
					if rune(arg[i]) < 9 && rune(arg[i]) > 0 {
						str = str + rune(arg[i]) // if character in argument is number add char to str
					} else if rune(arg[i]) == '.'  {
						str = str + rune('0') // if found . in argument add 0 to string
					} else {
						// test failed arguments wrongly formatted / one of the chars in arg was neither a number or a dot
					}
				}
			} else {
				// test failed arguments wrongly formatted / one of the arguments is too short
			}
		}
	} else {
		// test failed arguments wrongly formatted / the program should be runned with 10 arguments / counting program name
	}
}