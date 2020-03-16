import React, { Component } from "react";
import { withRouter } from 'react-router-dom';
import classNames from "classnames";

import { Row, Col, Card, CardBody, Nav, NavItem, NavLink, TabContent, TabPane } from 'reactstrap';

import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import OrgBreadCumb from '../../components/OrgBreadcrumb';

import SessionStorage from "../../stores/SessionStore";
import TopupCrypto from "./TopupCrypto";
import TopupHistory from "./TopupHistory";

import { Alert } from 'reactstrap';


class Topup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeTab: "0",
    };

    this.onTabToggle = this.onTabToggle.bind(this);
  }

  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
    }
  }

  onTabToggle(tab) {
    this.setState({ activeTab: tab });
  }

  onSubmit = () => {
    /* if (SessionStorage.getUser().isAdmin) {
      this.props.history.push(`/control-panel/modify-account`);
    } else {
      this.props.history.push(`/modify-account/${this.props.match.params.organizationID}`);
    } */
  }

  render() {
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;

    return (<React.Fragment>
      <TitleBar>
        <OrgBreadCumb orgListCallback={() => { this.props.switchToSidebarId('DEFAULT'); }}
          orgNameCallback={() => { this.props.switchToSidebarId('DEFAULT'); }}
          organizationID={currentOrgID} items={[
            { label: i18n.t(`${packageNS}:tr000568`), active: false },
            { label: i18n.t(`${packageNS}:menu.topup.topup`), active: true }]}></OrgBreadCumb>
      </TitleBar>
      <div>
        <Alert color="info">
          <h4 className="alert-heading">{i18n.t(`${packageNS}:menu.topup.info.title`)}</h4>
            <p>
              {i18n.t(`${packageNS}:menu.topup.info.notice`)}
            </p>
            <hr />
          <p className="mb-0"></p>
        </Alert>
      </div>
      <Row>
        <Col>
          <Card>
            <CardBody className="pb-0">
              <Nav tabs>
                <NavItem>
                  <NavLink
                    className={classNames('nav-link', { active: this.state.activeTab === '0' })} href='#'
                    onClick={(e) => this.onTabToggle("0")}
                  >{i18n.t(`${packageNS}:menu.topup.crypto`)}</NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classNames('nav-link', { active: this.state.activeTab === '1' })} href='#'
                    onClick={(e) => this.onTabToggle("1")} disabled
                  >{i18n.t(`${packageNS}:menu.topup.otc`)}</NavLink>
                </NavItem>
              </Nav>

              <TabContent activeTab={this.state.activeTab}>
                <TabPane tabId="0">
                  <TopupCrypto />
                </TabPane>
                <TabPane tabId="1">
                </TabPane>
              </TabContent>

            </CardBody>
          </Card>
        </Col>
      </Row>

      <TopupHistory organizationID={currentOrgID} />
    </React.Fragment>
    );
  }
}

export default withRouter(Topup);