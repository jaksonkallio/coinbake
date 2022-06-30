import React from 'react';

export default function ErrorMessage(props) {
	return (
		<div className="error_message">
			{props.children}
		</div>
	);
}