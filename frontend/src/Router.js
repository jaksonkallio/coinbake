import React from "react";
import { Route, Switch } from 'react-router-dom';
import { HashRouter } from 'react-router-dom';
import Portfolios from "./Portfolios";
import Nav from "./Nav";
import Portfolio from "./Portfolio";
import Assets from "./Assets";
import Login from "./Login";

export default function Router(props){
	return (
		<HashRouter basename="/">
			<Nav/>
			<Switch>
				<Route path='/assets' component={Assets} />
				<Route path='/portfolios' component={Portfolios} />
				<Route path='/portfolio/:id' component={Portfolio} />
				<Route path='/login' component={Login} />
				<Route path='/' component={Portfolios} />
			</Switch>
		</HashRouter>
	);
}