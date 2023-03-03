import './Loader.css'; 

export function Loader({
	type
}: {
	type?: 'primary' | 'secondary' | 'tertiary' | 'quaternary' | 'quinary';
}) {
	return (
		<div className={`lds-ellipsis ${type}`}>
			<div />
			<div />
			<div />
			<div />
		</div>
	);
}

Loader.defaultProps = {
	type: 'primary'
}