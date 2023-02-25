import React from "react";

export function ThirdStep({ resultUrl }: { resultUrl: string }) {
	return (
		<>
			<p>Here's your new url</p>
			<a className="url" href={resultUrl}>
				{resultUrl}
			</a>
		</>
	);
}
