import React, { useEffect, useState } from 'react';
import AssetListItem from './AssetListItem';
import contactEndpoint from './contactEndpoint';
import InlineLayout from './InlineLayout';
import Stack from './Stack';
import CurrencyFormat from 'react-currency-format';
import StyledText from './StyledText';
import List from './List';
import Button from './Button';

export default function Login(props){

	async function startLoginFlow(){
		let data = await contactEndpoint('GET', 'oauth_login_url');
		window.location.href = data.OauthLoginUrl;
	}

	return (
		<Stack>
			<Button label="Login with Google" onClick={startLoginFlow}/>
		</Stack>
	);
}