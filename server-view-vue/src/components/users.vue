<template>
	<div class="container">
		<div class="row">
			<div class="col">
				<div class="btn btn-success" @click="addMqttUserShow">添加账号</div>
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col">
				用户名
			</div>
			<div class="col">
				密码
			</div>
			<div class="col">
				创建时间
			</div>
			<div class="col">
				操作
			</div>
		</div>
		<hr>
		<div class="user-list">
			<div class="row" v-for="(item, index) in userList" :key="item.uid">
				<div class="col">
					{{item.user}}
				</div>
				<div class="col">
					{{item.passwd}}

				</div>
				<div class="col">
					{{item.date}}
				</div>
				<div class="col">
					<div class="btn btn-danger" style="margin-right: 5px;" @click="delMqttUser(item.topic)">删除</div>
				</div>
				<hr>
			</div>

		</div>

		<popup ref="addUserPopup">
			<template #header>添加MQTT账户</template>
			<template #default>
				<table>
					<tbody>
						<tr>
							<td>账号：</td>
							<td><input type="text" v-model="mqttUserAccountInput"></td>
						</tr>
						<tr>
							<td>密码：</td>
							<td><input type="text" v-model="mqttUserPasswdInput"></td>
						</tr>
					</tbody>
				</table>
				<span style="color: red;">{{mqttUserAddHint}}</span>
			</template>
			<template #footer>
				<div class="btn btn-success" style="float: right;" @click="addMqttUser">确定</div>
			</template>
		</popup>

	</div>
</template>


<script lang="ts">
import { ref } from "vue"

import axioUtil from '@/util/axioUtil';

import popup from "./popup.vue"

export default {
	data() {
		return {
			mqttUserAccountInput: "",
			mqttUserPasswdInput: "",
			mqttUserAddHint: "",
			userList: [],

		}
	},
	components: {
		popup,
	},
	methods: {
		// 删除账号
		delMqttUser(topic: string) {
			let _this = (this as any);
			console.log(topic);
		},
		// 显示创建账户界面
		addMqttUserShow() {
			let _this = (this as any);
			_this.mqttUserAddHint = ""
			_this.$refs.addUserPopup.show()
		},
		// 创建 mqtt 账户
		async addMqttUser() {
			let _this = (this as any);
			let accoust = _this.mqttUserAccountInput;
			let passwd = _this.mqttUserPasswdInput;

			try {

				if (/[^a-z|A-Z|0-9]/g.test(accoust) || /[^a-z|A-Z|0-9]/g.test(passwd)) {
					_this.mqttUserAddHint = "【账号、密码】不能包含【字母、数字以】外的字符";

				}

				if (accoust.length < 6) {
					_this.mqttUserAddHint = "账号长度不能小于 6 .";
					return;
				}

				if (passwd.length < 6) {
					_this.mqttUserAddHint = ("密码长度不能小于 6 .");
					return;
				}


				let ret = await axioUtil.postAwait("/server/addMqttUser", {

				})
			} catch (error) {
				console.error(error);

				alert("添加失败！");
			}
		}
	},
	async mounted() {
		let _this = (this as any);

		// 加载 mqtt 节点账号列表
		try {
			let ret = await axioUtil.postAwait("/server/getMqttUserList", {
				token: localStorage.getItem("token")
			});

			if (ret.status != "ok") {
				throw ""
			}

			let _ul:any = [];
			ret.users.forEach((el: any) => {
				_ul.push({
					user: el.user,
					date: el.date,
				})
			});

			this.userList = _ul;

		} catch (error) {
			alert("获取节点列表失败");
		}
	},
}
</script>
