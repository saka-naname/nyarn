package cats

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Cat struct {
    //  Cat's name as subcommand(Required)
    Use     string

    //  Cat's name
    Name    string

    //  Cat's long description
    Desc    string

    //  Cat's body
    //  Cat must have Body or Run property
    Body    []string

    //  Cat's function
    //  Cat must have Body or Run property
    Run     func()
}

func (cat Cat) Meow(cmd *cobra.Command, args []string){
    if len(cat.Body) == 0 {
        cat.Run()
    }else{
        for _, v := range cat.Body {
            fmt.Println(v)
        }
    }
}

var cathouse []Cat = []Cat{}

func RegisterCat (cat Cat){
    cathouse = append(cathouse, cat)
}

func GetCatHouse () []Cat {
    return cathouse
}

func InitCats (root *cobra.Command, cats []Cat){
    for _, cat := range cats {
        cm := &cobra.Command{
            Use:    cat.Use,
            Short:  "Print " + cat.Name,
            Long:   cat.Desc,
            Run:    cat.Meow,
        }

        root.AddCommand(cm)
    }
}