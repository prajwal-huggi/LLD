package main
import(
    "fmt"
)

type Coffee interface{
    Cost() int
    Description() string
}

type Plain struct{}

func(Plain) Cost() int{
    return 50
}

func(Plain) Description() string{
    return "Plain "
}

type Milk struct{
    base Coffee
}

func(m Milk) Cost() int{
    initialCost:= m.base.Cost()
    return initialCost+ 20
}

func(m Milk) Description() string{
    return m.base.Description()+" + Milk"
}

type Sugar struct{
    base Coffee
}

func(s Sugar) Cost() int{
    return s.base.Cost()+ 10
}

func(s Sugar) Description() string{
    return s.base.Description()+ " + sugar"
}

func main(){
    // c:= Plain{}
    // c1:= Milk{base: c}
    // c2:= Sugar{base: c}
    // c3:= Sugar{base: c1}
    
    c4:= Sugar{base: Milk{base: Plain{}}}
    fmt.Println("meetha dudh: ", c4.Cost(), " ", c4.Description())
    
    // fmt.Println("Plain coffee: ", c.Cost(), " ", c.Description())
    // fmt.Println("Milk coffee: ", c1.Cost(), " ", c1.Description())
    // fmt.Println("Sugar coffee: ", c2.Cost(), " ", c2.Description())
    // fmt.Println("Milk+ sugar coffee: ", c3.Cost(), " ", c3.Description())
}

