import { BASE_URL } from "../constants";
import { requestInit } from "./requestInit";

export const shorten = (payload: {
	targetUrl: string;
	namespace: string;
	segment: string;
}) =>
	fetch(`${BASE_URL}/url-shorten`, {
		...requestInit,
		method: "POST",
		body: JSON.stringify(payload),
	});
