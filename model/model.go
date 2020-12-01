package model

type MessageCard struct {
	CorrelationId string
	Summary string
	Title string
	ThemeColor string
	Text string
	Sections[] Section
}
type Section struct {
	Title string
	Activity Activity
	Facts[] Fact
	PotentialActions[] PotentialAction
}
type Activity struct {
	Title string
	Subtitle string
	Image string
	Text string
}
type Fact struct {
	Name string
	Value string
}
type PotentialAction struct {
	Name string
	Uri string
}
