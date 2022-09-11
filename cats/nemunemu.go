package cats

func init(){
    nemunemu := Cat{
        Use: "nemunemu",
        Name: "ねむねむにゃんこ",
        Desc: "ねむねむにゃんこなのだ...",
        Body: []string{
            "んぅ...ねむねむにゃんこなのだ...",
            "             ＿＿",
            "          ／＞    フ",
            "          |    _  _l",
            "         ／` ミ＿xノ",
            "        /          |",
            "       /   ヽ     ﾉ",
            "      │     |  |  |",
            "  ／￣|     |  |  |",
            "  | (￣ヽ＿_ヽ_)__)",
            "  ＼二つ",
        },
    }

    RegisterCat(nemunemu)
}