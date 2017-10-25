<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link
	href="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1502527377034&di=bd3b84165bf98b08e52c2b7618fdb853&imgtype=0&src=http%3A%2F%2Fec4.images-amazon.com%2Fimages%2FI%2F61QFHDUcvQL._SL1000_.jpg"
	rel="shortcut icon" type="image/x-icon" />
<meta name="description" content="Express">
<meta name="author" content="JerryChu">
<!-- Bootstrap core CSS -->
<link href="//cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css"
	rel="stylesheet">
<title>{{.Title}}</title>
</head>
<body>
	<nav class="navbar navbar-inverse navbar-fixed-top">
		<div class="container">
			<div class="navbar-header">
				<a class="navbar-brand" href="/express">{{.Info}}</a>
			</div>
		</div>
	</nav>

	<div class="jumbotron">
		<div class="container">
		    <h3>...</h3>
		</div>
	</div>

	<div class="container">
		<div class="row">
			<div class="col-md-2"></div>

			<div class="col-md-8">
				<form class="form-horizontal" id="userForm" method="POST"
					action="express/search">

					<div class="form-group">
						<label for="com_type" class="col-sm-2 control-label">快递公司</label>
						<div class="col-sm-3">
							<input type="text" class="form-control" id="com_type"
								name="com_type" placeholder="快递公司">
						</div>
					</div>

                    <div class="form-group">
                        <span for="speak" class="col-sm-10" style="font-size: 12px; color: gray; border: 1px solid #CCC; background-color: whiteSmoke; padding: 5px; margin: 5px 60px;"><p>快递公司编码:</p> <p>申通="shentong" EMS="ems" 顺丰="shunfeng" 圆通="yuantong" 中通="zhongtong" 韵达="yunda"</p> <p>天天="tiantian" 汇通="huitongkuaidi" 全峰="quanfengkuaidi" 德邦="debangwuliu" 宅急送="zhaijisong"</p></span>
                    </div>

					<div class="form-group">
						<label for="post_id" class="col-sm-2 control-label">快递单号</label>
						<div class="col-sm-4">
							<input type="text" class="form-control" id="post_id"
								name="post_id" placeholder="快递单号">
						</div>
					</div>

					<div class="form-group">
						<div class="col-sm-offset-3 col-sm-4">
							<input type="submit" class="btn btn-primary" name="submit" onclick=""
								value="查找">
						</div>
					</div>
				</form>
			</div>

			<div class="col-md-2"></div>

		</div>

		<hr>

		<footer>
			<nav class="navbar navbar-inverse ">
				<div class="container">
					<div class="navbar-header">
						<a class="navbar-brand" href="#">Welcome to go
							programming...</a>
					</div>
				</div>
			</nav>
		</footer>
	</div>
	<!-- /container -->
	<script type="text/javascript"
		src="http://code.jquery.com/jquery-latest.js"></script>
	<script type="text/javascript">
    </script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</body>
</html>
