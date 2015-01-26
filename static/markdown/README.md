#js解析markdown转html

**下载源码，直接运行hello.html看简单的代码即可**

	<script type="text/javascript" src="showdown.js"></script>
	<script type="text/javascript">
		var converter = new Showdown.converter();
		var hello = converter.makeHtml('#hello!');
		// alert(hello);
	</script>
