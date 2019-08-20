import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Router, Route, IndexRoute, browserHistory, applyRouterMiddleware } from 'react-router';
import useScroll from 'react-router-scroll/lib/useScroll';
import './index.css';
import App from './containers/App';
import * as serviceWorker from './serviceWorker';

import IndexPage from './containers/IndexPage';
import AdminLoginPage from './containers/AdminLoginPage';
import VotePage from './containers/VotePage';
import AdminApp from './containers/AdminApp';

const routes = (
    <Router history={browserHistory} render={applyRouterMiddleware(useScroll())}>
        <Route path="/" component={App}>
            <IndexRoute component={IndexPage} />
            <Route path="/vote" component={VotePage} />
            <Route path="/thank" component={VotePage} />    
        </Route>
        <Route path="/admin" component={AdminApp}>
            <Route path="/admin/login" component={AdminLoginPage} />
            <Route path="/admin" component={AdminLoginPage} />
            <Route path="/admin" component={AdminLoginPage} />
        </Route>
    </Router>
);

ReactDOM.render(
    <Provider>
      {routes}
    </Provider>,
    document.getElementById('root')
  );

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
