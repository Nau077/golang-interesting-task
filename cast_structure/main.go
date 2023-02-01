package main

type drinker interface {
	drinkVodka()
	drinkNastoika()
}

type drinkerKz interface {
	drinkVodka()
	drinkKumis()
}

type valihan struct{}

func (v *valihan) drinkVodka()    {}
func (v *valihan) drinkNastoika() {}
func (v *valihan) drinkKumis()    {}

func testBody() {
	var a drinker = &valihan{}
	var b drinkerKz = a.(drinkerKz)

	b.drinkKumis()
}

func main() {
	testBody()
}
