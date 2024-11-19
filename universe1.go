package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/cweel/zhouyi"
	"github.com/nosixtools/solarlunar"
)

func main() {
	yi := zhouyi.Text()
	m, n, l, f := universe()
	a := zhouyi.ReGuaNu(m, n, l, f, yi)
	writeToFile(f, a, n, yi)
	writeToTerminal(a, n, yi)
	f.Close()
}

func universe() (m, n, l []uint, f *os.File) {
	m = []uint{0, 0, 0, 0, 0, 0}
	n = []uint{0, 0, 0, 0, 0, 0}
	l = []uint{0, 0, 0, 0, 0, 0}
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(now.Day())
	if len(day) == 1 {
		day = "0" + day
	}
	date := year + "-" + month + "-" + day
	universeTime := solarlunar.SolarToChineseLuanr(date)
	r := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < 6; i++ {
		a := r.Intn(2) + 2
		b := r.Intn(2) + 2
		c := r.Intn(2) + 2
		m[i] = uint(a + b + c)
		if m[i] == 6 || m[i] == 9 {
			n[i] = m[i]
		}
		if m[i] == 7 || m[i] == 9 {
			l[i] = 1
		}
	}

	//输出文本文件
	f, err := os.OpenFile("./"+date, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error: %v\n,creating...\n", err)
		f, err = os.Create("./" + date)
	}
	//defer f.Close()

	fmt.Println(m)
	appendToFile(f, fmt.Sprintln(m))
	fmt.Println(n)
	appendToFile(f, fmt.Sprintln(n))
	fmt.Print(l, " ", date, " ", now.Hour(), ":", now.Minute(), "  ", universeTime)
	appendToFile(f, fmt.Sprint(l, " ", date, " ", now.Hour(), ":", now.Minute(), "  ", universeTime))

	return
}
func writeToTerminal(a uint, n []uint, yijing []zhouyi.Gua) {
	fmt.Printf(" 第%d卦\n", a)
	fmt.Println(yijing[a-1].Tuan)
	//动爻 判断
	for i := 0; i < 6; i++ {
		fmt.Println(yijing[a-1].Xi[i])
	}
	fmt.Println("\n<－－动爻－－>")
	for i := 0; i < 6; i++ {
		if n[i] != 0 {
			fmt.Println(yijing[a-1].Xi[i])
		}
	}
}
func writeToFile(f *os.File, a uint, n []uint, yijing []zhouyi.Gua) {

	appendToFile(f, fmt.Sprintf(" 第%d卦\n", a))
	appendToFile(f, fmt.Sprintln(yijing[a-1].Tuan))
	//动爻 判断
	for i := 0; i < 6; i++ {
		appendToFile(f, fmt.Sprintln(yijing[a-1].Xi[i]))
	}
	appendToFile(f, fmt.Sprintln("\n<－－动爻－－>"))
	for i := 0; i < 6; i++ {
		if n[i] != 0 {
			appendToFile(f, fmt.Sprintln(yijing[a-1].Xi[i]))
		}
	}

}

func appendToFile(f *os.File, content string) {
	// 以只写的模式，打开文件

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	f.WriteAt([]byte(content), n)

}
