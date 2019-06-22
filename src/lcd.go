package main

import (
        "github.com/davecheney/i2c"
        "flag"
        "log"
        "fmt"
        "time"
        s "strings"
)

func check(err error) {
        if err != nil { log.Fatal(err) }
}

func main() {
        i, err := i2c.New(0x27, 1)
        check(err)
        lcd, err := i2c.NewLcd(i, 2, 1, 0, 4, 5, 6, 7, 3)
        check(err)
        lcd.BacklightOn()
        lcd.Clear()

        // get text from argument -text
        var text string
        flag.StringVar(&text, "text", "Welcome to workshop", "insert text to display to lcd")
        flag.Parse()
        count, txt := len(text), s.Split(text, "")

        for {
                for i := 0; i <= count; i++ {
                        lcd.Home()
                        t := time.Now()
                        new_text := s.Join(txt[i:], "") + s.Repeat(" ", i)
                        lcd.SetPosition(1, 0)
                        fmt.Fprint(lcd, new_text)

                        lcd.SetPosition(2, 0)
                        fmt.Fprint(lcd, t.Format("15:04:05 2006"))

                        time.Sleep(333 * time.Millisecond)
                }
        }
}