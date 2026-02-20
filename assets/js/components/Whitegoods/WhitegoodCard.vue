<template>
	<div class="col-12 col-md-6 mb-4">
		<div class="card h-100">
			<div class="card-body d-flex flex-column justify-content-between">
				<div>
					<h3 class="card-title">{{ name }}</h3>
					<div class="d-flex align-items-center mb-2">
						<span :class="statusBadgeClass">{{ statusLabel }}</span>
					</div>
					<div class="text-muted small">
						{{ $t("main.whitegoods.requiredPower") }}: {{ requiredPower }}W
					</div>
				</div>
				<div class="mt-3 text-end" v-if="status === 'idle'">
					<button
						class="btn btn-sm btn-outline-primary"
						@click="startAppliance"
						:disabled="starting"
					>
						{{ starting ? $t("main.whitegoods.starting") : $t("main.whitegoods.startNow") }}
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import api from "@/api";

export default defineComponent({
	name: "WhitegoodCard",
	props: {
		name: { type: String, required: true },
		status: { type: String, required: true },
		requiredPower: { type: Number, required: true },
	},
	data() {
		return {
			starting: false,
		};
	},
	computed: {
		statusBadgeClass() {
			switch (this.status) {
				case "running":
					return "badge text-bg-success";
				case "error":
					return "badge text-bg-danger";
				default:
					return "badge text-bg-secondary";
			}
		},
		statusLabel() {
			// Translating the states
			// You'd need to add these keys to `assets/i18n/` files (e.g. en.json, de.json)
			// main.whitegoods.status.idle, running, error, etc.
			return this.$t("main.whitegoods.status." + this.status) || this.status;
		},
	},
	methods: {
		async startAppliance(): Promise<void> {
			this.starting = true;
			try {
				await api.post(`/whitegoods/${encodeURIComponent(this.name)}/start`);
				// Notification or state will update via Websocket soon after
			} catch (err) {
				console.error("Failed to start whitegood:", err);
			} finally {
				this.starting = false;
			}
		},
	},
});
</script>
