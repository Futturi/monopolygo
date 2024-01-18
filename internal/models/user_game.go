package models

type UserGame struct {
	Userid    int
	Username  string
	Money     int
	Position  int
	Business  []Business
	Monopolys []int
}

type Board struct {
	Businesses     []Business
	PlayerPosition []int
}

type Business struct {
	Name             string
	Price            int
	Owner            UserGame
	PriceIfAnotherIn int
	PriceHome        int
	OneHome          int
	TwoHome          int
	ThreeHome        int
	FourHome         int
	Hotel            int
	NumberOfMonopoly int // Каждая монополия - имеет отдельный индекс, у игрока есть массив интов длиной 8, если в массиве под индексом монополии набирается 3( или 2 если 0 или 8 монополия) то игроку разрешается ставить домики
}
