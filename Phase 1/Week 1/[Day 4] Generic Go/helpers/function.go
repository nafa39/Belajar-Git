package helpers

// generic struct
func (p *Person[T]) AddScores(val T) {
	p.Scores = append(p.Scores, val)
}

// generic Interface
func (r *Redis[T]) Get() T {
	return r.Value
}

func (r *Redis[T]) Set(newVal T) {
	r.Value = newVal
}
