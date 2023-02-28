import { describe, it, expect } from "vitest";
import { validateUrl } from "../validateUrl";

describe("validateUrl", () => {
    it("should error string if value is not url", () => {
        expect(validateUrl("foo", { fieldName: "Field" })).toBe("Field should contain valid http or https URL");
    });

    it("should empty string if value is valid url", () => {
        expect(validateUrl("https://foo.com", { fieldName: "Field" })).toBe("");
    });
});