<html>
<head>
   	<title>Загрузка файлов</title>
</head>
<body>
	<!-- action="http://127.0.0.1:80/upload" -->
<form enctype="multipart/form-data" method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{{.}}"/>
	<input type="submit" value="upload" />
</form>
</body>
</html>