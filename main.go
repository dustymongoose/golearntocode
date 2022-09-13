package main

import
(
"fmt"
"os/exec"
"log"
"os"
)

func main() {
	file, err := os.Create("notes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
        fmt.Println("file created")
	defer file.Close()

	if err := os.Chdir("test"); err != nil {
		log.Fatal(err)
	}

	file2, err := os.Create("notes2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
        fmt.Println("file created")
	defer file2.Close()

	if err := os.Chmod("notes2.txt", 0644); err != nil {
		log.Fatal(err)
	}
	
	if err := os.Chown("notes2.txt", 1000, 1000); err != nil {
		log.Fatal(err)
	}
	
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hostname)

	cmd := exec.Command("dnf", "install", "-y", "nano")
	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v. Package successfully installed", err)

	cmd2 := exec.Command("dnf", "remove", "-y", "nano")
	log.Printf("Running command and waiting for it to finish...")
	err = cmd2.Run()
	log.Printf("Command finished with error: %v. Package successfully removed", err)

}
