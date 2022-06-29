import React from "react";
import Stack from "./Stack";
import StyledText from "./StyledText";
import Button from './Button';

export default function CreatePortfolio(props) {
	async function createPortfolio(){
		// TODO: implement
	}

	return (
		<Stack>
			<StyledText styling='heading'>Create Portfolio</StyledText>
			<Button label="Create" onClick={createPortfolio} />
		</Stack>
	);
}