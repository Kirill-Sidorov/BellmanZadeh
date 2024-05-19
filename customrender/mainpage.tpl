{{- define "mainpage" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
    {{- template "base" . }}
	<style>
		.container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
		.item {margin: auto;}
	</style>

	<div align="center">
		<textarea rows="80" cols="100" readonly="true">{{.ResultStringData}}</textarea>
	</div>

</body>
</html>
{{ end }}