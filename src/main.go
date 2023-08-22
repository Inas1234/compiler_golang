package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)



func Token_to_Asm(tokens[] Token, node_exit NodeExit) string {
	var buffer bytes.Buffer

	buffer.WriteString("global _start\n _start:\n")

	buffer.WriteString("	mov rax, 60\n")
	buffer.WriteString("	mov rdi, " + node_exit.expr.int_lit.Value + "\n")
	buffer.WriteString("	syscall\n")

	return buffer.String()

}


func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	
	tokens := Tokenize(string(data))
	fmt.Println(tokens)

	var buffer bytes.Buffer
	tree, err := Parse(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(tree)
	buffer.WriteString(Token_to_Asm(tokens, tree))

	os.WriteFile("out.asm", []byte(buffer.Bytes()), 0644)


	// executa a command 
	cmd1 := exec.Command("nasm", "-felf64", "out.asm", "-o", "out.o")
	var stdout, stderr bytes.Buffer
	cmd1.Stdout = &stdout
	cmd1.Stderr = &stderr
	err = cmd1.Run()
	if err != nil {
		log.Fatalf("Failed to run cmd1: %s\nStdout: %s\nStderr: %s", err, stdout.String(), stderr.String())
	}

	

	cmd2 := exec.Command("ld", "out.o", "-o", "out")
	err2 := cmd2.Run()
	if err2 != nil {
		panic(err2)
	}

}