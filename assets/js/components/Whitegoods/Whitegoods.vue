<template>
	<div class="whitegoods">
		<h2 class="mb-3 d-flex align-items-center">
			{{ $t("config.section.whitegoods") }}
		</h2>
		<div class="row">
			<WhitegoodCard
				v-for="(wg, index) in whitegoods"
				:key="index"
				v-bind="wg"
			/>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, type PropType } from "vue";
import type { Whitegood } from "@/types/evcc";
import WhitegoodCard from "./WhitegoodCard.vue";

export default defineComponent({
	name: "Whitegoods",
	components: { WhitegoodCard },
	props: {
		whitegoods: { type: Array as PropType<Whitegood[]>, default: () => [] },
	},
});
</script>

<style scoped>
.whitegoods {
	margin-top: 2rem;
}
</style>
