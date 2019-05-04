package alcoholics

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Option struct {
	Probability  int
	StatusCode   int
	ErrorMessage string
	isError      bool
}

func shuffle(data []Option) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func New(os []Option) *Drunker {
	ret := &Drunker{
		o: []Option{},
	}

	for _, o := range os {
		if o.StatusCode == 200 {
			o.isError = false
		} else {
			o.isError = true
		}

		for i := 0; i < o.Probability; i++ {
			ret.o = append(ret.o, o)
		}
	}

	shuffle(ret.o)

	return ret
}

type Drunker struct {
	o []Option
}

func (d *Drunker) Drunk() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idx := rand.Int() % len(d.o)
		fmt.Println(len(d.o), idx)
		o := d.o[idx]
		if o.isError {
			ctx.JSON(o.StatusCode, o.ErrorMessage)
			ctx.Abort()
		}
		ctx.Next()
	}
}
