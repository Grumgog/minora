import { expect, test } from "@playwright/test";

test("login into system", async ({ page }) => {
	// Переходим на страницу авторизации
	await page.goto("/login");

	// Ввводим данные для входа
	await page.locator("#login").fill("admin");
	await page.locator("#password").fill("admin");

	// Нажимаем кнопку "войти"
	await page.locator('button[type="submit"]').click();

	// Должно перенести на базовую страницу
	await expect(page).toHaveURL("/");
});
