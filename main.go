package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type mode string

const (
	empty  mode = "empty"
	normal mode = "normal"
	slow   mode = "slow"
	hexxed mode = "hexxed"
)

type flag string

var (
	flagFormat string = "flag{%s}"
	level1Flag flag   = flag(fmt.Sprintf(flagFormat, "level1flag"))
	level2Flag flag   = flag(fmt.Sprintf(flagFormat, "level2flag"))
	normalFlag flag   = flag(fmt.Sprintf(flagFormat, "normalflag"))
	slowFlag   flag   = flag(fmt.Sprintf(flagFormat, "slowflag"))
	hexFlag    flag   = flag(fmt.Sprintf(flagFormat, "hexflag"))
)

type Challenge struct {
	question string
	answer   string
	cutoff   int64
}

type Level struct {
	congrats_msg string
	challenges   []Challenge
}

type ResponseDetails struct {
	challenge *Challenge
	response  string
	start     time.Time
	end       time.Time
}

func ask(c *Challenge) *ResponseDetails {
	start := time.Now()
	myfmt.Println("Q: " + c.question)
	var response string
	myfmt.Print("A: ")
	myfmt.Scanln(&response)
	end := time.Now()

	rd := ResponseDetails{c, response, start, end}

	return &rd
}

func hex(answer string) string {
	int_answer, _ := strconv.Atoi(answer)
	hex_int_answer := 6 * int_answer
	hex_str_answer := strconv.Itoa(hex_int_answer)
	return hex_str_answer
}

const long_mode_time int64 = 5000

func check(rd *ResponseDetails) mode {
	response_time := rd.end.Sub(rd.start).Milliseconds()
	correct := rd.response == rd.challenge.answer
	slow_response := response_time > long_mode_time
	timely_response := response_time <= rd.challenge.cutoff
	hexxed_response := rd.response == hex(rd.challenge.answer)

	if slow_response {
		return slow
	} else if !timely_response {
		myfmt.Printf("Sorry, you must answer within %d milliseconds.\n", rd.challenge.cutoff)
		myfmt.Println("Goodbye.")
		os.Exit(2)
	} else if hexxed_response {
		if rd.response == "0" {
			return empty
		}
		return hexxed
	} else if correct {
		return normal
	} else {
		myfmt.Println("Sorry, this is not an acceptable response.")
		myfmt.Println("Goodbye.")
		os.Exit(3)
	}
	return empty
}

func check_modes(m mode, n mode) mode {
	// Checks if the modes match or if one of them is empty. Returns the non-empty mode.
	if m == empty {
		return n
	} else if n == empty {
		return m
	} else if m == n {
		return m
	} else {
		myfmt.Println("Hmmm, something doesn't seem right here.")
		myfmt.Println("Goodbye.")
		os.Exit(4)
	}
	return empty
}

func run_level(l *Level) mode {
	var m mode = empty
	for _, c := range l.challenges {
		// Check the mode that a response would be valid for.
		user_mode := check(ask(&c))
		// Error if that mode is incompatible with the current mode
		m = check_modes(user_mode, m)
	}

	myfmt.Println(l.congrats_msg)
	return m
}

// Level 1: Very simple math. Should be able to be done in the person's head within 10 seconds.
var level1 Level = Level{
	fmt.Sprintf("Congratulations, you passed!\n%s", level1Flag),
	[]Challenge{
		{"1+1", "2", 10000},
		{"1+2", "3", 10000},
		{"2+2", "4", 10000},
		{"4+4", "8", 10000},
		{"12+10", "22", 10000},
		{"15+7", "22", 10000},
		{"30+2", "32", 10000},
		{"45+7", "52", 10000},
		{"15+17", "32", 10000},
		{"32+65", "97", 10000},
	},
}

// Level 2: More difficult math with the same amount of time.
var level2 Level = Level{
	fmt.Sprintf("Congratulations, you passed Level 2.\n%s\nLets speed things up.", level2Flag),
	[]Challenge{
		{"595+434", "1029", 10000},
		{"127+778", "905", 10000},
		{"441+931", "1372", 10000},
		{"939+631", "1570", 10000},
		{"754+634", "1388", 10000},
		{"172+123", "295", 10000},
		{"266+404", "670", 10000},
		{"103+958", "1061", 10000},
		{"254+898", "1152", 10000},
		{"910+117", "1027", 10000},
	},
}

