import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import InlineLayout from './InlineLayout.js';
import Icon from './Icon.js';

const ProcessingStatus = {
	IDLE: 1,
	PROCESSING: 2
};

export default function Button(props){
	const [processing_status, setProcessingStatus] = useState(ProcessingStatus.IDLE);

	async function onClick(){
		if(!props.onClick){
			throw 'No on-click action defined';
		}

		if(processing_status != ProcessingStatus.IDLE){
			// Processing status is not idle, don't do click event
			return;
		}

		try {
			// Mark as processing
			setProcessingStatus(ProcessingStatus.PROCESSING);
			await props.onClick();
		}finally{
			setProcessingStatus(ProcessingStatus.IDLE);
		}
	}

	let classes = ['button'];

	if(['standard', 'subtle', 'text'].includes(props.styling)){
		classes.push(props.styling);
	}

	if(props.full){
		classes.push('full');
	}

	if(processing_status == ProcessingStatus.PROCESSING){
		classes.push('processing');
	}

	const button_content = (
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
	);

	if(typeof props.onClick == 'string'){
		return (
			<Link className={classes.join(' ')} to={props.onClick}>
				{button_content}
			</Link>
		);
	}else{
		return (
			<a className={classes.join(' ')} onClick={onClick}>
				{
					(processing_status == ProcessingStatus.PROCESSING &&
						<img className='spinner' src='./static/img/spinner.svg'/>
					)
				}
				{button_content}
			</a>
		);
	}
}