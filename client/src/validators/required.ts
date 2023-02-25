import { ValidatorOptions } from "./types";

export const required = (value: string, options: ValidatorOptions): string => {
	if (!value) {
		return `${options.fieldName} is required`;
	}
	return "";
};
