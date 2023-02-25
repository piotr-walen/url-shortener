import React from "react";
import { BASE_URL } from "../../constants";

export function SecondStep({
	namespace,
	targetUrl,
	segment,
	onTargetUrlChange,
	onSegmentChange,
	targetUrlError,
	segmentError,
}: {
	namespace: string;
	targetUrl: string;
	segment: string;
	onTargetUrlChange: React.ChangeEventHandler<HTMLInputElement>;
	onSegmentChange: React.ChangeEventHandler<HTMLInputElement>;
	targetUrlError: string;
	segmentError: string;
}) {
	return (
		<>
			<p>Enter URL that you would like to create alias for</p>
			<form>
				<fieldset>
					<label htmlFor="urlField">Target URL</label>
					<input
						className="input"
						type="text"
						placeholder=""
						id="urlField"
						value={targetUrl}
						onChange={onTargetUrlChange}
					/>
					<div className="input-error">
						<sub>{targetUrlError}</sub>
					</div>
					<label htmlFor="aliasField">Alias</label>
					<div style={{ display: "flex", alignItems: "center" }}>
						<div className="segment-label">
							{BASE_URL}/{namespace}/
						</div>
						<input
							className="input"
							type="text"
							placeholder=""
							id="aliasField"
							value={segment}
							onChange={onSegmentChange}
						/>
					</div>
					<div className="input-error">
						<sub>{segmentError}</sub>
					</div>
				</fieldset>
			</form>
		</>
	);
}
