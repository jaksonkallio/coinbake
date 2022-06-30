import React from "react";
import Stack from "./Stack";
import StyledText from "./StyledText";

export default function Input(props){
	return (
		<Stack gap_size='shim'>
			{props.label &&
				<label>{props.label}</label>
			}
			{props.note &&
				<StyledText styling='note'>{props.note}</StyledText>
			}
			<input type={props.type} {...props.register(props.name, { required: props.required })} />
		</Stack>
	);
}