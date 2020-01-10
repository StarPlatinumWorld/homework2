package main

import "fmt"



type character struct{
	name string
	hp float64
	mp float64
	agg float64

}
var x*float64
var y*float64

func main() {
	A:=character{
		name:"A",
		hp:100,
		mp:0,
		agg:10,
    }
    B:=character{
	    name:"B",
	    hp:300,
	    mp:50,
	    agg:20,

    }
    m:=0.0
    b:=0.0
    x=&m
    y=&b
    for i:=1;i<=5;i++{
    	Round1(A,B)
    }
    for i:=1;i<=5;i++{
        Round2(A,B)
    }

    var Awp float64
    var Bwp float64
    Awp=float64(*x/10.0)
    Bwp=float64(*y/10.0)
    fmt.Println(A.name,"胜率：",Awp)
    fmt.Println(B.name,"胜率：",Bwp)

}
func Random()int{
	ch:=make(chan int,1)
	for{
		select{
		case ch<-0:
		case ch<-1:
		}
		i:=<-ch
		return i
	}
}
func Round1(A character,B character) {
    for {
    	B.hp -=A.agg
    	if Random()==0{
    		B.hp -=A.agg
        }
        A.hp -=B.agg
        B.mp +=10
        if B.mp==50{
        	A.agg=A.agg*0.9
        	B.mp=0
        }
        if B.hp<=0{
            fmt.Println(A.name,"胜利")
            *x+=1
            return
        }else if A.hp<=0{
        	fmt.Println(B.name,"胜利")
        	*y+=1
        	return
        }
    }
}
func Round2(A character,B character){
	for{
		A.hp-=B.agg
		B.mp+=10
		if B.mp==50{
			A.agg=A.agg*0.9
			B.mp=0
		}
		B.hp-=A.agg
		if Random()==0{
			B.hp-=A.agg
		}
		if B.hp<=0{
			fmt.Println(A.name,"胜利")
			*x+=1
			return
		}else if A.hp<=0{
			fmt.Println(B.name,"胜利")
			*y+=1
			return
		}
	}
}

