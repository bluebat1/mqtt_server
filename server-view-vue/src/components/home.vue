<template>
	<div class="container">
		<hr>
		<div class="row">
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">节点账户资源</h4>
						<p class="card-text">{{nodeAccountMax}}</p>
					</div>
				</div>
			</div>
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">当前节点账户数量</h4>
						<p class="card-text">{{nodeAccountCount}}</p>
					</div>
				</div>
			</div>
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">剩余节点账户资源</h4>
						<p class="card-text">{{nodeAccountResidue}}</p>
					</div>
				</div>
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">主题资源</h4>
						<p class="card-text">{{TopicMax}}</p>
					</div>
				</div>
			</div>
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">当前主题数量</h4>
						<p class="card-text">{{TopicCount}}</p>
					</div>
				</div>
			</div>
			<div class="col">
				<div class="card">
					<div class="card-body">
						<h4 class="card-title">剩余主题资源</h4>
						<p class="card-text">{{TopicResidue}}</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<script lang="ts">
import axioUtil from '@/util/axioUtil';

export default {
	data() {
		return {
			nodeAccountMax: "0",
			nodeAccountResidue: "0",
			nodeAccountCount: "0",
			TopicMax: "0",
			TopicCount: "0",
			TopicResidue: "0",
		}
	},
	async mounted() {
		try {
			let ret = await axioUtil.postAwait("/server/getUserAssets", {
				token: localStorage.getItem("token")
			});
			if (ret.status != "ok") {
				return;
			}
			console.log(ret);

			this.nodeAccountMax = ret.nodeAssetsMax;
			this.nodeAccountCount = ret.nodeAssetsNum;
			this.nodeAccountResidue = Number(this.nodeAccountMax) - Number(this.nodeAccountCount)
			this.TopicMax = ret.topicAssetsMax;
			this.TopicCount = ret.topicAssetsNum;
			this.TopicResidue = Number(this.TopicMax) - Number(this.TopicCount)
		} catch (error) {
			alert("加载失败");
		}
	},
}
</script>