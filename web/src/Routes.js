import {
    BrowserRouter as Router,
    Link,
    Route,
    Switch,
} from 'react-router-dom';
import React from 'react';
import LandingPageContainer from "./containers/LandingPageContainer";

const Routes = () => (
    <Router>
        <div>
            {/*<Link to="/">Home</Link>{' '}*/}
            {/*<Link to={{pathname: '/about'}}>About</Link>{' '}*/}
            {/*<Link to="/contact">Contact</Link>*/}

            <Switch>
                <Route path="/" exact component={LandingPageContainer} />
                <Route path="/about" component={LandingPageContainer} />
                <Route
                    path="/contact"
                    render={() => <h1>Contact Us</h1>} />
                <Route path="/blog" children={({match}) => (
                    <li className={match ? 'active' : ''}>
                        <Link to="/blog">Blog</Link>
                    </li>)} />
                <Route render={() => <h1 className="text-center">Page not found</h1>} />
            </Switch>
        </div>
    </Router>
);

export default Routes;
