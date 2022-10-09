package encapsulation

// 定义一个结构体
type student struct {
	Name  string
	score float64
}

// *student 返回结构体的指针
func NewStudent(name string, score float64) *student {
	// 外包无法访问结构体的字段，只能通过工厂模式来访问
	return &student{
		Name:  name,
		score: score,
	}
}

// 结构体中的score是小写需要一个方法来访问
func (s *student) GetScore() float64 {
	return s.score
}

func (s *student) SetScore(score float64) {
	s.score = score
}
