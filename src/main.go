package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	repo := flag.String("repo", ".", "путь к Git-репозиторию")
	days := flag.Int("days", 7, "количество дней для анализа")
	flag.Parse()

	cmd := exec.Command("git", "-C", *repo, "log",
		"--since="+time.Now().AddDate(0,0, -*days).Format("2006-01-02"),
		"--pretty=format:%ad", "--date=short")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Ошибка вызова git:", err)
		return
	}

	commitsPerDay := map[string]int{}
	for _, date := range strings.Split(string(out), "\n") {
		if date != "" {
			commitsPerDay[date]++
		}
	}

	// Вывод ASCII-графика
	for i := 0; i < *days; i++ {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		count := commitsPerDay[day]
		bar := strings.Repeat("#", count)
		fmt.Printf("%s | %s (%d)\n", day, bar, count)
	}
}
