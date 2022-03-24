package main

import (
	"fmt"
	"os"
	"time"
)

type mode int64

const (
	normal mode = 1
	funky       = 7
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

	fmt.Println(c.question)
	var response string
	fmt.Scanln(&response)

	end := time.Now()

	rd := ResponseDetails{c, response, start, end}

	return &rd
}

func check(rd *ResponseDetails) {
	response_time := rd.end.Sub(rd.start).Milliseconds()
	if response_time >= rd.challenge.cutoff {
		fmt.Printf("Sorry, you must answer within %d milliseconds.n", rd.challenge.cutoff)
		fmt.Println("Goodbye.")
		os.Exit(2)
	} else if rd.response != rd.challenge.answer {
		fmt.Println("Sorry, this is not an acceptable response.")
		fmt.Println("Goodbye.")
		os.Exit(3)
	}
}

func run_level(l *Level) bool {
	for _, c := range l.challenges {
		check(ask(&c))
	}
	fmt.Println(l.congrats_msg)
	return true
}

// Level 1: Very simple math. Should be able to be done in the person's head within 10 seconds.
var level1 Level = Level{
	"Congratulations, you passed!\nFlag:abcdefghijklmnop",
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
	"Congratulations, you passed Level 2.\nFlag:1234\nLets speed things up.",
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
	"Congratulations, you passed Level 5.\nWe're just getting started.",
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

func main() {
	_ = run_level(&level1)
	for i := 1; i <= 3; i++ {
		fmt.Scanln()
	}
	fmt.Println("Oh, there's more here? Interesting....")
	for i := 1; i <= 7; i++ {
		fmt.Scanln()
	}
	fmt.Println("Let's start Level 2, I guess.")
	_ = run_level(&level2)
	_ = run_level(&level3)
	_ = run_level(&level4)
	_ = run_level(&level5)

}
