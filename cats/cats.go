package cats

import "fmt"

type Cat struct {
    //  Cat's name(Required)
    Name    string

    //  Cat's short description
    Short   string

    //  Cat's body
    //  Cat must have Body or Run property
    Body    []string

    //  Cat's function
    //  Cat must have Body or Run property
    Run     func()
}

func (cat Cat) Meow(){
    if len(cat.Body) == 0 {
        cat.Run()
    }else{
        for _, v := range cat.Body {
            fmt.Println(v)
        }
    }
}
