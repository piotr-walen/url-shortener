import { describe, it, expect } from "vitest";
import { validateFilenameSafe } from "./validateFilenameSafe";


describe("required", () => {
    it("should return error string if value is not filename safe", () => {
        expect(validateFilenameSafe("foo/", { fieldName: "Field" })).toBe(`Field field can only contain alphanumeric or "-", "_" characters`);
    });

    it("should return empty string if value is valid filename", () => {
        expect(validateFilenameSafe("foo", { fieldName: "Field" })).toBe("");
    });
});
