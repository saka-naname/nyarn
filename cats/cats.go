package cats

import "fmt"

type Cat struct {
    //  Cat's name(Required)
    Name    string

    //  Cat's body
    Body    []string

    //  Cat's function
    Run     func()
}

func Meow(cat Cat){
    if len(cat.Body) == 0 {
        cat.Run()
    }else{
        for _, v := range cat.Body {
            fmt.Println(v)
        }
    }
}
