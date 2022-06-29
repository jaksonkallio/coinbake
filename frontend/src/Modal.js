import React from "react";

export default function Modal(props){
	return (
		<div className="modal_background">
			<div className="modal">
				<div
					className="close_button"
					onClick={props.onClose}
				>
					<img src="static/img/close.svg" />
				</div>
				{props.children}
			</div>
		</div>
	);
}