//lint: file-ignore ST1006

package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (self *SpecialPosition) SpecialMove(x,y float64){
	self.x += (x * x);
	self.y += (y * y);
}

type Position struct {
	x float64
	y float64
}

func (self *Position) Move(x float64, y float64){
	self.x += x;
	self.y += y;
}

func (self *Position) Teleport(x float64, y float64){
	self.x = x;
	self.y = y;
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}


func main(){
	boss := NewEnemy();
	boss.Move(1.1, 10.4);
	fmt.Println("boss", boss.Position);
	boss.SpecialMove(42.0, 77.0);
	fmt.Println("boss", boss.Position);

	player := NewPlayer();
	player.Move(1.1, 10.0);
	fmt.Println(player.Position);
	player.Teleport(100.5, 200.8);
	fmt.Println(player.Position);
}