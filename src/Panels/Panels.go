package panels

type PanelType int

const (
	Status PanelType = iota
	Notification
	QuestInfo
	QuestRewards
)

type Panel struct {
	Kind PanelType
	Title string
	Lines []string
}
