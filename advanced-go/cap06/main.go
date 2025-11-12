package main

import "fmt"

//buffers
func main(){
	buf := NewBuffer()

	go func (){
		buf.mu.Lock()
		for len(buf.item) == 0{
			buf.cond.Wait()
		}
		items := buf.item[0]
		buf.item = buf.item[1:]
		buf.mu.Unlock()
		fmt.Printf("Consumido:" , items)
	}()
}

