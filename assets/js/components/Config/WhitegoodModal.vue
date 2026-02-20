<template>
	<DeviceModalBase
		:id="id"
		name="whitegood"
		device-type="whitegood"
		:is-sponsor="isSponsor"
		:modal-title="modalTitle"
		:provide-template-options="provideTemplateOptions"
		:initial-values="initialValues"
		:is-yaml-input-type="isYamlInput"
		:transform-api-data="transformApiData"
		@added="(name) => emitChanged('added', name)"
		@updated="() => emitChanged('updated')"
		@removed="() => emitChanged('removed')"
		@close="$emit('close')"
	/>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import DeviceModalBase from "./DeviceModal/DeviceModalBase.vue";
import { ConfigType } from "@/types/evcc";
import { getModal } from "@/configModal";
import {
	type DeviceValues,
	type Product,
	type ApiData,
} from "./DeviceModal";
import { type TemplateGroup } from "./DeviceModal/TemplateSelector.vue";

const initialValues = {
	type: ConfigType.Template,
	icon: undefined,
	deviceProduct: undefined,
	yaml: undefined,
	template: null,
};

export default defineComponent({
	name: "WhitegoodModal",
	components: {
		DeviceModalBase,
	},
	props: {
		isSponsor: Boolean,
	},
	emits: ["changed", "close"],
	data() {
		return {
		};
	},
	computed: {
		id(): number | undefined {
			return getModal("whitegood")?.id;
		},
		initialValues(): DeviceValues {
			const modal = getModal("whitegood");
			if (modal?.template) {
				return {
					...initialValues,
					template: modal.template,
					deviceProduct: modal.template,
					...modal.values,
				};
			}
			return initialValues;
		},
		modalTitle(): string {
			if (this.isNew) {
				return this.$t("config.whitegood.titleAdd");
			}
			return this.$t("config.whitegood.titleEdit");
		},
		isNew(): boolean {
			return this.id === undefined;
		},
	},
	methods: {
		provideTemplateOptions(products: Product[]): TemplateGroup[] {
			return [
				{
					label: "whitegoods",
					options: products,
				},
			];
		},
		transformApiData(data: ApiData): ApiData {
			if (data.type && this.isYamlInput(data.type)) {
				delete data.icon;
			}
			return data;
		},
		isYamlInput(type: ConfigType): boolean {
			return [ConfigType.Custom].includes(type);
		},
		async emitChanged(action: "added" | "updated" | "removed", name?: string) {
			const result = { action, name };
			this.$emit("changed", result);
		},
	},
});
</script>
