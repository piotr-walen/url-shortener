import React, { useCallback, useMemo, useReducer } from "react";
import { shorten } from "./api/shorten";
import "./App.css";
import { Loader } from "./components/Loader";

import { FirstStep } from "./components/steps/FirstStep";
import { SecondStep } from "./components/steps/SecondStep";
import { ThirdStep } from "./components/steps/ThirdStep";
import { BASE_URL } from "./constants";
import { reducer, initialState, stepOrder } from "./reducers/app";
import { required } from "./validators/required";
import { returnOnFirstError } from "./validators/returnOnFirstError";
import { validateFilenameSafe } from "./validators/validateFilenameSafe";
import { validateUrl } from "./validators/validateUrl";

function App() {
	const [state, dispatch] = useReducer(reducer, initialState);
	const formErrors = useMemo(
		() => ({
			targetUrl: returnOnFirstError(required, validateUrl)(
				state.formValues.targetUrl,
				{
					fieldName: "Target URL",
				},
			),
			namespace: returnOnFirstError(required, validateFilenameSafe)(
				state.formValues.namespace,
				{
					fieldName: "Namespace",
				},
			),
			segment: returnOnFirstError(required, validateFilenameSafe)( // TODO: refactor to field to "alias"
				state.formValues.segment,
				{ fieldName: "Alias" },
			),
		}),
		[state.formValues],
	);

	const canProceed = useCallback(() => {
		switch (state.step) {
			case "first":
				return !formErrors.namespace;
			case "second":
				return !(formErrors.targetUrl || formErrors.segment);
			default:
				return true;
		}
	}, [state.step, formErrors]);

	const onPrevClick: React.MouseEventHandler<HTMLButtonElement> = useCallback(
		(event) => {
			event.currentTarget.blur();
			dispatch({ type: "prevStep" });
		},
		[],
	);
	const onNextClick: React.MouseEventHandler<HTMLButtonElement> = useCallback(
		(event) => {
			event.currentTarget.blur();
			if (!canProceed()) {
				return;
			}
			if (state.step === "second") {
				dispatch({ type: "shorten" });
				shorten(state.formValues)
					.then((response) => {
						if (response.status === 201) {
							dispatch({
								type: "shortenSuccess",
								data: {
									status: "created",
									resultUrl: `${BASE_URL}/${state.formValues.namespace}/${state.formValues.segment}`,
								},
							});
							dispatch({ type: "nextStep" });
							return;
						}
						if (response.status === 409) {
							dispatch({
								type: "shortenFailed",
								data: { status: "alreadyExists" },
							});
							return;
						}
						if (response.status === 400) {
							dispatch({
								type: "shortenFailed",
								data: { status: "invalidPayload" },
							});
							return;
						}
						dispatch({ type: "shortenFailed", data: { status: "error" } });
					})
					.catch((error) => {
						dispatch({ type: "shortenFailed", data: { status: "error" } });
					});
				return;
			}
			dispatch({ type: "nextStep" });
		},
		[state],
	);

	const onNamespaceChange: React.ChangeEventHandler<HTMLInputElement> =
		useCallback((event) => {
			dispatch({
				type: "setNamespace",
				data: {
					value: event.target.value,
				},
			});
		}, []);
	const onTargetUrlChange: React.ChangeEventHandler<HTMLInputElement> =
		useCallback((event) => {
			dispatch({
				type: "setTargetUrl",
				data: {
					value: event.target.value,
				},
			});
		}, []);
	const onSegmentChange: React.ChangeEventHandler<HTMLInputElement> =
		useCallback((event) => {
			dispatch({
				type: "setSegment",
				data: {
					value: event.target.value,
				},
			});
		}, []);

	const getError = () => {
		const { errorStatus } = state;
		if (!errorStatus) {
			return null;
		}
		if (errorStatus === "alreadyExists") {
			return <sub>{"Url is already taken. Pl.ease enter another alias."}</sub>;
		}
		return (
			<sub>
				{"Internal error. Please "}
				<a href="report">report</a> {" the issue to page administrator."}
			</sub>
		);
	};

	return (
		<div className='App'>
			<div className="container box">
				<div>
					<h2 className="title">
						<strong>URL</strong>
					</h2>
					{state.step === "first" ? (
						<FirstStep
							namespace={state.formValues.namespace}
							onNamespaceChange={onNamespaceChange}
							namespaceError={formErrors.namespace}
						/>
					) : null}
					{state.step === "second" ? (
						<SecondStep
							namespace={state.formValues.namespace}
							targetUrl={state.formValues.targetUrl}
							segment={state.formValues.segment}
							onTargetUrlChange={onTargetUrlChange}
							onSegmentChange={onSegmentChange}
							targetUrlError={formErrors.targetUrl}
							segmentError={formErrors.segment}
						/>
					) : null}
					{state.step === "third" ? (
						<ThirdStep resultUrl={state.resultUrl} />
					) : null}
				</div>

				<div>
					<div className="error-box">{getError()}</div>
					<div className="controls">
						<button
							onClick={onPrevClick}
							disabled={stepOrder[0] === state.step}
						>
							Prev
						</button>
						<button
							onClick={onNextClick}
							disabled={stepOrder[stepOrder.length - 1] === state.step}
						>
							{state.isFetching ? <Loader /> : "Next"}
						</button>
					</div>
				</div>
			</div>
		</div>
	);
}

export default App;
