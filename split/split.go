package main

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"syscall/js"
	"time"
)

var ch chan int

func consoleLog(s string) {
	js.Global().Get("console").Call("log", s)
	js.Global().Get("document").Call("getElementById", "text").Set("value", s)
}
func registerFunc(fname string, fn func(js.Value, []js.Value) interface{}) {
	js.Global().Get("document").Set(fname, js.ValueOf(js.FuncOf(fn)))
}
func getElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}
func main() {
	consoleLog(fmt.Sprintf("split/wasm loaded! %s", time.Now().String()))
	registerFunc("submitMsg", submitMsg)
	registerFunc("nextSection", nextSection)
	go ttsLoop()
	getElementById("text").Set("value", "split/wasm loaded!")
	ch = make(chan int, 1)
	ch <- 1
	getElementById("hecheng").Set("disabled", "")
	select {}
}

func nextSection(p js.Value, val []js.Value) interface{} {
	ch <- 1
	return nil
}

func submitMsg(p js.Value, val []js.Value) interface{} {
	entry1 := getElementById("text")
	msg := entry1.Get("value").String()

	split(msg)

	return nil
}

func split(s string) {
	const size1 = 2000
	src := bytes.NewBufferString(s)

	reader1 := bufio.NewReader(src)

	buf := bytes.NewBufferString("")

	secSize := 0
	num := 1
	regex1 := regexp.MustCompile("\\s+")
	for {
		line1, _, err := reader1.ReadLine()
		if err != nil {
			if secSize > 0 {
				saveSection(buf.Bytes(), num)
				buf.Reset()
			}
			break
		}
		line1 = bytes.TrimSpace(line1)
		line1 = regex1.ReplaceAll(line1, []byte("!"))
		runes1 := bytes.Runes(line1)
		lsize := len(runes1)
		if lsize == 0 {
			continue
		}

		for (secSize + lsize) > size1 {
			pos := getBreakPos(runes1, size1-secSize)
			for n := 0; n <= pos; n++ {
				buf.WriteRune(runes1[n])
			}
			runes1 = runes1[pos+1:]
			saveSection(buf.Bytes(), num)
			buf.Reset()
			secSize = 0
			lsize = len(runes1)
			num++
		}

		if lsize > 0 {
			for i := 0; i < lsize; i++ {
				buf.WriteRune(runes1[i])
			}
		}
		buf.WriteByte('!')
		secSize += lsize + 1
	}
}

var msgs chan string

func ttsLoop() {
	msgs = make(chan string, 100)
	for {
		<-ch
		data := <-msgs
		js.Global().Call("ttsWasm", data)

	}
}
func saveSection(data []byte, num int) {
	buf := make([]byte, len(data))
	copy(buf, data)
	msgs <- string(buf)
}

func getBreakPos(data []rune, max int) (pos int) {
	const ends1 = `。！？?!.`
	runes1 := bytes.Runes([]byte(ends1))
	pos = max
	for {
		if pos == -1 {
			return
		}
		for _, r := range runes1 {
			if r == data[pos] {
				return
			}
		}

		pos--
	}
	return
}
