package main

type Path struct {
	view     string
	resource string
	name     string
	group    string
}

func (p *Path) URI(uri string) string {
	path := "/manage/v2"
	if p.resource != "" {
		path = path + "/" + p.resource
	}
	if p.name != "" {
		path = path + "/" + p.name
	}
	path = path + "?format=json"
	if p.view != "" {
		path = path + "&view=" + p.view
	}
	if p.group != "" {
		path = path + "&group-id=" + p.group
	}
	return uri + path
}
