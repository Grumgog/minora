<script setup lang="ts">
import { getLastUserFixtures } from "@/fixtures";
import { ref } from "vue";
import dayjs from "dayjs";
import { useRouter } from "vue-router";

const router = useRouter();

const lastUsers = ref(getLastUserFixtures());
const selectedUser = ref(null);

function onRowDblClick(event): void {
	console.log("on rw dbl click", event.data.login);
	router.push({ name: "user", params: { login: event.data.login } });
}
</script>

<template>
	<Card>
		<template #title>{{ $t("lastUser.title") }}</template>
		<template #content>
			<DataTable
				:value="lastUsers"
				size="small"
				selection-mode="single"
				v-model:selection="selectedUser"
				:meta-key-selection="true"
				@row-dblclick="onRowDblClick($event)">
				<Column field="login" :header="$t('lastUser.login')"></Column>
				<Column field="lastLogin" :header="$t('lastUser.lastLogin')">
					<template #body="{ data }">
						{{ dayjs(data.loastLogin).format("DD.MM.YYYY HH:mm:ss") }}
					</template>
				</Column>
				<Column field="ip" header="IP"></Column>
				<Column field="lastActionDate" :header="$t('lastUser.lastActionDate')">
					<template #body="{ data }">
						{{ dayjs(data.lastActionDate).format("DD.MM.YYYY HH:mm:ss") }}
					</template>
				</Column>
				<Column field="lastAction" :header="$t('lastUser.lastAction')"></Column>
			</DataTable>
		</template>
	</Card>
</template>
