package cats

func init(){
    dev := Cat{
        Use: "dev",
        Name: "DevDevにゃんこ",
        Desc: "DevDevにゃんこなのだ...",
        Body: []string{
            "んぅ...DevDevにゃんこなのだ...",
            "             ＿＿",
            "          ／＞    フ",
            "          |    _  _l",
            "         ／` ミ＿xノ",
            "        /          |",
            "       /   ヽ     ﾉ",
            "      |     |  |  |       ＿＿＿＿",
            "  ／￣|     |  | ＿＿＿＿/       /",
            "  | (￣ヽ＿_ヽ_)_|＼  ヽ/       /＼",
            "  ＼二つ         |＼＼  o＝＝＝o   ＼",
            "                    ＼|￣￣￣￣￣￣ |",
            "                      |￣￣￣￣￣￣ |",
        },
    }

    RegisterCat(dev)
}