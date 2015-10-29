package ccache

import (
	. "github.com/karlseguin/expect"
	"testing"
	"time"
)

type ItemTests struct{}

func Test_Item(t *testing.T) {
	Expectify(new(ItemTests), t)
}

func (_ *ItemTests) Promotability() {
	item := &Item{promotions: 4}
	Expect(item.shouldPromote(5)).To.Equal(true)
	Expect(item.shouldPromote(5)).To.Equal(false)
}

func (_ *ItemTests) Expired() {
	now := time.Now().UnixNano() / 1e6
	item1 := &Item{expires: now + 1}
	item2 := &Item{expires: now - 1}
	Expect(item1.Expired()).To.Equal(false)
	Expect(item2.Expired()).To.Equal(true)
}

func (_ *ItemTests) TTL() {
	now := time.Now().UnixNano() / 1e6
	item1 := &Item{expires: now + 10*1e3}
	item2 := &Item{expires: now - 10*1e3}
	Expect(item1.TTL()).To.Equal(time.Second * 10)
	Expect(item2.TTL()).To.Equal(time.Second * -10)
}

func (_ *ItemTests) PTTL() {
	now := time.Now().UnixNano() / 1e6
	item1 := &Item{expires: now + 10}
	item2 := &Item{expires: now - 10}
	Expect(item1.PTTL()).To.Equal(time.Millisecond * 10)
	Expect(item2.PTTL()).To.Equal(time.Millisecond * -10)
}

func (_ *ItemTests) Expires() {
	now := time.Now().UnixNano() / 1e6
	item := &Item{expires: now + 10}
	Expect(item.Expires().UnixNano() / 1e6).To.Equal(now + 10)
}

func (_ *ItemTests) Extend() {
	item := &Item{expires: time.Now().Unix() + 10}
	item.Extend(time.Minute * 2)
	Expect(item.Expires().Unix()).To.Equal(time.Now().Unix() + 120)
}
