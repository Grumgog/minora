import dayjs from "dayjs";

export function getLastUserFixtures(): {
	login: string;
	lastLogin: Date;
	ip: string;
	lastActionDate: Date;
	lastAction: string;
}[] {
	return [
		{
			login: "Admin",
			ip: "127.0.0.1",
			lastAction: "Development CMS",
			lastActionDate: new Date(),
			lastLogin: new Date(),
		},
		{
			login: "mihabsky",
			ip: "192.42.42.1",
			lastAction: "Play Tabletops",
			lastActionDate: new Date(),
			lastLogin: dayjs().subtract(20, "minutes").toDate(),
		},
		{
			login: "БеБ Ыб",
			ip: "192.56.23.56",
			lastAction: "Hate HP",
			lastActionDate: dayjs().subtract(20, "minutes").toDate(),
			lastLogin: dayjs().subtract(75, "minutes").toDate(),
		},
		{
			login: "Kapitan Romka",
			ip: "192.42.42.65",
			lastAction: "Return to Overwatch",
			lastActionDate: dayjs().subtract(9, "hours").toDate(),
			lastLogin: dayjs().subtract(12, "hours").toDate(),
		},
	];
}
