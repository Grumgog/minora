import { ref, computed } from "vue";
import { defineStore } from "pinia";
import { getUsersFixtures, type UserInfo } from "@/fixtures";

export const useUsersStore = defineStore("users", () => {
	async function getUsers(): Promise<Array<UserInfo>> {
		return getUsersFixtures();
	}

	return { getUsers };
});
