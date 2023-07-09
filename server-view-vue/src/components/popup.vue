<template>
	<div class="popup" v-if="isShow">
		<div class="card box">
			<div class="card-header">
				<slot name="header"></slot>
			</div>
			<div class="card-body">
				<slot></slot>
			</div>
			<div class="card-footer">
				<slot name="footer"></slot>
			</div>
		</div>
	</div>
</template>
<style scoped>
.popup {
	position: fixed;
	top: 0px;
	left: 0px;
	width: 100vw;
	height: 100vh;
	display: flex;
	align-items: center;
	justify-content: center;
	z-index: 999;
	background: #0008;

}

.box {
	box-shadow: -5px 5px 20px #000;
	min-width: 300px;
	min-height: 200px;
}
</style>
<script lang="ts">
export default {
	data() {
		return {
			isShow: false,
		}
	},
	methods: {
		show() {
			(this as any).isShow = true;
		},
		hied() {
			(this as any).isShow = false;
		}
	},
}
</script>

<!-- 
	<popup ref="popup">
		<template v-slot:header></template>
		<template></template>
		<template v-slot:footer></template>
	</popup>
 -->