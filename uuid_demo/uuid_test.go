package uuid_demo

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	u1 := uuid.New()
	u2, _ := uuid.NewRandom()
	fmt.Println(u1.ID())
	fmt.Println(u2.ID())
}
