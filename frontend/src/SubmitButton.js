import React from "react";
import Icon from "./Icon";
import InlineLayout from "./InlineLayout";

export default function SubmitButton(props){
	let classes = [];

	if(props.full){
		classes.push('full');
	}

	return (
		<button type="submit" className={classes.join(' ')}>
			<InlineLayout align='center'>
				{
					(props.icon && !props.icon_side_flip &&
						<Icon name={props.icon} />
					) || (
						null
					)
				}
				<div className='label'>
					{props.label}
				</div>
				{
					(props.icon && props.icon_side_flip &&
						<Icon name={props.icon} />
					) || (
						null
					)
				}
			</InlineLayout>
		</button>
	);
}