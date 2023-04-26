<template>
	<div>
		<!-- 评论区 -->
		<div id="SOHUCS" :sid="sourceId"></div>
		<a href="#SOHUCS" id="changyan_count_unit"></a>
		<a href="#SOHUCS" id="changyan_parti_unit"></a>
	</div>
</template>
<script>
export default {
	name: 'Comment',
	props: {
		sourceId: {
			type: Number,
			require: false,
			default: -1
		}
	},
	mounted() {
		// 云评论
		(function () {
			var appid = 'cywE58eoF';
			var conf = 'prod_287c236820392be64661bae328651187';
			var width =
				window.innerWidth || document.documentElement.clientWidth;
			if (width < 1000) {
				var head =
					document.getElementsByTagName("head")[0] ||
					document.head ||
					document.documentElement;
				var script = document.createElement("script");
				script.type = "text/javascript";
				script.charset = "utf-8";
				script.id = "changyan_mobile_js";
				script.src =
					"https://cy-cdn.kuaizhan.com/upload/mobile/wap-js/changyan_mobile.js?client_id=" +
					appid +
					"&conf=" +
					conf;
				head.appendChild(script);
			} else {
				var loadJs = function (d, a) {
					var c =
						document.getElementsByTagName("head")[0] ||
						document.head ||
						document.documentElement;
					var b = document.createElement("script");
					b.setAttribute("type", "text/javascript");
					b.setAttribute("charset", "UTF-8");
					b.setAttribute("src", d);
					if (typeof a === "function") {
						if (window.attachEvent) {
							b.onreadystatechange = function () {
								var e = b.readyState;
								if (e === "loaded" || e === "complete") {
									b.onreadystatechange = null;
									a();
								}
							};
						} else {
							b.onload = a;
						}
					}
					c.appendChild(b);
				};
				loadJs(
					"https://cy-cdn.kuaizhan.com/upload/changyan.js",
					function () {
						window.changyan.api.config({
							appid: appid,
							conf: conf,
						});
					}
				);
			}
		})();
	},
}
</script>

<style></style>