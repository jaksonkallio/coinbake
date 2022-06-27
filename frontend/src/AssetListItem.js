import React from "react";

export default function AssetListItem(props) {
	return (
		<div className='asset_list_item'>
			{props.asset.Symbol}
		</div>
	);
}