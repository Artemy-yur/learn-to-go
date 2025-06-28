package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	scan_os()
}

func scan_os() {
	switch runtime.GOOS {
	case "windows":
		os.Create("windows.txt")
		os.WriteFile("windows.txt", []byte("Windows"), 0644)

		dir, err := os.Open("E:\\golang")
		if err != nil {
			fmt.Println("Error opening directory")
		}
		defer dir.Close()
		files, err := dir.Readdir(0)
		if err != nil {
			fmt.Println("Error reading directory")
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		scan_windows()
	case "linux":
		os.Create("linux.txt")
		os.WriteFile("linux.txt", []byte("Linux"), 0644)
		dir, err := os.Open("/home/golang")
		if err != nil {
			fmt.Println("Error opening directory")
		}
		defer dir.Close()
		files, err := dir.Readdir(0)
		if err != nil {
			fmt.Println("Error reading directory")
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		fp, err := os.OpenFile("linux.txt", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file")
		}
		defer fp.Close()
		fp.WriteString("Linux")
		fp.WriteString("\n")
		fp, err = os.OpenFile("linux.txt", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file")
		}
		defer fp.Close()
		scan_linux()


	case "darwin":
		os.Create("macos.txt")
		os.WriteFile("macos.txt", []byte("Mac OS"), 0644)
		dir, err := os.Open("/Users/golang")
		if err != nil {
			fmt.Println("Error opening directory")
		}
		defer dir.Close()
		files, err := dir.Readdir(0)
		if err != nil {
			fmt.Println("Error reading directory")
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		fp, err := os.OpenFile("macos.txt", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file")
		}
		data, err := os.ReadFile("macos.txt")
		if err != nil {
			fmt.Println("Error reading file")
		}
		fmt.Printf("Read file: %s", string(data))
		defer fp.Close()
		scan_macos()
	
	}
}
func scan_windows() {
	
	num_cpu := runtime.NumCPU()
	currentProcs := runtime.GOMAXPROCS(0)
	availableProcs := runtime.GOMAXPROCS(-1)
	
	fp, err := os.OpenFile("windows.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer fp.Close()
	fp.WriteString("\n--------------------------------\n")
	fp.WriteString("Windows\n")
	fp.WriteString(fmt.Sprintf("Time: %s", time.Now().Format("2006-01-02 15:04:05")))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Number of CPU: %d", num_cpu))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Current GOMAXPROCS: %d", currentProcs))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Available GOMAXPROCS: %d", availableProcs))
	fp.WriteString("\n")
	data, err := os.ReadFile("windows.txt")
		if err != nil {
			fmt.Println("Error reading file")
		}
		fmt.Printf("Read file: %s", string(data))
	return
}
func scan_linux() {
	num_cpu := runtime.NumCPU()
	currentProcs := runtime.GOMAXPROCS(0)
	availableProcs := runtime.GOMAXPROCS(-1)
	
	fp, err := os.OpenFile("linux.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer fp.Close()
	fp.WriteString("\n--------------------------------\n")
	fp.WriteString("Linux\n")
	fp.WriteString(fmt.Sprintf("Time: %s", time.Now().Format("2006-01-02 15:04:05")))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Number of CPU: %d", num_cpu))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Current GOMAXPROCS: %d", currentProcs))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Available GOMAXPROCS: %d", availableProcs))
	fp.WriteString("\n")
	data, err := os.ReadFile("windows.txt")
		if err != nil {
			fmt.Println("Error reading file")
		}
		fmt.Printf("Read file: %s", string(data))
	return
}
func scan_macos() {
	num_cpu := runtime.NumCPU()
	currentProcs := runtime.GOMAXPROCS(0)
	availableProcs := runtime.GOMAXPROCS(-1)
	
	fp, err := os.OpenFile("macos.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer fp.Close()
	fp.WriteString("\n--------------------------------\n")
	fp.WriteString("Mac OS\n")
	fp.WriteString(fmt.Sprintf("Time: %s", time.Now().Format("2006-01-02 15:04:05")))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Number of CPU: %d", num_cpu))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Current GOMAXPROCS: %d", currentProcs))
	fp.WriteString("\n")
	fp.WriteString(fmt.Sprintf("Available GOMAXPROCS: %d", availableProcs))
	fp.WriteString("\n")
	data, err := os.ReadFile("windows.txt")
		if err != nil {
			fmt.Println("Error reading file")
		}
		fmt.Printf("Read file: %s", string(data))
	return
}
