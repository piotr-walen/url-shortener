import React from "react";

export function FirstStep({
	namespace,
	onNamespaceChange,
	namespaceError,
}: {
	namespace: string;
	onNamespaceChange: React.ChangeEventHandler<HTMLInputElement>;
	namespaceError: string;
}) {
	return (
		<>
			<p>"Log in" to your namespace</p>
			<form>
				<fieldset>
					<label htmlFor="namespaceField">Namespace</label>
					<input
						className="input"
						type="text"
						placeholder=""
						id="namespaceField"
						value={namespace}
						onChange={onNamespaceChange}
					/>
					<div className="input-error">
						<sub>{namespaceError}</sub>
					</div>
				</fieldset>
			</form>
		</>
	);
}
