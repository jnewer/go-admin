package nuts

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
	"testing"
	"time"
)

func init() {
	InitNuts()
}

func TestInitNuts(t *testing.T) {
	err := Instance().nuts.Update(func(t *nutsdb.Tx) error {
		key := []byte("name")
		val := []byte("val")
		bucket := "table1"
		if err := t.Put(bucket, key, val, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNutsDB(t *testing.T){
	err := Instance().Set("thiskey","123")
	if err != nil{
		log.Fatal(err.Error())
	}
	Instance().Set("abc","aaa",10)
	fmt.Println(Instance().Get("abc"))
	time.Sleep(time.Second * 10)
	fmt.Println(Instance().Get("abc"))

}