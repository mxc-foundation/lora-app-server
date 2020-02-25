import React, { Component } from "react";
import { Route, Switch, Link, withRouter } from "react-router-dom";
import classNames from "classnames";
import { Nav, NavItem, Row, Col, Card, CardBody } from 'reactstrap';

import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import OrgBreadCumb from '../../components/OrgBreadcrumb';

import StakeStore from "../../stores/StakeStore";
import EthAccount from "./EthAccount";
import Transactions from "./Transactions";
import NetworkActivityHistory from "./NetworkActivityHistory";
import Stakes from "./Stakes";


class HistoryLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeTab: '0',
      loading: false,
      admin: false,
    };

    this.onChangeTab = this.onChangeTab.bind(this);
    this.locationToTab = this.locationToTab.bind(this);
  }

  componentDidMount() {
    /*window.analytics.page();*/
    const prevLoc = this.props.location.search.split('=')[1];
    this.setState({ loading: true });
    this.locationToTab(prevLoc);
    this.setState({ loading: false });
  }

  componentDidUpdate(oldProps) {
    if (this.props == oldProps) {
      return;
    }

    this.locationToTab();
  }

  onChangeTab(e, v) {
    this.setState({
      tab: v,
    });
  }

  locationToTab(prevLoc) {
    let tab = 0;
    if (window.location.href.endsWith("/eth-account")) {
      tab = 1;
    } else if (window.location.href.endsWith("/network-activity")) {
      tab = 2;
    } else if (window.location.href.endsWith("/stake")) {
      tab = 3;
    }

    this.setState({
      activeTab: tab + '',
    });
  }

  unstake = (e) => {
    e.preventDefault();
    const resp = StakeStore.unstake(this.props.match.params.organizationID);
    resp.then((res) => {
    })
  }

  render() {
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;

    return (
      <React.Fragment>
        <TitleBar>
          <OrgBreadCumb orgListCallback={() => { this.props.switchToSidebarId('DEFAULT'); }}
            orgNameCallback={() => { this.props.switchToSidebarId('DEFAULT'); }}
            organizationID={currentOrgID} items={[
              { label: i18n.t(`${packageNS}:tr000568`), active: false },
              { label: i18n.t(`${packageNS}:menu.history.history`), active: true }]}></OrgBreadCumb>
        </TitleBar>

        <Row>
          <Col>
            <Card>
              <CardBody>
                <Nav tabs>
                  <NavItem>
                    <Link
                      className={classNames('nav-link', { active: this.state.activeTab === '0' })}
                      to={`/history/${this.props.match.params.organizationID}/`}
                    >{i18n.t(`${packageNS}:menu.history.transactions`)}</Link>
                  </NavItem>
                  {this.state.admin && <NavItem>
                    <Link
                      className={classNames('nav-link', { active: this.state.activeTab === '1' })}
                      to={`/history/${this.props.match.params.organizationID}/eth-account`}
                    >{i18n.t(`${packageNS}:menu.history.eth_account`)}</Link>
                  </NavItem>}
                  <NavItem>
                    <Link
                      className={classNames('nav-link', { active: this.state.activeTab === '2' })}
                      to={`/history/${this.props.match.params.organizationID}/network-activity`}
                    >{i18n.t(`${packageNS}:menu.history.network_activity`)}</Link>
                  </NavItem>
                  <NavItem>
                    <Link
                      className={classNames('nav-link', { active: this.state.activeTab === '3' })}
                      to={`/history/${this.props.match.params.organizationID}/stake`}
                    >{i18n.t(`${packageNS}:menu.history.staking`)}</Link>
                  </NavItem>
                </Nav>

                <Row className="pt-3">
                  <Col>
                    <Card className="card-box shadow-sm">
                      <Switch>
                        <Route exact path={`${this.props.match.path}`} render={props => <Transactions organizationID={currentOrgID} {...props} />} />
                        <Route exact path={`${this.props.match.path}/eth-account`} render={props => <EthAccount organizationID={currentOrgID} {...props} />} />
                        <Route exact path={`${this.props.match.path}/network-activity`} render={props => <NetworkActivityHistory organizationID={currentOrgID} {...props} />} />
                        <Route exact path={`${this.props.match.path}/stake`} render={props => <Stakes organizationID={currentOrgID} {...props} />} />
                      </Switch>
                    </Card>
                  </Col>
                </Row>
              </CardBody>
            </Card>
          </Col>
        </Row>
      </React.Fragment>



    );
  }
}

export default withRouter(HistoryLayout);
