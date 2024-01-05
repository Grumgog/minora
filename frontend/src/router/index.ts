import { createRouter, createWebHistory } from "vue-router";
import MainView from "@/views/MainView.vue";
import LoginView from "@/views/LoginView.vue";
import UserView from "@/views/UserView.vue";
import UsersView from "@/views/UsersView.vue";
import { useUserStore } from "@/stores/user";
import DashboardView from "@/views/DashboardView.vue";

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: "/",
			name: "main",
			component: MainView,
			children: [
				{
					path: "",
					name: "dashboard",
					component: DashboardView,
				},
				{
					path: "/users",
					name: "users",
					component: UsersView,
				},
				{
					path: "/user/:login",
					name: "user",
					component: UserView,
				},
			],
		},
		{
			path: "/about",
			name: "about",
			// route level code-splitting
			// this generates a separate chunk (About.[hash].js) for this route
			// which is lazy-loaded when the route is visited.
			component: () => import("../views/AboutView.vue"),
		},
		{
			path: "/login",
			name: "login",
			component: LoginView,
		},
	],
});

router.beforeEach((to, from) => {
	//const userStore = useUserStore();
	//if (!userStore.isLogined && to.name !== "login") {
	//	return { name: "login" };
	//}
});

export default router;
