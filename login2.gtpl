<html>
<head>
<title>kto do 5?</title>
</head>
<body>
<form action="/" method="post">
  first:<input type="checkbox" name="first" {{ .First }}><br>
  second:<input type="checkbox" name="second" {{ .Second }}><br>
  <input type="submit" value="OK">
</form>
</body>
</html>
