package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	type args struct {
		h   *SliceQueue
		run func()
	}
	type want struct {
		x interface{}
	}
	q1 := SliceQueue{}
	q2 := SliceQueue{}
	q3 := SliceQueue{}
	q4 := SliceQueue{}

	testcases := map[string]struct {
		args args
		want want
	}{
		"nil SliceQueue": {
			args: args{},
			want: want{
				x: nil,
			},
		},
		"empty SliceQueue": {
			args: args{
				h: &q1,
			},
			want: want{
				x: nil,
			},
		},
		"one element SliceQueue": {
			args: args{
				h: &q2,
				run: func() {
					q2.Push("a")
				},
			},
			want: want{
				x: "a",
			},
		},
		"two elements SliceQueue": {
			args: args{
				h: &q3,
				run: func() {
					q3.Push("a")
					q3.Push("b")
				},
			},
			want: want{
				x: "a",
			},
		},
		"11 elements SliceQueue": {
			args: args{
				h: &q4,
				run: func() {
					for i := 0; i < 11; i++ {
						q4.Push(i)
					}
					err := q4.Push(11)
					assert.Contains(t, err.Error(), "queue is full")
				},
			},
			want: want{
				x: 0,
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if tc.args.run != nil {
				tc.args.run()
			}
			got := tc.args.h.Pop()
			if !reflect.DeepEqual(got, tc.want.x) {
				t.Errorf("Pop() = %v, want %v", got, tc.want.x)
			}
		})
	}
}

func TestGorountineSafe(t *testing.T) {
	q := SliceQueue{}
	for i := 0; i < 10; i++ {
		go q.Push(i)
	}
	time.Sleep(5 * time.Second)
	t.Log("q.Len() = ", q.Len())
	for {
		t.Log(q.Pop())
		if q.Len() == 0 {
			break
		}
	}
	fmt.Println()
}
