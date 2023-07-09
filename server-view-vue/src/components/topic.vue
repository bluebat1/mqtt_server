<template>
	<div class="container">
		<div class="row">
			<div class="col">
				<div class="btn btn-success">添加主题</div>
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col">
				主题
			</div>
			<div class="col">
				订阅量
			</div>
			<div class="col">
				消息总量
			</div>
			<div class="col">
				操作
			</div>
		</div>
		<hr>
		<div class="theme-list">
			<div class="row" v-for="(item, index) in topicList" :key="index">
				<div class="col">
					{{item.topic}}
				</div>
				<div class="col">
					{{item.subscribeNum}}

				</div>
				<div class="col">
					{{item.msgNum}}
				</div>
				<div class="col">
					<div class="btn btn-danger" style="margin-right: 5px;" @click="delTopic(item.topic)">删除</div>
					<div class="btn btn-success" style="margin-right: 5px;" @click="showSubscribeList(item.topic)">查看订阅者
					</div>
					<div class="btn btn-success" style="margin-right: 5px;" @click="showMsgList(item.topic)">查看消息</div>
				</div>
				<hr>
			</div>

		</div>
	</div>
</template>
<script lang="ts" setup>

let topicList: any[] = [];

// 
for (let i = 0; i < 20; i++) {
	topicList.push({
		topic: "主题" + i,
		subscribeNum: Math.floor(Math.random() * 100),
		msgNum: Math.floor(Math.random() * 100),
	})
}

// 删除主题
let delTopic = (topic: string) => {
	console.log(topic);
}

// 查看订阅者
let showSubscribeList = (topic: string) => {
	console.log(topic);
}

// 查看消息
let showMsgList = (topic: string) => {
	console.log(topic);
}



</script>
