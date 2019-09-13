package template

import (
	"bytes"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	basicInfoHTML = `<!DOCTYPE html>
	<html>
		<head>
			<title>heartbeat</title>
		</head>
		<body>
			<h1>{{ .InfoHeaderName }}</h1>
			<p>Your requested endpoint is available under <a href="{{ .PathName }}">{{ .PathName }}</a> </p>
		</body>
	</html>`
)

// Info struct
type Info struct {
	Info     tmplInfo
	template *template.Template
	result   string
}

type tmplInfo struct {
	InfoHeaderName string
	PathName       string
}

// New returns new template info
func New(infoHeader, path string) (*Info, error) {
	tmpl, err := template.New("html-tmpl").Parse(basicInfoHTML)
	if err != nil {
		return nil, err
	}

	return &Info{
		Info: tmplInfo{
			InfoHeaderName: infoHeader,
			PathName:       path,
		},
		template: tmpl,
	}, nil
}

// GenerateInfoSite template a new info site
func (i *Info) generateInfoSite() error {
	buf := new(bytes.Buffer)
	if err := i.template.Execute(buf, i.Info); err != nil {
		return err
	}

	log.Debugf("Created template: \n%s", buf.String())
	i.result = buf.String()
	return nil

}

// GetTempaltedHandler returns templated handler info site
func (i *Info) GetTempaltedHandler() (http.Handler, error) {

	if err := i.generateInfoSite(); err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(i.result)); err != nil {
			log.Error(err)
		}
	}), nil
}
