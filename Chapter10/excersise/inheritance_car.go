package main

import "fmt"

// Engine 引擎
type Engine interface {
	Start()
	Stop()
}

// Car 车
type Car struct {
	wheelCount int
	Engine
}

// Mercedes 梅赛德斯
type Mercedes struct {
	Car
}

func (c *Car) numberOfWheels() int {
	return (*c).wheelCount
}

func (m *Mercedes) sayHiToMerker() {
	fmt.Printf("Hi merker!\n")
}

// Start 启动
func (c *Car) Start() {
	fmt.Println("Car is started")
}

// Stop 停止
func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

// goToWorkIn 去工作
func (c *Car) goToWorkIn() {
	c.Start()
	c.Stop()
}
