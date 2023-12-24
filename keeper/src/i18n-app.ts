import { createI18n } from "vue-i18n";
import ru_locale from "./assets/i18n/ru.json" with { type: "json" };
import en_locale from "./assets/i18n/en.json" with { type: "json" };

const messages = {
	en: en_locale,
	ru: ru_locale,
};

const i18n = createI18n({
	locale: "ru",
	fallbackLocale: "en",
	messages,
});

export default i18n;
