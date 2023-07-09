<template >
	<div>
		<div class="container-fluid">
			<div class="row header">
				<div class="col">MQTT 控制后台</div>
				<div class="col"></div>
				<div class="col">
					<div class="btn btn-success " style="float: right;" @click="loginExit">退出</div>
				</div>
			</div>
			<div class="row">
				<div class="col-2 side">
					<ul class="list-group">
						<li class="list-group-item" @click="gotoView('home')">首页</li>
						<li class="list-group-item" @click="gotoView('topic')">主题管理</li>
						<li class="list-group-item" @click="gotoView('users')">用户管理</li>
						<li class="list-group-item" @click="gotoView('supervisory')">实时监控</li>
					</ul>
				</div>
				<div class="col-10 main" style="padding-top: 10px;">
					<keep-alive>
						<component :is='nowView'></component>
					</keep-alive>
				</div>
			</div>
		</div>
	</div>
</template>
<style>


body {
	margin: 0px;
	padding: 0px;
	height: 100vh;
}

.header {
	background: #111;
	color: #eee;
	padding: 10px;
	height: 50px;
}

.side {
	background: #111;
	color: #eee;
	padding: 10px;
	height: calc(100vh - 50px);

	cursor: pointer;
}

.main {
	height: calc(100vh - 50px);
	overflow: auto;
}
</style>
<script lang="ts">

import axioUtil from "./util/axioUtil"

import home from "./components/home.vue"
import topic from "./components/topic.vue"
import users from "./components/users.vue"
import supervisory from "./components/supervisory.vue"

export default {
	setup() {
		return {
		}
	},
	data() {
		return {
			nowView: "home",
		}
	},
	components: {
		home, topic, users, supervisory,
	},
	methods: {
		async loginCheck() {
			if (!import.meta.env.DEV) {
				try {
					let ret: any = await axioUtil.postAwait("/server/checkLogin", {
						user: localStorage.getItem
					});
					if (ret.status != "ok") {
						return
					}
				} catch (error) {
					console.log(error);
				}

				location.href = "./login.html";
			}
		},
		loginExit(){
			localStorage.removeItem("user")
			localStorage.removeItem("token")
			location.href = "login.html"
		},
		gotoView(viewName: string) {
			this.nowView = viewName
		}
	},
	async mounted() {
		this.loginCheck();
	},

}
</script>
