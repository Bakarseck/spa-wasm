package components

import "stack/frameGo"

type Clash struct {
	IdClash      int
	ClashTechno1 string
	ClashTechno2 string
	NbComment    int
	NbVote       int
}

func (c *Clash) Render() string {
	component :=
		`
	<div class="clash-created" id="clash-{{.IdClash}}">
	<h2>
		{{.ClashTechno1}} vs {{.ClashTechno2}}
	</h2>
	<div class="progression-bar">
	  <div class="progress-bar progress-bar1 percentage"><p id="progress">70.0%</p></div>
	  <div class="progress-bar progress-bar2 percentage"><p id="progress">30.0%</p></div>
	</div>
	<div class="techno-clash">
	<p>{{.ClashTechno1}}</p>
	<p>{{.ClashTechno2}}</p>
	</div>
	<div class="stats">
	  <div class="comment">
		<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1"><path d="m3 21 1.9-5.7a8.5 8.5 0 1 1 3.8 3.8z"></path></svg>
		{{.NbComment}} comments
	  </div>
	  <div class="vote">
		<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1"><path d="m16.02 12 5.48 3.13a1 1 0 0 1 0 1.74L13 21.74a2 2 0 0 1-2 0l-8.5-4.87a1 1 0 0 1 0-1.74L7.98 12"></path><path d="M13 13.74a2 2 0 0 1-2 0L2.5 8.87a1 1 0 0 1 0-1.74L11 2.26a2 2 0 0 1 2 0l8.5 4.87a1 1 0 0 1 0 1.74Z"></path></svg>
		{{.NbVote}} vote
	  </div>
	</div>
	</div>
	`
	return frameGo.ParseComponent(component, &c)
}

func NewClash(id, comment, vote int, techno1, techno2 string) Clash {
	return Clash{
		IdClash:      id,
		ClashTechno1: techno1,
		ClashTechno2: techno2,
		NbComment:    comment,
		NbVote:       vote,
	}
}