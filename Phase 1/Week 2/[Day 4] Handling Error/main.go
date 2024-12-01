// handling error in golang
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

//review materi

// // custom error
// func validPassword(password string) (string, error) {
// 	if len(password) < 8 {
// 		return "ada error", errors.New("password harus lebih dari 8 karakter")
// 	}
// 	return "password valid", nil
// }

// // panic
// func validPassword(password string) (string, error) {
// 	if len(password) < 8 {
// 		return "ada error", errors.New("password harus lebih dari 8 karakter")
// 	}
// 	return "password valid", nil
// }

// // recover
// func validPassword(password string) (string, error) {
// 	if len(password) < 8 {
// 		return "ada error", errors.New("password harus lebih dari 8 karakter")
// 	}
// 	return "password valid", nil
// }

// func catchErr() {
// 	if r := recover(); r != nil {
// 		fmt.Println("error terjadi", r)
// 	} else {
// 		fmt.Println("tidak ada error")
// 	}
// }

// // graceful exit with panic recovery
// func catchErr() {
// 	if r := recover(); r != nil {
// 		fmt.Println(r)
// 	} else {
// 		fmt.Println("\ntidak ada error")
// 	}
// }

// func divide(a, b int) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("divide by zero occured")
// 	}
// 	result := float64(a) / float64(b)
// 	return result, nil
// }

// // CLI task list with error handling
// type Task struct {
// 	ID        int
// 	Title     string
// 	Completed bool
// }

// type TaskList struct {
// 	task []*Task
// }

// func (t *TaskList) addTask(title string) {
// 	newTask := &Task{
// 		ID:        len(t.task) + 1,
// 		Title:     title,
// 		Completed: false,
// 	}
// 	t.task = append(t.task, newTask)
// }

// func (t *TaskList) listTask() {
// 	if len(t.task) == 0 {
// 		fmt.Println("Task list is empty")
// 		return
// 	}
// 	for _, task := range t.task {
// 		status := "incomplete"
// 		if task.Completed {
// 			status = "completed"
// 		}
// 		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, status)
// 	}
// }

// func (t *TaskList) removeTask(id int) error {
// 	if id < 1 || id > len(t.task) {
// 		return fmt.Errorf("invalid task ID")
// 	}
// 	t.task = append(t.task[:id-1], t.task[id:]...)

// 	for i, task := range t.task {
// 		task.ID = i + 1
// 	}
// 	return nil
// }

// func (t *TaskList) Load(filname string) error {
// 	file, err := os.Open(filname)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return nil
// 		}
// 		return err
// 	}
// 	defer file.Close()

// 	t.task = []*Task{}

// 	// read file
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		//fmt.Println("read line:", line)

// 		var task Task
// 		var completed bool
// 		var id int
// 		_, err := fmt.Sscanf(line, "%d|%q|%t\n", &id, &task.Title, &completed)
// 		if err != nil {
// 			return fmt.Errorf("failed to read task: %v", err)
// 		}
// 		t.task = append(t.task, &Task{
// 			ID:        id,
// 			Title:     task.Title,
// 			Completed: completed,
// 		})
// 	}
// 	return scanner.Err()
// }

// func (t *TaskList) Save(filname string) error {
// 	file, err := os.Create(filname)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// write to file
// 	for _, task := range t.task {
// 		line := fmt.Sprintf("%d|\"%s\"|%t\n", task.ID, task.Title, task.Completed)
// 		_, err = file.WriteString(line)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// CLI Log Analyzer
type Log struct {
	logFilePath string
}

func newLog(logFilePath string) *Log {
	return &Log{logFilePath: logFilePath}
}

func (l *Log) analyze() error {
	// open log file
	file, err := os.Open(l.logFilePath)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	logLevelCount := map[string]int{}

	// read log file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		level := extractLogLevel(line)
		logLevelCount[level]++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read log file: %v", err)
	}

	fmt.Println("Log analysis result:")
	for level, count := range logLevelCount {
		fmt.Printf("[%s] level : %d occurences\n", level, count)
	}

	return nil
}

func extractLogLevel(line string) string {
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	if start == -1 || end == -1 {
		return "unknown"
	}

	return line[start+1 : end]
}

//INCLASS
// type MyError struct {
// 	Msg string
// 	Err error
// 	Code int
// }

// func insertData(nama string, umur int) (inserted_id int, err error) {
// 	// insert data to database
// 	// if success return inserted_id
// 	// if failed return error
// 	fmt.Println(nama, umur)
// 	return 0, errors.New("Error") // ini belom ngerti
// }

