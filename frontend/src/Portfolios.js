import React, { useEffect, useState } from "react";
import PortfolioTile from "./PortfolioTile";
import StyledText from "./StyledText";
import Stack from "./Stack";
import InlineLayout from "./InlineLayout";
import contactEndpoint from './contactEndpoint';
import Modal from "./Modal";
import CreatePortfolio from "./CreatePortfolio";
import Button from "./Button";

export default function Portfolios(props){
	const [portfolios, setPortfolios] = useState([]);
	const [createPortfolioIntent, setCreatePortfolioIntent] = useState(false);

	async function loadPortfolios(){
		try {
			let data = await contactEndpoint('GET', 'portfolios');
			setPortfolios(data.Portfolios);
		}catch(e){
			console.log("could not load portfolios: "+e);
		}
	}

	useEffect(
		() => {
			loadPortfolios();
		},
		[]
	);

	// TODO: show a message if there are no portfolios.
	return (
		<Stack>
			{
				createPortfolioIntent && (
					<Modal onClose={ () => { setCreatePortfolioIntent(false); }}>
						<CreatePortfolio/>
					</Modal>
				)
			}
			<InlineLayout>
				{
					portfolios.map(
						(portfolio) => {
							return (
								<React.Fragment key={portfolio.ID}>
									<PortfolioTile portfolio={portfolio} />
								</React.Fragment>
							);
						}
					)
				}
			</InlineLayout>
			<Button label="Create Portfolio" onClick={() => { setCreatePortfolioIntent(true); }}/>
		</Stack>
	);
}