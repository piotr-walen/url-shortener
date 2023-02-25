// Base 64 Encoding with URL and Filename Safe Alphabet

import { ValidatorOptions } from "./types";

// https://datatracker.ietf.org/doc/html/rfc4648#section-5
const filenameSafeRegex = /^[a-zA-Z0-9_-]+$/;
export const validateFilenameSafe = (
	value: string,
	options: ValidatorOptions,
): string => {
	if (!filenameSafeRegex.test(value)) {
		return `${options.fieldName} field can only contain alphanumeric or "-", "_" characters`;
	}
	return "";
};
