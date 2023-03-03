import { BASE_URL } from "../constants";
import { requestInit } from "./requestInit";

export type ShortenStatus =
	| "created"
	| "alreadyExists"
	| "invalidPayload"
	| "serverError";

const getStatus = ({ status }: { status: number }): ShortenStatus => {
	if (status === 201) return "created";
	if (status === 409) return "alreadyExists";
	if (status === 400) return "invalidPayload";
	return "serverError";
};

export const shorten = (payload: {
	targetUrl: string;
	namespace: string;
	segment: string;
}) =>
	fetch(`${BASE_URL}/url-shorten`, {
		...requestInit,
		method: "POST",
		body: JSON.stringify(payload),
	}).then((response) => ({
		status: getStatus(response),
	}));

