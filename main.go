package main
import (
"fmt"
 "github.com/yuin/gopher-lua"
)


func Double(L *lua.LState) int {
    lv := L.ToInt(1)             /* get argument */
    L.Push(lua.LNumber(lv * 2)) /* push result */
    return 1                     /* number of results */
}

func calllua(){
L := lua.NewState()
defer L.Close()
if err := L.DoFile("double.lua"); err != nil {
    panic(err)
}
if err := L.CallByParam(lua.P{
    Fn: L.GetGlobal("double"),
    NRet: 1,
    Protect: true,
    }, lua.LNumber(99)); err != nil {
    panic(err)
}
ret := L.Get(-1) // returned value
L.Pop(1)  // remove received value
fmt.Println(ret)
}

func main(){

 L := lua.NewState()
 defer L.Close()
 L.SetGlobal("double", L.NewFunction(Double)) 

 if err := L.DoString(`print(double(12))`); err != nil {
    panic(err)
 }

 L2 := lua.NewState()
 defer L2.Close()
 if err := L2.DoFile("hello.lua"); err != nil {
     panic(err)
 }

 calllua()

}
