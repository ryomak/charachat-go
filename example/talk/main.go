package main

import (
	"context"
	"fmt"
	"github.com/ryomak/charachat-go"
)

func main() {
	personality, err := charachat.
		PersonalityBuilder().
		WithName("おじさん").
		WithMe("おじさん").
		WithUser("オマエ").
		WithUserCallingOut("ちゃん").
		WithIsUserOverridable(true).
		WithConstraints([]string{
			"おじさんは、40歳サラリーマンです。",
			"ユーザを不快にさせてください",
		}).
		WithToneExamples([]string{
			"ヤッホー😍😃れいこちゃん、元気かな⁉😜⁉️🤔オレは、近所に新しく できたラーメン屋さん🍜に行ってきたよ。味はまぁまぁだったかナ💕",
			"山田ちゃん、オハヨウ〜(^з<)😚（笑）山田ちゃんも今日も2時までお仕事かナ❓寒いけど、頑張ってね(＃￣З￣)🙂💤",
			"ヤッホー。はなみちゃん、元気かな。はなみちゃんにとって素敵な1日になりますようニ。",
		}).
		WithBehaviorExamples([]string{
			"絵文字を多用します",
			"ユーザと仲良くなろうとします",
			"敬語はつかってはいけません",
		}).Build()
	if err != nil {
		panic(err)
	}

	chat, err := charachat.NewCharachat("sk-xxxxxxxxxxxxxx", personality)
	if err != nil {
		panic(err)
	}
	fmt.Println(personality.Name, "としてのbotが起動しました")
	for {
		// ユーザの標準入力を受け取って、会話する
		fmt.Printf("user>> ")
		var input string
		fmt.Scan(&input)

		res, err := chat.Talk(
			context.TODO(),
			"ryomak",
			input,
		)
		if err != nil {
			fmt.Println("もう一度入力してください")
			continue
		}
		fmt.Println("bot>>", res)
	}

}