func main() {

	//review materi
	// //error
	// number, err := strconv.Atoi("123GH")
	// if err != nil {
	// 	fmt.Println("error apa", err)
	// } else {
	// 	fmt.Println(number)
	// }

	// number, err = strconv.Atoi("123")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(number)
	// }

	// //custom error
	// var password string
	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	fmt.Println(valid, err)
	// } else {
	// 	fmt.Println(valid)
	// }

	// //panic
	// var password string
	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(valid)
	// }

	// //recover
	// defer catchErr()
	// var password string
	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(valid)
	// }

	// //graceful exit with panic recovery
	// defer catchErr()
	// var a, b int
	// fmt.Scanln(&a)
	// fmt.Scanln(&b)

	// if result, err := divide(a, b); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("result of %.f / %.f is %.2f", float64(a), float64(b), result)
	// }

	// // CLI task list with error handling
	// const filename = "tasks.txt"
	// tasklist := &TaskList{}

	// err := tasklist.Load(filename)
	// if err != nil {
	// 	fmt.Printf("failed to load task list: %v\n", err)
	// 	os.Exit(1)
	// }

	// addFlag := flag.Bool("add", false, "add new task")
	// listFlag := flag.Bool("list", false, "list all tasks")
	// removeFlag := flag.Int("remove", 0, "remove task by ID")
	// flag.Parse()

	// if *addFlag {
	// 	args := flag.Args()
	// 	if len(args) == 0 {
	// 		fmt.Println("please provide task title")
	// 		os.Exit(1)
	// 	}
	// 	title := args[0]
	// 	tasklist.addTask(title)
	// 	err := tasklist.Save(filename)
	// 	if err != nil {
	// 		fmt.Printf("failed to save task list: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println("task added")
	// } else if *listFlag {
	// 	tasklist.listTask()
	// } else if *removeFlag != 0 {
	// 	err := tasklist.removeTask(*removeFlag)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	err = tasklist.Save(filename)
	// 	if err != nil {
	// 		fmt.Printf("failed to save task list: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println("task removed")
	// } else {
	// 	fmt.Println("invalid command. Usage:")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// CLI Log Analyzer
	logFilePath := flag.String("log", "", "path to log file")
	flag.Parse()

	if *logFilePath == "" {
		fmt.Println("please provide log file path")
		os.Exit(1)
	}

	log := newLog(*logFilePath)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
		}
	}()

	err := log.analyze()
	if err != nil {
		fmt.Println("failed to analyze log:", err)
		os.Exit(1)
	}

	fmt.Println("log analysis completed")

	//INCLASS
	// var str1 string = "123"
	// var str2 string = "123.123"

	// var int1 int
	// var int2 int

	// var err error

	// _, err = insertData("John Doe", 20)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Convert string to int
	// int1, err = strconv.Atoi(str1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// } else {
	// 	// 	fmt.Println(int1)
	// }

	// // Convert string to int
	// int2, err = strconv.Atoi(str2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// } else {
	// 	// 	fmt.Println(int2)
	// }

	// fmt.Println(int1)
	// fmt.Println(int2)

}

//case kak dharma
// 	package main
// import (
//     "errors"
//     "fmt"
// )
// func division(n1, n2 int) (int, error) {
//     //validasi input kalo 0
//     if n2 == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     //bagiin
//     result := n1 / n2
//     //balikin result
//     return result, nil
// }
// func main() {
//     n1 := 1
//     n2 := 0
//     //manggil sama bikin variable error
//     _, err := division(n1, n2)
//     //kalo error ngapain kalo success ngapain
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Success")
//     }
// }

//contoh kak dharma lagi
// package main
// import (
//     "errors"
//     "fmt"
// )
// func division(n1, n2 int) (int, error) {
//     //validasi input kalo 0
//     if n2 == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     //bagiin
//     result := n1 / n2
//     //balikin result
//     return result, nil
// }
// func main() {
//     n1 := 1
//     n2 := 0
//     //manggil sama bikin variable error
//     _, err := division(n1, n2)
//     //kalo error ngapain kalo success ngapain
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Success")
//     }
// }

//recover belom aman, defer belom ngerti, fmt.Errorf belom ngerti
//defer func() // kalo panic tetep jalan tapi ga masuk ke recover
//if recover() != nil { => ini buat ngambil errornya
