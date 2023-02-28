import { describe, expect, it } from 'vitest';
import { returnOnFirstError } from "../returnOnFirstError";

describe("returnOnFirstError", () => {
	it("should return first error", () => {
		const validator = returnOnFirstError(
			(value: string) => (value ? "" : "Required"),
			(value: string) =>	/^[a-zA-Z0-9]*$/.test(value) ? "" : "Alfanumeric only"
		);
		expect(validator("__", { fieldName: "Foo" })).toBe("Alfanumeric only")
	});

	it('should return empty string if no errors', () => {
		const validator = returnOnFirstError(
			(value: string) => (value ? "" : "Required"),
			(value: string) =>	/^[a-zA-Z0-9]*$/.test(value) ? "" : "Alfanumeric only"
		);
		expect(validator("Foo", { fieldName: "Foo" })).toBe("")
	});

	it("should pass options to validators", () => {
		const validator = returnOnFirstError(
			(value: string, options: { fieldName: string }) => (value ? "" : `${options.fieldName} is required`),
			(value: string, options: { fieldName: string }) =>	/^[a-zA-Z0-9]*$/.test(value) ? "" : `${options.fieldName} must be alfanumeric`
		);
		expect(validator("", { fieldName: "Foo" })).toBe("Foo is required")
	});
});
