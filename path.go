package main

type Path struct {
	resource 	string
	name 		string
	group 		string
}

func (p *Path) URI(uri string) (string) {
	path := "/manage/v2"
	if p.resource != "" {
		path = path + "/" + p.resource
	}
	if p.name != "" {
		path = path + "/" + p.name
	}
	path = path + "?view=status&format=json"
	if p.group != "" {
		path = path + "&group-id=" + p.group
	}
	return uri + path
}