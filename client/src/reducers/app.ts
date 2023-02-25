export type Action =
	| { type: "nextStep" }
	| { type: "prevStep" }
	| { type: "shorten" }
	| { type: "shortenSuccess"; data: { status: "created"; resultUrl: string } }
	| {
			type: "shortenFailed";
			data: { status: "error" | "invalidPayload" | "alreadyExists" };
	  }
	| { type: "setNamespace"; data: { value: string } }
	| { type: "setTargetUrl"; data: { value: string } }
	| { type: "setSegment"; data: { value: string } };

export type Step = "first" | "second" | "third";

export const stepOrder: Step[] = ["first", "second", "third"];
function getNextStep(step: Step) {
	const index = stepOrder.findIndex((s) => s === step);
	if (index >= stepOrder.length - 1) {
		return step;
	}
	return stepOrder[index + 1];
}
function getPrevStep(step: Step) {
	const index = stepOrder.findIndex((s) => s === step);
	if (index <= 0) {
		return step;
	}
	return stepOrder[index - 1];
}

export type State = {
	step: Step;
	isFetching: boolean;
	formValues: {
		namespace: string;
		targetUrl: string;
		segment: string;
	};
	resultUrl: string;
	errorStatus: "error" | "invalidPayload" | "alreadyExists" | "";
};

export const initialState: State = {
	step: "first",
	isFetching: false,
	formValues: {
		namespace: "",
		targetUrl: "",
		segment: "",
	},
	resultUrl: "",
	errorStatus: "",
};

export function reducer(state: State, action: Action): State {
	switch (action.type) {
		case "prevStep":
			return {
				...state,
				step: getPrevStep(state.step),
				errorStatus: initialState.errorStatus,
			};
		case "nextStep":
			return {
				...state,
				step: getNextStep(state.step),
			};
		case "shorten":
			return {
				...state,
				isFetching: true,
				resultUrl: initialState.resultUrl,
				errorStatus: initialState.errorStatus,
			};
		case "shortenSuccess":
			return {
				...state,
				isFetching: false,
				resultUrl: action.data.resultUrl,
				errorStatus: initialState.errorStatus,
			};
		case "shortenFailed":
			return {
				...state,
				isFetching: false,
				resultUrl: initialState.resultUrl,
				errorStatus: action.data.status,
			};
		case "setNamespace":
			return {
				...state,
				formValues: {
					...state.formValues,
					namespace: action.data.value,
				},
				errorStatus: initialState.errorStatus,
			};
		case "setTargetUrl":
			return {
				...state,
				formValues: {
					...state.formValues,
					targetUrl: action.data.value,
				},
				errorStatus: initialState.errorStatus,
			};
		case "setSegment":
			return {
				...state,
				formValues: {
					...state.formValues,
					segment: action.data.value,
				},
				errorStatus: initialState.errorStatus,
			};
		default:
			return state;
	}
}
