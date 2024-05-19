{{- define "mainpage" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>

	<div align="center">
		<textarea rows="80" cols="100" readonly="true">{{.ResultStringData}}</textarea>
	</div>

	{{ if eq .Layout "none" }}
    {{- range .Charts }} {{ template "base" . }} {{- end }}
	{{ end }}

	{{ if eq .Layout "center" }}
    <style> .container {display: flex;justify-content: center;align-items: center;} .item {margin: auto;} </style>
    {{- range .Charts }} {{ template "base" . }} {{- end }}
	{{ end }}

	{{ if eq .Layout "flex" }}
    <style> .box { justify-content:center; display:flex; flex-wrap:wrap } </style>
    <div class="box"> {{- range .Charts }} {{ template "base" . }} {{- end }} </div>
	{{ end }}
	
</body>
</html>
{{ end }}
