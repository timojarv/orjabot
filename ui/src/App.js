import React, { useState } from 'react';
import { Navbar, Button, Alignment, Card, Elevation, H3 } from '@blueprintjs/core';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import * as api from './api';
import GroupSelector from './GroupSelector';
import Chats from './Chats';

export const GroupContext = React.createContext(false);

const App = props => {
	const [user, setUser] = useState(false);
	const [group, setGroup] = useState(false);

	const handleLogin = () => api.login()
		.then(res => {
			setUser(res.user);
		});

	const handleLogout = () => api.logout()
		.then(() => {
			setUser(false);
		})

	return (
		<div className="" style={{ width: '90%', margin: 'auto', padding: '2rem' }}>
			<GroupContext.Provider value={group}>
				<Router>
					<Navbar style={{ marginBottom: '2rem' }}>
						<Navbar.Group align={Alignment.LEFT}>
							<Navbar.Heading>OrjaData</Navbar.Heading>
							<Navbar.Divider />
							<Link to="/">
								<Button minimal icon="home" text="Etusivu" />
							</Link>
							<Link to="/chats">
								<Button minimal icon="chat" text="Chatit" />
							</Link>
							<Button minimal disabled icon="briefcase" text="Soppa" />
							<Button minimal disabled icon="feed" text="Tiedotus" />
						</Navbar.Group>
						<Navbar.Group align={Alignment.RIGHT}>
							<GroupSelector onChange={groupID => setGroup(groupID)} />
							{/* <Navbar.Divider />
							<em>{user.displayName}</em>
							<Navbar.Divider />
							{
								user
									? <Button onClick={handleLogout} minimal icon="log-out" text="Kirjaudu ulos" />
									: <Button onClick={handleLogin} minimal icon="log-in" text="Kirjaudu sisään" />
							} */}
						</Navbar.Group>
					</Navbar>
					<Switch>
						<Route exact path="/">
							<Card elevation={Elevation.TWO}>
								<H3>Tervetuloa OrjaDataan.</H3>
								<p>Koska tietojohtaminen on selkeää ja helppoa.</p>
							</Card>
						</Route>

						<Route path="/chats" component={Chats} />
					</Switch>
				</Router>
			</GroupContext.Provider>
		</div>
	);
}

export default App;
