import { describe, it, expect } from "vitest";
import { required } from "./required";


describe("required", () => {
    it("should return error if value is empty", () => {
        expect(required("", { fieldName: "Field" })).toBe("Field is required");
    });

    it("should return error if value is not empty", () => {
        expect(required("foo", { fieldName: "Field" })).toBe("");
    });
});
