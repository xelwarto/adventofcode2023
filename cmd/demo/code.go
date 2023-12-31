package demo

import (
	"bufio"
	"code/util"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var demoJson = `{"widget": {
    "debug": "on",
    "window": {
        "title": "Sample Konfabulator Widget",
        "name": "main_window",
        "width": 500,
        "height": 500
    },
    "image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    },
    "text": {
        "data": "Click Here",
        "size": 36,
        "style": "bold",
        "name": "text1",
        "hOffset": 250,
        "vOffset": 100,
        "alignment": "center",
        "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
}}`

type AutoGenerated struct {
	Widget struct {
		Debug  string `json:"debug"`
		Window struct {
			Title  string `json:"title"`
			Name   string `json:"name"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"window"`
		Image struct {
			Src       string `json:"src"`
			Name      string `json:"name"`
			HOffset   int    `json:"hOffset"`
			VOffset   int    `json:"vOffset"`
			Alignment string `json:"alignment"`
		} `json:"image"`
		Text struct {
			Data      string `json:"data"`
			Size      int    `json:"size"`
			Style     string `json:"style"`
			Name      string `json:"name"`
			HOffset   int    `json:"hOffset"`
			VOffset   int    `json:"vOffset"`
			Alignment string `json:"alignment"`
			OnMouseUp string `json:"onMouseUp,omitempty"`
		} `json:"text"`
	} `json:"widget"`
}

var DemoCmd = &cobra.Command{
	Use:   "demo",
	Short: "",
}

var testCodeCmd = &cobra.Command{
	Use:   "test-code",
	Short: "Test Code",
	Run: func(cmd *cobra.Command, args []string) {
		readFromFile()
		readFromFile2()
		unmarshalJson()
		unmarshalJsonFile()
		marshalJson()
		splitString()
		splitString2()
		splitString3()
		stringToInt()
		stringToInt2()
		parseTime()
	},
}

func init() {
	DemoCmd.AddCommand(testCodeCmd)
}

func unmarshalJson() {
	j := AutoGenerated{}
	if err := json.Unmarshal([]byte(demoJson), &j); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(j.Widget.Window.Name)
	fmt.Println(j.Widget.Image.Name)
}

func unmarshalJsonFile() {
	b, err := os.ReadFile("inputs/demo/file1.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	j := AutoGenerated{}
	if err := json.Unmarshal(b, &j); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(j.Widget.Window.Name)
	fmt.Println(j.Widget.Image.Name)
	fmt.Println(j.Widget.Image.Alignment)

	fmt.Println(string(b))
}

func marshalJson() {
	j := AutoGenerated{}
	j.Widget.Window.Name = "Test1"
	j.Widget.Debug = "false"
	j.Widget.Image.Name = "Test2"
	j.Widget.Window.Width = 30
	j.Widget.Window.Height = 40

	b, err := json.Marshal(j)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(string(b))

	b2, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(string(b2))
}

func readFromFile() {
	s, err := util.File2Array("inputs/demo/file1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	lineNum := 1
	for p, x := range s {
		if x != "" {
			fmt.Printf("line %v: %s (%v)\n", lineNum, x, p)
		}
		lineNum++
	}
}

func readFromFile2() {
	s, err := util.File2Array("inputs/demo/file1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	v := 0
	for _, x := range s {
		if x != "" {
			i, err := strconv.Atoi(x)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			v = v + i
		}
	}
	fmt.Println(v)
	// return strconv.Itoa(v), nil
}

func splitString() {
	s, err := util.File2Array("inputs/demo/file2.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	lineNum := 1
	for _, x := range s {
		data := strings.Split(x, " ")
		if x != "" {
			if len(data) < 2 {
				fmt.Printf("line %v: %s\n", lineNum, data[0])
			} else {
				fmt.Printf("line %v: %s (%s)\n", lineNum, data[0], data[1])
			}
			lineNum++
		}
	}
}

func splitString2() {
	file, err := os.Open("inputs/demo/file2.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		data := strings.Split(t, " ")
		if t != "" {
			if data[0] == "addx" {
				if len(data) == 2 {
					fmt.Println(data[1])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func splitString3() {
	s, err := util.File2Array("inputs/demo/file1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			fmt.Printf("%s (%v) %s\n", x, len(data), data[0])
		}
	}

}

func stringToInt() {
	// b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415", 64)
	// i, err := strconv.ParseInt("-42", 10, 64)
	// u, err := strconv.ParseUint("42", 10, 64)

	file, err := os.Open("inputs/demo/file2.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		data := strings.Split(t, " ")
		if t != "" {
			if len(data) == 2 {
				// i, err := strconv.Atoi(data[1])
				i, err := strconv.ParseInt(data[1], 10, 64)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}

				if i < 0 {
					fmt.Println(i)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func stringToInt2() {
	s, err := util.File2Array("inputs/demo/file2.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, x := range s {
		data := strings.Split(x, " ")
		if x != "" {
			if len(data) == 2 {
				i, err := strconv.ParseInt(data[1], 10, 64)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}

				fmt.Printf("100-(%v) = %v\n", i, 100-(i))
				fmt.Printf("%v-100 = %v\n", i, i-100)
				fmt.Printf("%v+100 = %v\n", i, i+100)
				fmt.Printf("100+(%v) = %v\n", i, 100+(i))
				fmt.Println()
			}
		}
	}
}

// Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
// ANSIC       = "Mon Jan _2 15:04:05 2006"
// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// RFC822      = "02 Jan 06 15:04 MST"
// RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// RFC3339     = "2006-01-02T15:04:05Z07:00"
// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// Kitchen     = "3:04PM"
// // Handy time stamps.
// Stamp      = "Jan _2 15:04:05"
// StampMilli = "Jan _2 15:04:05.000"
// StampMicro = "Jan _2 15:04:05.000000"
// StampNano  = "Jan _2 15:04:05.000000000"
// DateTime   = "2006-01-02 15:04:05"
// DateOnly   = "2006-01-02"
// TimeOnly   = "15:04:05"

func parseTime() {
	t, err := time.Parse(time.RFC3339, "2023-01-01T10:10:11Z")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(t)
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
}
