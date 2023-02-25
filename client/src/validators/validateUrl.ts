import { ValidatorOptions } from "./types";

function isValidHttpUrl(url: string) {
	try {
		const newUrl = new URL(url);
		return newUrl.protocol === "http:" || newUrl.protocol === "https:";
	} catch (err) {
		return false;
	}
}

export const validateUrl = (value: string, options: ValidatorOptions) => {
	if (!isValidHttpUrl(value)) {
		return `${options.fieldName} should contain valid http or https URL`;
	}
	return "";
};
