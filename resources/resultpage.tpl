{{- define "resultpage" }}
<!DOCTYPE html>
<html>
    <head>
    <meta charset="utf-8">
    <title>Bellman-Zadeh Task</title>
	{{- range .JSAssets.Values }}
    <script src="{{ . }}"></script>
	{{- end }}
	{{- range .CustomizedJSAssets.Values }}
    <script src="{{ . }}"></script>
	{{- end }}
	{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
	{{- end }}
	{{- range .CustomizedCSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
	{{- end }}
</head>
<body>

	<div align="center">
		<a href="/"><h2>Back to load new data json</h2></a>
		<h1>Task Result:</h1>
		<textarea rows="80" cols="100" readonly="true" wrap="off">{{.ResultStringData}}</textarea>
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
