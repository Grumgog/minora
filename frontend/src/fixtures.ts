import dayjs from "dayjs";

export interface User {
	login: string;
	lastLogin: Date;
	ip: string;
	lastActionDate: Date;
	lastAction: string;
}

export function getLastUserFixtures(): User[] {
	const result: Array<User> = [
		{
			login: "Admin",
			ip: "127.0.0.1",
			lastAction: "Разрабатывает СУК",
			lastActionDate: new Date(),
			lastLogin: new Date(),
		},
		{
			login: "mihabsky",
			ip: "192.42.42.1",
			lastAction: "Играет в настолки",
			lastActionDate: new Date(),
			lastLogin: dayjs().subtract(20, "minutes").toDate(),
		},
		{
			login: "БеБ Ыб",
			ip: "192.56.23.56",
			lastAction: "Нелюбит HP",
			lastActionDate: dayjs().subtract(20, "minutes").toDate(),
			lastLogin: dayjs().subtract(75, "minutes").toDate(),
		},
		{
			login: "Kapitan Romka",
			ip: "192.42.42.65",
			lastAction: "Уехал к родителям",
			lastActionDate: dayjs().subtract(9, "hours").toDate(),
			lastLogin: dayjs().subtract(12, "hours").toDate(),
		},
	];
	return result;
}

export type UserInfo = User & { ActionHistory: { action: string; date: Date; affected?: string }[] };

export function getUsersFixtures(): Array<UserInfo> {
	const result: Array<UserInfo> = [
		{
			login: "Admin",
			ip: "127.0.0.1",
			lastAction: "Разрабатывает СУК",
			lastActionDate: new Date(),
			lastLogin: new Date(),
			ActionHistory: [
				{
					action: "Добавил тестовые данные",
					date: new Date(),
				},
				{
					action: "Добавил новый ресурс",
					date: dayjs().subtract(15, "minutes").toDate(),
					affected: "How to make CMS for russian users for goverment office workers",
				},
				{
					action: "Добавлен новый параметр",
					date: dayjs().subtract(45, "minutes").toDate(),
					affected: "inline photo",
				},
			],
		},
		{
			login: "mihabsky",
			ip: "192.42.42.1",
			lastAction: "Play Tabletops",
			lastActionDate: new Date(),
			lastLogin: dayjs().subtract(20, "minutes").toDate(),
			ActionHistory: [
				{
					action: "Штурмует крепость",
					date: new Date(),
				},
				{
					action: "Нанимает крестьян выращивать хмель",
					date: dayjs().subtract(15, "minutes").toDate(),
					affected: "Stronghold: Definitive Editon",
				},
				{
					action: "Добавил новый ресурс",
					date: dayjs().subtract(45, "minutes").toDate(),
					affected: "Разработка колоды MTG",
				},
			],
		},
		{
			login: "БеБ Ыб",
			ip: "192.56.23.56",
			lastAction: "Hate HP",
			lastActionDate: dayjs().subtract(20, "minutes").toDate(),
			lastLogin: dayjs().subtract(75, "minutes").toDate(),
			ActionHistory: [
				{
					action: "Добавил тестовые данные",
					date: new Date(),
				},
				{
					action: "Добавил новый ресурс",
					date: dayjs().subtract(15, "minutes").toDate(),
					affected: "Разборка принтера в оффисных условиях: руководство для бабы Зины",
				},
				{
					action: "Добавлен новый параметр",
					date: dayjs().subtract(45, "minutes").toDate(),
					affected: "summary",
				},
			],
		},
		{
			login: "Kapitan Romka",
			ip: "192.42.42.65",
			lastAction: "Return to Overwatch",
			lastActionDate: dayjs().subtract(9, "hours").toDate(),
			lastLogin: dayjs().subtract(12, "hours").toDate(),
			ActionHistory: [
				{
					action: "Добавил тестовые данные",
					date: new Date(),
				},
				{
					action: "Добавил новый ресурс",
					date: dayjs().subtract(15, "minutes").toDate(),
					affected: "Стельба за Ханзо, краткий гайд для начинающих",
				},
				{
					action: "Добавлен новый параметр",
					date: dayjs().subtract(45, "minutes").toDate(),
					affected: "fragMovie",
				},
			],
		},
	];
	return result;
}
