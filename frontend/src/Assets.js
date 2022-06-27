import React, { useEffect, useState } from 'react';
import AssetListItem from './AssetListItem';
import contactEndpoint from './contactEndpoint';
import InlineLayout from './InlineLayout';
import Stack from './Stack';
import CurrencyFormat from 'react-currency-format';
import StyledText from './StyledText';
import List from './List';

export default function Assets(props){
	const [assets, setAssets] = useState([]);
	
	async function loadAssets(){
		try {
			let data = await contactEndpoint('GET', 'assets', {});
			setAssets(data.Assets);
		}catch(e){
			// TODO: use a page error user can see
			console.log("Could not load assets: "+e);
		}
	}

	useEffect(
		() => {
			loadAssets();
		},
		[]
	)

	// TODO: add paging

	return (
		<Stack>
			<List>
				{
					assets.map(
						(asset) => {
							return (
								<React.Fragment key={asset.ID}>
									<InlineLayout>
										<StyledText styling='title'>{asset.Name}</StyledText>
										<StyledText>{asset.Symbol}</StyledText>
										<StyledText><CurrencyFormat value={asset.ApproxPrice} displayType='text' prefix='$' thousandSeparator={true} decimalScale={2}/></StyledText>
									</InlineLayout>
								</React.Fragment>
							);
						}
					)
				}
			</List>
		</Stack>
	);
}