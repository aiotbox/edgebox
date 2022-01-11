package process

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Process struct {
	name string
	Cmd string
	Args []string
	c* exec.Cmd
	pid int
}
func NewProcess(name string,cmd string,args string) *Process{
	p :=Process{}
	p.name = name
	p.Cmd = cmd
	p.Args = strings.Split(args,",")
	p.pid = -1
	return &p
}
func (p* Process)Print(){
	fmt.Println("Name:",p.name,",Cmd:",p.Cmd,",Args:",p.Args,",Pid=",p.Pid())
}
func (p* Process)Pid() int{
	cp:=p.pid
	return cp
}
func (p* Process)Start(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	if p.pid >= 0{
		return
	}
	ctx, _:= context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx,p.Cmd, p.Args...)
	cmd.Stdout =os.Stdout
	cmd.Start()
	p.pid = cmd.Process.Pid
	p.c = cmd
	time.Sleep(time.Second)
	cmd.Wait()
	//defer cancel()
}
func (p* Process)Stop(){
	if p.c != nil{
		p.c.Process.Signal(os.Interrupt)
		p.pid = -1
	}
}
func (p* Process)Kill(){
	if p.c != nil{
		p.c.Process.Kill()
		p.pid = -1
	}
}