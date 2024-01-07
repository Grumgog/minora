import { test, expect } from "@playwright/test";

test.describe("Welcome splash panel", () => {
	test.beforeEach(async ({ page }) => {
		await page.goto("/");
	});

	test("Close splash", async ({ page }) => {
		await page.getByTestId("welcome-splash-close").click();
		expect(await page.getByTestId("welcome-splash").count()).toEqual(0);
	});

	test("Link of Splash screen is correct", ({ page }) => {});
});
