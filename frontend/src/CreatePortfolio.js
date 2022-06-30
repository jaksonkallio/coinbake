import React from "react";
import Stack from "./Stack";
import StyledText from "./StyledText";
import Button from './ActionButton';
import { useForm } from "react-hook-form";
import SubmitButton from "./SubmitButton";
import Input from "./Input";
import ErrorMessage from "./ErrorMessage";

export default function CreatePortfolio(props) {
	const { register, handleSubmit, watch, formState } = useForm();

	//console.log(watch("example")); // watch input value by passing the name of it

	async function createPortfolio(formdata){
		// TODO: implement
		console.log('creating portfolio, form data:');
		console.log(formdata);
	}

  return (
		<Stack>
			<StyledText styling='heading'>Create Portfolio</StyledText>
			<form onSubmit={handleSubmit(createPortfolio)}>
				<Stack>
					<Input type="text" label="Name of Portfolio" name="name" required register={register} />
					{formState.errors.name && <ErrorMessage>Name of portfolio is required.</ErrorMessage>}
					<SubmitButton label="Create" />
				</Stack>
			</form>
		</Stack>
  );
}