// Level 3: Level1 but with faster responses. People may begin making scripts at this point.
var level3 Level = Level{
	"Congratulations, you passed Level 3.\nI wonder what's next....",
	[]Challenge{
		{"1+1", "2", 2000},
		{"1+2", "3", 2000},
		{"2+2", "4", 2000},
		{"4+4", "8", 2000},
		{"12+10", "22", 2000},
		{"15+7", "22", 2000},
		{"30+2", "32", 2000},
		{"45+7", "52", 2000},
		{"15+17", "32", 2000},
		{"32+65", "97", 2000},
	},
}

// Level 4: Level2 but with faster responses. People may begin making scripts at this point.
var level4 Level = Level{
	"Congratulations, you passed Level 4.\nWe're just getting started.",
	[]Challenge{
		{"595+434", "1029", 2000},
		{"127+778", "905", 2000},
		{"441+931", "1372", 2000},
		{"939+631", "1570", 2000},
		{"754+634", "1388", 2000},
		{"172+123", "295", 2000},
		{"266+404", "670", 2000},
		{"103+958", "1061", 2000},
		{"254+898", "1152", 2000},
		{"910+117", "1027", 2000},
	},
}

// Level 5: Slow multiplication
var level5 Level = Level{
	"Congratulations, you passed Level 5.\nDon't forget PEMDAS.",
	[]Challenge{
		{"45*54", "2430", 10000},
		{"82*5", "410", 10000},
		{"10*14", "140", 10000},
		{"15*54", "810", 10000},
		{"24*45", "1080", 10000},
		{"94*55", "5170", 10000},
		{"62*50", "3100", 10000},
		{"65*43", "2795", 10000},
		{"2*35", "70", 10000},
		{"51*25", "1275", 10000},
	},
}

// Level 6: Order of operations.
var level6 Level = Level{
	"Congratulations, you passed Level 6.\nNext round is going to be A LOT bigger.",
	[]Challenge{
		{"3*2+7", "13", 6789},
		{"9+6*3", "27", 6789},
		{"1*0+6", "6", 6789},
		{"7+9*6", "61", 6789},
		{"2*4+3", "11", 6789},
		{"7*4+3", "31", 6789},
		{"6*3+5", "23", 6789},
		{"1*4+7", "11", 6789},
		{"0*7+1", "1", 6789},
		{"4+5*6", "34", 6789},
	},
}

// Level 7: Lots of simple math.
func gen_level7() *[]Challenge {
	var output []Challenge
	for i := 1; i <= 100; i++ {
		arg1 := i % 7
		arg2 := i % 11
		sum := arg1 + arg2
		question := myfmt.Sprintf("%d+%d", arg1, arg2)
		answer := myfmt.Sprint(sum)
		output = append(output, Challenge{question, answer, 10000})
	}
	return &output
}

var level7 Level = Level{
	"Congratulations, you passed Level 7.\nNot even I know what's coming up next. No, seriously.",
	*gen_level7(),
}

// Level 8: First random level.
func gen_level8() *[]Challenge {
	var output []Challenge
	for i := 1; i <= 10; i++ {
		arg1 := rand.Intn(97) + 1
		arg2 := rand.Intn(98) + 1
		sum := arg1 + arg2
		question := myfmt.Sprintf("%d+%d", arg1, arg2)
		answer := myfmt.Sprint(sum)
		output = append(output, Challenge{question, answer, 10000})
	}
	return &output
}

var level8 Level = Level{
	"Congratulations, you passed Level 8.\nWe're gonna speed it up again.",
	*gen_level8(),
}

// Level 9: Fast random level.
func gen_level9() *[]Challenge {
	var output []Challenge
	for i := 1; i <= 1000; i++ {
		arg1 := rand.Intn(97) + 1
		arg2 := rand.Intn(98) + 1
		sum := arg1 + arg2
		question := myfmt.Sprintf("%d+%d", arg1, arg2)
		answer := myfmt.Sprint(sum)
		output = append(output, Challenge{question, answer, 2000})
	}
	return &output
}

