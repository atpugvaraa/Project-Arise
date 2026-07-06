package panels

func NewQuestInfoPanel(objective string) Panel {
	return Panel{
		Kind:  QuestInfo,
		Title: "QUEST INFO",
		Lines: []string{
			objective,
		},
	}
}

func NewQuestRewardsPanel(rewards []string) Panel {
	return Panel{
		Kind:  QuestRewards,
		Title: "QUEST REWARDS",
		Lines: append([]string{
			"You have the following rewards!",
		}, rewards...),
	}
}
