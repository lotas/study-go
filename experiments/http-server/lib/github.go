package lib

import (
	"fmt"
	"io"

	"gopl.io/ch4/github"
)

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

// SearchGithub searches issues
func SearchGithub(args []string, wr io.Writer) {
	result, err := github.SearchIssues(args)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(wr, "%s", err)
		return
	}
	if err := issueList.Execute(wr, result); err != nil {
		fmt.Println(err)
		fmt.Fprintf(wr, "%s", err)
	}
}
