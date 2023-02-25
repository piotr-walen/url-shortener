import { ValidatorOptions } from "./types";

export const returnOnFirstError =
	(...chain: ((value: string, options: ValidatorOptions) => string)[]) =>
	(value: string, options: ValidatorOptions): string => {
		for (const validator of chain) {
			const result = validator(value, options);
			if (result) {
				return result;
			}
		}
		return "";
	};
