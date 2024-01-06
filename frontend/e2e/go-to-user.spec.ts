import { expect, test } from "@playwright/test";

test("go to user", async ({ page }) => {
	// Переходим на страницу авторизации
	await page.goto("/");

	// Ввводим данные для входа
	await page.getByRole('cell').nth(0).dblclick();

	// Должно перенести на базовую страницу
	await expect(page).toHaveURL("/user/Admin");
});