var level9 Level = Level{
	"Congratulations, you passed Level 9.\nNext round is going to pretty similar....but watch out for the curveballs.",
	*gen_level9(),
}

// Level 10: Fast random level again but with wildcards.
func gen_level10_helper_meta() func(int) string {
	var bash_ps_exit_q int = (rand.Intn(80) + 20)
	var bash_ps_exit_str string = "exit"

	var python_exit_q int = rand.Intn(100) + 100
	var python_exit_str string = "exit()"

	var js_exit_q int = rand.Intn(100) + 200
	var js_exit_str string = "throw new Error('Hm')"

	var php_exit_q int = rand.Intn(100) + 300
	var php_exit_str string = "die('Hm')"

	var return_q int = rand.Intn(100) + 400
	var reuturn_str string = "return"

	return func(i int) string {
		prefix := ""
		switch i {
		case bash_ps_exit_q:
			prefix = bash_ps_exit_str
		case python_exit_q:
			prefix = python_exit_str
		case js_exit_q:
			prefix = js_exit_str
		case php_exit_q:
			prefix = php_exit_str
		case return_q:
			prefix = reuturn_str
		default:
			return ""
		}
		return prefix + "\n"

	}

}

func gen_level10() *[]Challenge {
	gen_level10_helper := gen_level10_helper_meta()
	var output []Challenge
	for i := 1; i <= 1000; i++ {
		arg1 := rand.Intn(97) + 1
		arg2 := rand.Intn(98) + 1
		sum := arg1 + arg2
		question := myfmt.Sprintf("%d+%d", arg1, arg2)
		question = gen_level10_helper(i) + question
		answer := myfmt.Sprint(sum)
		output = append(output, Challenge{question, answer, 2000})
	}
	return &output
}

var level10 Level = Level{
	"Congratulations, you passed Level 10!",
	*gen_level10(),
}

func level1_break() {
	for i := 1; i <= 3; i++ {
		//fmt.Scanln OK
		fmt.Scanln()
	}
	myfmt.Println("Oh, there's more here? Interesting....")
	for i := 1; i <= 7; i++ {
		//fmt.Scanln OK
		fmt.Scanln()
	}
	myfmt.Println("Let's start Level 2, I guess.")
}

func main() {
	myfmt.Println("Shell challenge. Beta 1")
	var current_mode mode = run_level(&level1)
	if current_mode == normal {
		level1_break()
	} else if current_mode == hexxed {
		myfmt.Println("Ooh, you found the secret mode.")
	} else if current_mode == slow {
		myfmt.Println("Ooh, you found the secret mode.")
	}
	current_mode = check_modes(current_mode, run_level(&level2))
	current_mode = check_modes(current_mode, run_level(&level3))
	current_mode = check_modes(current_mode, run_level(&level4))
	current_mode = check_modes(current_mode, run_level(&level5))
	current_mode = check_modes(current_mode, run_level(&level6))
	current_mode = check_modes(current_mode, run_level(&level7))
	current_mode = check_modes(current_mode, run_level(&level8))
	current_mode = check_modes(current_mode, run_level(&level9))
	current_mode = check_modes(current_mode, run_level(&level10))

	if current_mode == normal {
		myfmt.Printf("Nice job! You won!\n%s\n", normalFlag)
		myfmt.Println("Thanks for playing! Y'all come back now, you hear?")
		// Psych them out by accepting any number of lines.
		i := 0
		for true {
			fmt.Scanln() // Using fmt instead of myfmt because I don't want to log a bunch of blank lines.
			i += 1
			if i == 10 {
				myfmt.Println("There's nothing else, I don't know what you're trying to do.")
			} else if i == 100 {
				myfmt.Println("Seriously, what do you want from me? You're done.")
			} else if i == 1000 {
				myfmt.Println("Really think that I'm going to reuse the same tricks?")
			}
		}
	} else if current_mode == hexxed {
		myfmt.Println("Congratulations on passing the secret mode!")
		myfmt.Println(hexFlag)
	} else if current_mode == slow {
		myfmt.Println("Congratulations on passing the secret mode!")
		myfmt.Println(slowFlag)
	}
}
