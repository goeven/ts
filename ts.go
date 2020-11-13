package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	durationFlag     bool
	jsonFlag         bool
	formatFlag       string
	timestampKeyFlag string
	textKeyFlag      string
	precisionFlag    int
	widthFlag        int
)

func init() {
	flag.BoolVar(&durationFlag, "d", false, "annotate with the elapsed time since the previous line instead of an absolute timestamp")
	flag.BoolVar(&jsonFlag, "json", false, "output each line as a JSON object")
	flag.StringVar(&formatFlag, "fmt", "", "format for the timestamp; ignored when -d is specified")
	flag.StringVar(&timestampKeyFlag, "timestamp-key", "timestamp", "JSON key for the timestamp value; only when -json is specified")
	flag.StringVar(&textKeyFlag, "text-key", "text", "JSON key for a line; only when -json is specified")
	flag.IntVar(&precisionFlag, "p", 1, "precision for the duration; a multiple of milliseconds")
	flag.IntVar(&widthFlag, "w", 5, "width of the timestamp; ignored when -json is set")
}

func main() {
	flag.Parse()

	formatter := Formatter{
		duration:     durationFlag,
		json:         jsonFlag,
		format:       formatFlag,
		timestampKey: timestampKeyFlag,
		textKey:      textKeyFlag,
		precision:    precisionFlag,
		width:        widthFlag,
	}

	scanner := bufio.NewScanner(os.Stdin)
	prev := time.Now()
	for scanner.Scan() {
		line := scanner.Text()

		now := time.Now()

		out, err := formatter.Format(line, prev, now)
		if err != nil {
			log.Fatal(err)
		}

		prev = now

		fmt.Println(out)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Formatter struct {
	duration     bool
	json         bool
	format       string
	timestampKey string
	textKey      string
	precision    int
	width        int
}

func (f Formatter) Format(text string, prev, now time.Time) (string, error) {
	ts := f.ts(prev, now)
	if !f.json {
		return fmt.Sprintf("%*v %s", f.width, ts, text), nil
	}

	data, err := json.Marshal(map[string]interface{}{
		f.timestampKey: ts,
		f.textKey:      text,
	})
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (f Formatter) ts(prev, now time.Time) interface{} {
	if f.duration {
		dur := now.Sub(prev)
		prec := time.Duration(f.precision) * time.Millisecond

		return dur.Truncate(prec).String()
	}

	if f.format != "" {
		return now.Format(f.format)
	}

	return now.Unix()
}
